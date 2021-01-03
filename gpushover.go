// Package gpushover is a Go wrapper for Pushover's notification API.
//
// Copyright (c) 2020 José Manuel Díez. <j.diezlopez@protonmail.ch>
// Copyright (c) 2020 Gridfinity, LLC. <admin@gridfinity.com>
// Copyright (c) 2014 Damian Gryski. <damian@gryski.com>
// Copyright (c) 2020 Jeffrey H. Johnson. <jeff@gridfinity.com>
// Copyright (c) 2014 Adam Lazzarato.
//
// All Rights reserved.
//
// All use of this code is governed by the MIT license.
// The complete license is available in the LICENSE file.
package gpushover // import "go.gridfinity.dev/gpushover"

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	json "github.com/json-iterator/go"
	gpushoverLegal "go4.org/legal"
)

const endpoint string = "https://api.pushover.net/1/messages.json"

// ErrorPushover is the generic Pushover error
var ErrorPushover = errors.New("ErrorPushover")

// P defines the Pushover configuration
type P struct {
	UserKey, AppKey string
	Client          *http.Client
}

// Response defines the Pushover API response
type Response struct {
	Status  int
	Errors  []interface{}
	Message string
}

// Notification defines the Pushover API request
type Notification struct {
	Message, Title, URL, URLTitle, Sound, Device, Callback string
	Timestamp                                              time.Time
	Priority, Retry, Expire                                int
}

func (n Notification) toValues(p P) url.Values {
	return url.Values{
		"user":      {p.UserKey},
		"token":     {p.AppKey},
		"message":   {n.Message},
		"title":     {n.Title},
		"url":       {n.URL},
		"url_title": {n.URLTitle},
		"sound":     {n.Sound},
		"device":    {n.Device},
		"timestamp": {fmt.Sprintf("%d", n.Timestamp.Unix())},
		"priority":  {fmt.Sprintf("%d", n.Priority)},
		"retry":     {fmt.Sprintf("%d", n.Retry)},
		"expire":    {fmt.Sprintf("%d", n.Expire)},
		"callback":  {n.Callback},
	}
}

// Notify sends the Pushover notification
func (p P) Notify(n Notification) (*Response, error) {
	client := p.Client
	if client == nil {
		client = http.DefaultClient
	}

	resp, err := client.PostForm(endpoint, n.toValues(p))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode == 200 {
		return nil, nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	response := new(Response)
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}
	return response, ErrorPushover
}

func init() {
	// Register licensing
	gpushoverLegal.RegisterLicense(
		"\nCopyright (c) 2020 José Manuel Díez. <j.diezlopez@protonmail.ch>\nCopyright (c) 2020 Gridfinity, LLC. <admin@gridfinity.com>\nCopyright (c) 2014 Damian Gryski. <damian@gryski.com>\nCopyright (c) 2020 Jeffrey H. Johnson. <jeff@gridfinity.com>\nCopyright (c) 2014 Adam Lazzarato.\n\nPermission is hereby granted, free of charge, to any person obtaining a\ncopy of this software and associated documentation files (the \"Software\"),\nto deal in the Software without restriction, including, without limitation,\nthe rights to use, copy, modify, merge, publish, distribute, sub-license,\nand/or sell copies of the Software, and to permit persons to whom the\nSoftware is furnished to do so, subject to the following conditions:\n\nThe above copyright notice and this permission notice shall be\nincluded in all copies or substantial portions of the Software.\n\nTHE SOFTWARE IS PROVIDED \"AS IS\", WITHOUT WARRANTY OF ANY KIND, EXPRESS\nOR IMPLIED, INCLUDING, BUT NOT LIMITED, TO THE WARRANTIES OF\nMERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE, AND NON-INFRINGEMENT.\nIN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM,\nDAMAGES, OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT, OR\nOTHERWISE, ARISING FROM, OUT OF, OR IN CONNECTION WITH, THE SOFTWARE, OR\nTHE USE OR OTHER DEALINGS IN THE SOFTWARE.\n",
	)
}
