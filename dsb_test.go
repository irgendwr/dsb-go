package dsb

import (
	"log"
	"os"
	"testing"
)

var data *Response

func TestMain(m *testing.M) {
	account := NewAccount("196041", "DVfSiW")
	var err error
	data, err = account.GetData()
	if err != nil {
		log.Fatalf("Error: %s\n", err)
		os.Exit(1)
	}
	os.Exit(m.Run())
}

func TestGetTimetables(t *testing.T) {
	content := data.GetContent()
	if content == nil {
		t.Error("Unable to get content")
	}

	timetables := content.GetTimetables()
	if tCount := len(timetables); tCount == 0 {
		t.Error("No timetables found")
	} else {
		t.Logf("%d timetable(s) found", tCount)
	}
}
