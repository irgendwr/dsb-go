package dsb_test

import (
	"log"
	"os"
	"testing"

	"github.com/irgendwr/dsb-go"
)

var data *dsb.Response

func TestMain(m *testing.M) {
	account := dsb.NewAccount("196041", "DVfSiW")

	if response, err := account.GetData(); err != nil {
		log.Fatalf("Unable to get data: %s\n", err)
	} else {
		data = response
		os.Exit(m.Run())
	}
}

func TestGetTimetables(t *testing.T) {
	content := data.GetContent()
	if content == nil {
		t.Error("Unable to get content")
	}

	timetables := content.GetTimetables()
	if l := len(timetables); l == 0 {
		t.Error("No timetables found")
	} else {
		t.Logf("%d timetable(s) found", l)
		for i, timetable := range timetables {
			t.Logf("URL #%d: %s", i, timetable.GetURL())
		}
	}
}
