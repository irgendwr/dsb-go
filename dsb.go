// Package dsb lets you access content from DSBmobile in golang.
//
// Example:
//
//		import (
//			"fmt"
//
//			"github.com/irgendwr/dsb-go"
//		)
//
//		func main() {
//			account := dsb.NewAccount("123456", "password")
//			content, err := account.GetContent()
//
//			if err != nil {
//				fmt.Printf("Error: %s", err)
//			}
//
//			timetables, err := content.GetTimetables()
//
//			if err != nil {
//				fmt.Printf("Error: %s", err)
//			}
//
//			timetables[0].GetURL()
//			fmt.Printf("URL: %s", url)
//		}
//
package dsb

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pkg/errors"
)

// constants
const (
	bundleID   = "de.heinekingmedia.inhouse.dsbmobile.web"
	webservice = "https://mobile.dsbcontrol.de/JsonHandlerWeb.ashx/GetData"
	appVersion = "2.3"
	lang       = "de"
	device     = "WebApp"
	success    = 0
)

// request types
const (
	unknownRequest = iota
	dataRequest
	mailRequest
	feedbackRequest
	subjectsRequest
)

// NewAccount creates a new account interface
func NewAccount(username string, password string) Account {
	return Account{username, password}
}

// GetData returns all available information of the account
func (account *Account) GetData() (*Response, error) {
	JSONdata, err := json.Marshal(map[string]interface{}{
		"UserId":     account.username,
		"UserPw":     account.password,
		"Abos":       []string{},
		"AppVersion": appVersion,
		"Language":   lang,
		"OsVersion":  "",
		"AppId":      "",
		"Device":     device,
		"PushId":     "",
		"BundleId":   bundleID,
		"Date":       "",
		"LastUpdate": "",
	})
	if err != nil {
		return nil, errors.Wrap(err, "json encode failed")
	}

	// gzip encode
	var gzipBuffer bytes.Buffer
	gzipWriter := gzip.NewWriter(&gzipBuffer)

	if _, err = gzipWriter.Write(JSONdata); err != nil {
		return nil, errors.Wrap(err, "gzip encode failed")
	}
	gzipWriter.Close()

	// base64 encode
	base64data := base64.StdEncoding.EncodeToString(gzipBuffer.Bytes())

	JSONdata, err = json.Marshal(map[string]interface{}{
		"req": map[string]interface{}{
			"Data":     base64data,
			"DataType": dataRequest,
		},
	})
	if err != nil {
		return nil, errors.Wrap(err, "2nd json encode failed")
	}

	// build request
	req, err := http.NewRequest("POST", webservice, bytes.NewReader(JSONdata))
	if err != nil {
		return nil, errors.Wrap(err, "could not create request")
	}

	// set headers
	req.Header.Add("bundle_id", bundleID)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Referer", "https://www.dsbmobile.de/default.aspx")

	// TODO: don't use default client
	// send request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "request failed")
	}
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed, status: %s (%d)", res.Status, res.StatusCode)
	}

	// decode json
	var data struct {
		Data string `json:"d"`
	}
	if err := json.NewDecoder(res.Body).Decode(&data); err != nil {
		return nil, errors.Wrap(err, "json unmarshal failed")
	}

	if data.Data == "" {
		return nil, errors.New("got empty response")
	}

	// base64 decode
	decodedBase64, err := base64.StdEncoding.DecodeString(data.Data)
	if err != nil {
		return nil, errors.Wrap(err, "base64 decode failed")
	}

	// gzip decode
	gzipReader, err := gzip.NewReader(bytes.NewReader(decodedBase64))
	if err != nil {
		return nil, errors.Wrap(err, "gzip reader failed")
	}
	defer gzipReader.Close()

	// decode json
	var response Response
	if err := json.NewDecoder(gzipReader).Decode(&response); err != nil {
		return nil, errors.Wrap(err, "json unmarshal of response failed")
	}

	if response.Resultcode != success {
		return nil, fmt.Errorf("Unknown Resultcode: %d", response.Resultcode)
	}

	return &response, nil
}

// GetContent returns the MenuItem containing the main content
func (data *Response) GetContent() (*MenuItem, error) {
	for _, menuItem := range data.MenuItems {
		// TODO: use Method or Index
		if menuItem.Title == "Inhalte" {
			return &menuItem, nil
		}
	}
	return nil, errors.New("content not found")
}

// GetContent returns the MenuItem containing the main content
func (account *Account) GetContent() (*MenuItem, error) {
	data, err := account.GetData()

	if err != nil {
		return nil, err
	}

	return data.GetContent()
}

// GetTimetables returns all timetables,
// use GetContent() to get the according menuItem first
func (menuItem *MenuItem) GetTimetables() ([]MenuItemChildItem, error) {
	for _, menuItemChild := range menuItem.Childs {
		if menuItemChild.Method == "timetable" {
			return menuItemChild.Root.Childs, nil
		}
	}
	return []MenuItemChildItem{}, errors.New("no timetables found")
}

// GetURL returns the url of specified content
func (menuItemChildItem *MenuItemChildItem) GetURL() string {
	return menuItemChildItem.Childs[0].Detail
}

// TODO: implement more getters
