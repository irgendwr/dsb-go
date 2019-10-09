# dsb-go

[![GoDoc](https://godoc.org/github.com/irgendwr/dsb-go?status.svg)](https://godoc.org/github.com/irgendwr/dsb-go)
[![Travis](https://travis-ci.org/irgendwr/dsb-go.svg)](https://travis-ci.org/irgendwr/dsb-go)

This library lets you access content from [DSBmobile](https://www.dsbmobile.de) in [Go (golang)](https://golang.org).

Unfortunately the heinekingmedia GmbH does not provide any official APIs nor documentation for developers. That sucks! This is why I reverse-engineered their app + website and [developed a php script back in 2016](https://github.com/irgendwr/dsbmobile-php-api/commit/7c34d99ce97818053b158811f91a5af145a7ca74) when I was still in school. I later moved on to using [Go](https://golang.org) and created this package.

Contributions are welcome! Please give this repository a star :star: if it helped you.

## Disclaimer

This is currently not actively maintained since I do not use DSBmobile.

"DSBmobile" and "Digitales Schwarzes Brett" are part of the heinekingmedia GmbH. I am **not** affiliated with them in any way nor do I endorse their products.

## Installation

```bash
go get -u github.com/irgendwr/dsb-go
```

## Example

```go
package main

import (
	"fmt"
	"os"

	"github.com/irgendwr/dsb-go"
)

func main() {
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
```