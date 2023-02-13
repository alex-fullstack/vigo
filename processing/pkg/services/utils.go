package services

import (
	"encoding/binary"
	"fmt"
	"math/bits"
)

func pow2(exp uint8) uint64 {
	if exp == 1 {
		return 2
	}
	return 2 * pow2(exp-1)
}

func decodeUint32(data ...byte) uint32 {
	input := make([]byte, 4)
	start := 4 - len(data)
	for i := start; i < 4; i++ {
		input[i] = data[i-start]
	}
	return binary.BigEndian.Uint32(input)
}

func decodeUint64(data ...byte) uint64 {
	input := make([]byte, 8)
	start := 8 - len(data)
	for i := start; i < 8; i++ {
		input[i] = data[i-start]
	}
	return binary.BigEndian.Uint64(input)
}

func decodeEBML(data ...byte) (result uint64, length int, err error) {
	for i := 0; i < 8 && i < len(data); i++ {
		chunk := data[:i+1]
		result = decodeUint64(chunk...)
		shift := 7 * (i + 1)
		if result>>shift == 1 {
			result -= pow2(uint8(shift))
			length = i + 1
			return
		}
	}
	err = fmt.Errorf("can't decode to EMBL")
	return
}
func encodeUint(value uint64) []byte {
	buf := make([]byte, 8)
	binary.BigEndian.PutUint64(buf, value)
	return buf[bits.LeadingZeros64(value)>>3:]
}
