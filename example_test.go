package dsb_test

import (
	"fmt"
	"os"

	"github.com/irgendwr/dsb-go"
)

func Example() {
	// set credentials
	account := dsb.NewAccount("user", "password")
	// get content
	content, err := account.GetContent()

	if err != nil {
		// exit on error
		fmt.Printf("Error: %s", err)
		os.Exit(1)
	}

	// get timetables
	timetables := content.GetTimetables()
	if len(timetables) == 0 {
		fmt.Println("no timetables found")
		return
	}

	// print all url's
	for i, timetable := range timetables {
		fmt.Printf("URL #%d: %s", i, timetable.GetURL())
	}
}
