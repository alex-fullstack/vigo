package services

import (
	"encoding/json"
	"fmt"
	"sync"
)

type FileService struct {
	Specification
	mux        *sync.Mutex
	downloader Downloader
}

func NewFileService(downloader Downloader) *FileService {
	return &FileService{
		downloader:    downloader,
		mux:           &sync.Mutex{},
		Specification: *newSpecification(),
	}
}

func (fs *FileService) OnGetFileSpecification(incomingMsg []byte) (replayMsg []byte, err error) {
	message := FileSpecificationGetMessage{}
	err = json.Unmarshal(incomingMsg, &message)
	if err != nil {
		return
	}
	fileData, err := fs.downloader.Download(message.FileId)
	if err != nil {
		return
	}
	fs.mux.Lock()
	err = fs.fill(fileData, 0, uint64(len(fileData)))
	if err == nil {
		replayMessage := FileSpecificationGetReplayMessage{message, fs.Specification}
		replayMsg, err = json.Marshal(replayMessage)
		fs.Specification = *newSpecification()
	}
	fs.mux.Unlock()
	return
}

func (fs *FileService) fill(data []byte, start uint64, size uint64) error {
	var filledSize uint64 = 0
	for filledSize < size {
		if data[start] == byte(CRCIdValue) && filledSize == 0 {
			if crcLength, sizeDataLength, err := decodeEBML(data[start+1:]...); err != nil {
				return err
			} else {
				filledSize += crcLength + uint64(sizeDataLength) + 1
			}
		}
		if idType, err := getEMBLIDType(data[start+filledSize : start+size]); err != nil {
			return err
		} else {
			element, hasChildren := getElement(&fs.Specification, idType)
			err = element.fill(data, idType, start+filledSize)
			if err != nil {
				return err
			}
			if hasChildren {
				dataStart := element.DataPosition
				if err = fs.fill(data, dataStart, element.Size); err != nil {
					return err
				}
			}
			filledSize += element.Size + (element.DataPosition - element.Position)
		}
	}
	return nil
}

func (el *SpecificationElement) fill(data []byte, idType EMBLIDType, position uint64) (err error) {
	el.Position = position
	idSize := getEMBLIDSize(idType)
	dataSizePosition := position + uint64(idSize)
	size, sizeLen, err := decodeEBML(data[dataSizePosition:]...)
	if err != nil {
		return
	}
	el.Size = size
	el.DataPosition = position + uint64(idSize) + uint64(sizeLen)
	return nil
}

func getEMBLIDType(data []byte) (idType EMBLIDType, err error) {
	for level := 0; level < EBMLMaxIDLengthDefault; level++ {
		idValue := EMBLIDValue(decodeUint32(data[:EBMLMaxIDLengthDefault-level]...))
		idType = convertIdValueToIdType(idValue)
		if idType != WrongType {
			return
		}
	}
	err = fmt.Errorf("can't find EMBL type: %d", data[:EBMLMaxIDLengthDefault])
	return
}

func getEMBLIDSize(idType EMBLIDType) uint8 {
	var idValue = convertIdTypeToIdValue(idType)
	return uint8(len(encodeUint(uint64(idValue))))
}
