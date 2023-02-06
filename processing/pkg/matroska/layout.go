package matroska

import (
	"encoding/json"
	"fmt"
	"os"
)

type Filler interface {
	Fill(data []byte, start uint64, size uint64)
}

type JsonWriter interface {
	WriteToJson(out *os.File) error
}

type SpecificationElement struct {
	Size         uint64 `json:"size,omitempty"`
	Position     uint64 `json:"position,omitempty"`
	DataPosition uint64 `json:"dataPosition,omitempty"`
}

func newSpecificationElement() *SpecificationElement {
	return &SpecificationElement{}
}

type Layout struct {
	Specification
}

func NewLayout() *Layout {
	return &Layout{*newSpecification()}
}

func (layout *Layout) Fill(data []byte, start uint64, size uint64) error {
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
			element, hasChildren := getElement(&layout.Specification, idType)
			err = element.fill(data, idType, start+filledSize)
			if err != nil {
				return err
			}
			if hasChildren {
				dataStart := element.DataPosition
				if err = layout.Fill(data, dataStart, element.Size); err != nil {
					return err
				}
			}
			filledSize += element.Size + (element.DataPosition - element.Position)
		}
	}
	return nil
}

func (layout *Layout) WriteToJson(out *os.File) error {
	layoutJson, err := json.MarshalIndent(&layout, "", "    ")
	if err != nil {
		return err
	}
	_, err = out.Write(layoutJson)
	return err
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
