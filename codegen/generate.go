package main

import (
	"bytes"
	"encoding/xml"
	"matroska-spec-generator/pkg"
	"os"
)

func main() {
	fileData, _ := os.ReadFile(os.Args[1])
	d := xml.NewDecoder(bytes.NewReader(fileData))
	elements := pkg.Parse(d)
	out, _ := os.Create(os.Args[2])
	pkg.Generate(out, elements)
}
