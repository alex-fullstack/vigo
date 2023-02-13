package infrastructure

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"processing/pkg/services"
	"strings"
	"sync"
	"time"
)

type NatsConsumer struct {
	conn              *nats.Conn
	streamName        string
	replaySubjectName string
	jsc               nats.JetStreamContext
	wg                *sync.WaitGroup
	messageChannels   map[BusSubject]chan *nats.Msg
	subscriptions     map[BusSubject]*nats.Subscription
	messageProcessor  services.MessageProcessor
}

func NewNatsConsumer(
	conn *nats.Conn,
	streamName string,
	jsc nats.JetStreamContext,
	wg *sync.WaitGroup,
	subjects []BusSubject,
	replaySubjectName string,
	messageProcessor services.MessageProcessor,
) *NatsConsumer {
	n := &NatsConsumer{
		conn:              conn,
		messageChannels:   make(map[BusSubject]chan *nats.Msg),
		subscriptions:     make(map[BusSubject]*nats.Subscription),
		messageProcessor:  messageProcessor,
		streamName:        streamName,
		replaySubjectName: replaySubjectName,
		wg:                wg,
		jsc:               jsc,
	}
	for _, subject := range subjects {
		n.messageChannels[subject] = make(chan *nats.Msg, 64)
	}
	return n
}

func (nc *NatsConsumer) Start() {
	for subject, ch := range nc.messageChannels {
		for {
			sub, err := nc.jsc.ChanSubscribe(
				strings.ToLower(nc.streamName)+"."+string(subject),
				ch,
				nats.BindStream(nc.streamName),
			)
			if err == nil {
				nc.subscriptions[subject] = sub
				break
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
	go func() {
		for nc.conn.IsConnected() {
			time.Sleep(time.Second)
		}
		nc.flush()
	}()
	for _, ch := range nc.messageChannels {
		nc.wg.Add(1)
		go func(ch chan *nats.Msg) {
			for {
				select {
				case msg, ok := <-ch:
					if ok {
						nc.handleMessageAsync(msg)
					} else {
						nc.wg.Done()
						return
					}
				}
			}
		}(ch)
	}
	fmt.Println("consumer started!")
}

func (nc *NatsConsumer) flush() {
	for subject, ch := range nc.messageChannels {
		close(ch)
		nc.subscriptions[subject].Unsubscribe()
	}
}

func (nc *NatsConsumer) handleMessageAsync(msg *nats.Msg) {
	go func(msg *nats.Msg, jsc nats.JetStreamContext) {
		subject := msg.Subject
		_, after, found := strings.Cut(subject, strings.ToLower(nc.streamName+"."))
		if !found {
			panic(fmt.Errorf("handling message with unknown subject: %s", subject))
		}
		replayMessage, err := nc.processMessage(BusSubject(after), msg.Data)
		if err != nil {
			panic(err)
		}
		_, err = jsc.PublishAsync(subject+"."+nc.replaySubjectName, replayMessage)
		if err != nil {
			panic(err)
		}
		err = msg.Ack()
		if err != nil {
			panic(err)
		}
	}(msg, nc.jsc)
}

func (nc *NatsConsumer) processMessage(subject BusSubject, msg []byte) (replayMessage []byte, err error) {
	switch subject {
	case FileSpecificationGetSubject:
		replayMessage, err = nc.messageProcessor.OnGetFileSpecification(msg)
	default:
		err = fmt.Errorf("processing message with unknown subject: %s", subject)
	}
	return
}
