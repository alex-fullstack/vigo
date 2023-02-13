package services

type BusMessage struct {
	Id string
}

type FileSpecificationGetMessage struct {
	BusMessage
	FileId string
}

type FileSpecificationGetReplayMessage struct {
	FileSpecificationGetMessage
	Specification Specification
}

type SpecificationElement struct {
	Size         uint64 `json:"size,omitempty"`
	Position     uint64 `json:"position,omitempty"`
	DataPosition uint64 `json:"dataPosition,omitempty"`
}

func newSpecificationElement() *SpecificationElement {
	return &SpecificationElement{}
}

type Downloader interface {
	Download(fileId string) ([]byte, error)
}

type MessageProcessor interface {
	OnGetFileSpecification(incomingMsg []byte) (replayMsg []byte, err error)
}
