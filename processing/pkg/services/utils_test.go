package services

import (
	"processing/pkg/infrastructure"
	"testing"
)

func TestDecodeEBMLCorrectly(t *testing.T) {
	actualData := [4][]byte{
		{0x81},
		{0x40, 0x1},
		{0x20, 0x0, 0x1},
		{0x10, 0x0, 0x0, 0x1},
	}
	expectedResult := [4]uint64{1, 1, 1, 1}
	for i := 0; i < len(actualData); i++ {
		actualResult, actualLength, err := infrastructure.decodeEBML(actualData[i]...)

		if err != nil {
			t.Errorf("unexpected result (have error with message: %s, want nil error)", err.Error())
		}
		if actualResult != expectedResult[i] {
			t.Errorf("unexpected result (have decoded result %d, want %d)", actualResult, expectedResult[i])
		}
		if actualLength != i+1 {
			t.Errorf("unexpected result (have length %d, want %d)", actualLength, i+1)
		}
	}

}

func TestDecodeEBMLWithError(t *testing.T) {
	var data = []byte{0}

	_, _, err := infrastructure.decodeEBML(data...)

	if err == nil {
		t.Errorf("unexpected result (have nil error, want error with message: %s)", "can't decode to EMBL")
	}
}
