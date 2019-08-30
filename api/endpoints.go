package api

import (
	"bytes"
	"fmt"
	"net/http"
)

var (
	version    = "v1"
	apiURL     = "api2.autopilothq.com"
	mockServer = "private-82d61-autopilot.apiary-mock.com"
	debugProxy = "private-82d61-autopilot.apiary-proxy.com"
	endpoint   = "contact"
)

// GetContact ..
func GetContact(contactID string) (resp *http.Response, err error) {
	url := fmt.Sprintf("https://%s/%s/%s/%s", mockServer, version, endpoint, contactID)
	return apiRequest(url, "GET", []byte{})
}

// UpdateContact ..
func UpdateContact(body []byte) (resp *http.Response, err error) {
	url := fmt.Sprintf("https://%s/%s/%s", mockServer, version, endpoint)
	return apiRequest(url, "PUT", body)
}

// CreateContact ..
func CreateContact(body []byte) (resp *http.Response, err error) {
	url := fmt.Sprintf("https://%s/%s/%s", mockServer, version, endpoint)
	return apiRequest(url, "POST", body)
}

func apiRequest(reqURL, requestType string, body []byte) (*http.Response, error) {
	client := &http.Client{}
	req, err := http.NewRequest(requestType, reqURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}
	setHeaders(req)
	return client.Do(req)
}

func setHeaders(req *http.Request) {
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/71.0.3578.98 Safari/537.36")
	req.Header.Set("content-type", "application/json")
}
