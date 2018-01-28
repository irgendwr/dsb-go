# dsb-go
> This library lets you access content from DSBmobile in golang.

[![GoDoc](https://godoc.org/github.com/irgendwr/dsb-go?status.svg)](https://godoc.org/github.com/irgendwr/dsb-go)
[![Travis](https://travis-ci.org/irgendwr/dsb-go.svg)](https://travis-ci.org/irgendwr/dsb-go)

```bash
go get -u github.com/irgendwr/dsb-go
```

## Example
```go
package main

import (
	"fmt"

	"github.com/irgendwr/dsb-go"
)

func main() {
    account := dsb.NewAccount("123456", "password")
    content, err := account.GetContent()

    if err != nil {
        fmt.Printf("Error: %s", err)
    }

    timetables, err := content.GetTimetables()

    if err != nil {
        fmt.Printf("Error: %s", err)
    }

    timetables[0].GetURL()
    fmt.Printf("URL: %s", url)
}
```