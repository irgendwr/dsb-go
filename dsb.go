// Package dsb lets you access content from DSBmobile in golang.
package dsb

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

// constants
const (
	bundleID   = "de.heinekingmedia.dsbmobile"
	webservice = "https://www.dsbmobile.de/JsonHandler.ashx/GetData"
	appVersion = "2.5.9"
	lang       = "de"
	device     = "Nexus 4"
	osVersion  = "27 8.1.0"
	success    = 0
)

// request types
const (
	UnknownRequest = iota
	DataRequest
	MailRequest
	FeedbackRequest
	SubjectsRequest
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
		"OsVersion":  osVersion,
		"AppId":      uuid.New().String(),
		"Device":     device,
		"PushId":     "",
		"BundleId":   bundleID,
		"Date":       time.Now(),
		"LastUpdate": time.Now(),
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
			"DataType": DataRequest,
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
	req.Header.Set("User-Agent", "dsb-go")

	// send request
	httpClient := &http.Client{
		Timeout: time.Second * 10,
	}
	res, err := httpClient.Do(req)
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

	// for debugging
	//jsonStr, _ := ioutil.ReadAll(gzipReader)
	//fmt.Printf("%s\n\n", jsonStr)

	// decode json
	var response Response
	if err := json.NewDecoder(gzipReader).Decode(&response); err != nil {
		return nil, errors.Wrap(err, "json unmarshal of response failed")
	}

	if response.StatusCode != success {
		return nil, fmt.Errorf("API request failed: %s (%d)", response.Status, response.StatusCode)
	}

	return &response, nil
}

// GetCategoryByIndex returns the category with the given index
func (data *Response) GetCategoryByIndex(index int) *Category {
	for _, category := range data.Categorys {
		if category.Index == index {
			return &category
		}
	}
	return nil
}

// GetCategoryByTitle returns the category with the given name
func (data *Response) GetCategoryByTitle(title string) *Category {
	for _, category := range data.Categorys {
		if category.Title == title {
			return &category
		}
	}
	return nil
}

// GetContent returns the category containing the main content
func (account *Account) GetContent() (*Category, error) {
	data, err := account.GetData()
	if err != nil {
		return nil, err
	}

	return data.GetContent(), nil
}

// GetContent returns the category containing the main content
func (data *Response) GetContent() *Category {
	return data.GetCategoryByIndex(0)
}

// GetMenuByMethod returns the menu with the given method
func (category *Category) GetMenuByMethod(method string) *Menu {
	for _, menu := range category.Menus {
		if menu.Method == method {
			return &menu
		}
	}
	return nil
}

// GetTimetables returns all timetables
func (category *Category) GetTimetables() []MenuItem {
	timetables := category.GetMenuByMethod("timetable")
	if timetables != nil {
		return timetables.Root.Childs
	}
	return []MenuItem{}
}

// GetNews returns all news
func (category *Category) GetNews() []MenuItem {
	news := category.GetMenuByMethod("news")
	if news != nil {
		return news.Root.Childs
	}
	return []MenuItem{}
}

// GetTiles returns all news
func (category *Category) GetTiles() *Menu {
	return category.GetMenuByMethod("tiles")
}

// GetURL returns the URL of a timetable
func (menuItem *MenuItem) GetURL() string {
	return menuItem.Childs[0].Detail
}
