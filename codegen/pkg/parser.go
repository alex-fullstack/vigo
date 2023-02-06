package pkg

import (
	"encoding/xml"
	"strconv"
)

func Parse(d *xml.Decoder) []XMLElement {
	elements := make([]XMLElement, 0)
	for t, _ := d.Token(); t != nil; t, _ = d.Token() {
		switch se := t.(type) {
		case xml.StartElement:
			name, id, level, multiple := "", "", -1, false
			for _, attr := range se.Attr {
				switch attr.Name.Local {
				case "name":
					name = attr.Value
				case "level":
					level, _ = strconv.Atoi(attr.Value)
				case "id":
					id = attr.Value
				case "multiple":
					multiple = attr.Value == "1"
				}
			}
			if id != "" {
				elements = append(elements, XMLElement{name, level, id, multiple})
			}
		}
	}
	return elements
}
