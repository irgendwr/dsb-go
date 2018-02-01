package dsb_test

import (
	"fmt"

	dsb "github.com/irgendwr/dsb-go"
)

func Example() {
	account := dsb.NewAccount("123456", "password")
	content, err := account.GetContent()
	if err != nil {
		fmt.Printf("could not get account content: %s", err)
	}

	timetables, err := content.GetTimetables()
	if err != nil {
		fmt.Printf("could not get timetable: %s", err)
	}

	url := timetables[0].GetURL()
	fmt.Printf("URL: %s", url)
}
