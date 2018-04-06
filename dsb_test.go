package dsb_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	dsb "."
)

var data *dsb.Response

func TestMain(m *testing.M) {
	account := dsb.NewAccount("196041", "DVfSiW")
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

func Example() {
	account := dsb.NewAccount("user", "password")
	content, err := account.GetContent()
	if err != nil {
		fmt.Printf("could not get account content: %s", err)
	}

	timetables := content.GetTimetables()
	if len(timetables) != 0 {
		fmt.Println("no timetables found")
		return
	}

	url := timetables[0].GetURL()
	fmt.Printf("URL: %s", url)
}
