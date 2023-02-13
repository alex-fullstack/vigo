package infrastructure

import (
	"github.com/nats-io/nats.go"
	"processing/pkg/services"
	"strings"
	"sync"
)

type NatsBroker struct {
	url                    string
	streamName             string
	replaySubjectName      string
	publishAsyncMaxPending int
	wg                     *sync.WaitGroup
	messageProcessor       services.MessageProcessor
}

type NatsBrokerConfig struct {
	Url                    string
	StreamName             string
	PublishAsyncMaxPending int
	ReplaySubjectName      string
}

func NewNatsBrokerConfig(url string, streamName string, publishAsyncMaxPending int, replaySubjectName string) NatsBrokerConfig {
	return NatsBrokerConfig{
		Url:                    url,
		StreamName:             streamName,
		PublishAsyncMaxPending: publishAsyncMaxPending,
		ReplaySubjectName:      replaySubjectName,
	}
}

func NewNatsBroker(cfg NatsBrokerConfig, messageProcessor services.MessageProcessor) *NatsBroker {
	return &NatsBroker{
		url:                    cfg.Url,
		streamName:             cfg.StreamName,
		replaySubjectName:      cfg.ReplaySubjectName,
		publishAsyncMaxPending: cfg.PublishAsyncMaxPending,
		wg:                     &sync.WaitGroup{},
		messageProcessor:       messageProcessor,
	}
}

func (nb *NatsBroker) Start() {
	nc, _ := nats.Connect(nb.url)
	defer nc.Drain()
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(nb.publishAsyncMaxPending))
	cfg := &nats.StreamConfig{
		Name:      nb.streamName,
		Subjects:  nb.getStreamSubjects(),
		Retention: nats.WorkQueuePolicy,
	}
	_, err := js.StreamInfo(cfg.Name)
	if err != nil {
		js.AddStream(cfg)
	}
	consumer := NewNatsConsumer(
		nc,
		cfg.Name,
		js,
		nb.wg,
		[]BusSubject{FileSpecificationGetSubject},
		nb.replaySubjectName,
		nb.messageProcessor,
	)
	consumer.Start()
	nb.wg.Wait()
}

func (nb *NatsBroker) getStreamSubjects() []string {
	return []string{strings.ToLower(nb.streamName) + ".>"}
}
