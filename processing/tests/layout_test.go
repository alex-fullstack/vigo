package tests

import (
	"encoding/json"
	"os"
	"processing/pkg/matroska"
	"strconv"
	"testing"
)

func TestFill_CorrectData_WorkCorrectly(t *testing.T) {
	for i := 0; i < 6; i++ {
		if i == 3 {
			continue
		}
		fileData, _ := os.ReadFile("./fixtures/test" + strconv.Itoa(i+1) + ".mkv")
		jsonData, _ := os.ReadFile("./fixtures/layout" + strconv.Itoa(i+1) + ".json")
		expectedJson := string(jsonData)
		layout := matroska.NewLayout()

		layout.Fill(fileData, 0, uint64(len(fileData)))
		actualJson, _ := json.MarshalIndent(layout, "", "    ")

		if expectedJson != string(actualJson) {
			t.Errorf("unexpected result (have %s, want %s)", actualJson, expectedJson)
		}
	}
}

// TODO: Добавить обработку данных 0x0a
func TestFill_WrongData_ReturnError(t *testing.T) {
	fileData, _ := os.ReadFile("./fixtures/test4.mkv")
	layout := matroska.NewLayout()
	expectedError := "can't find EMBL type: [10 10 10 10]"

	err := layout.Fill(fileData, 0, uint64(len(fileData)))

	if err == nil || err.Error() != expectedError {
		t.Errorf("unexpected result (have nil, want %s)", expectedError)
	}
}
