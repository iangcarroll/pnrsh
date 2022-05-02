package pnr

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	receiptEndpoint = `https://api.delta.com/shim2/receipts`
	receiptReqBody  = `{
	"lastName": "{lname}",
	"firstName": "{fname}",
	"pnr": "{conf}"
}`
)

var (
	receiptRequestHeaders = map[string]string{
		"X-Device-Resolution-Class": "2x",
		"Tlioscnx":                  "Wi-Fi",
		"Accept":                    "application/json",
		"Tliosloc":                  "gn=my_trips:find&ch=find",
		"Tliosid":                   "15.3, iPad Pro (12.9 inch) 3G (Wi-Fi), 5.14",
		"Content-Type":              "application/json; charset=utf-8",
		"response-json":             "true",
		"Accept-Language":           "en-us",
		"Accept-Encoding":           "gzip;q=1.0, compress;q=0.5",
		"User-Agent":                "Fly Delta iPhone, iOS 15.3, 5.14, Phone",
	}
)

func generateReceiptRequestBody(firstName, lastName, confirmationCode string) string {
	body := receiptReqBody
	body = strings.Replace(body, "{fname}", firstName, -1)
	body = strings.Replace(body, "{lname}", lastName, -1)
	body = strings.Replace(body, "{conf}", confirmationCode, -1)
	return body
}

func setReceiptRequestHeaders(r *http.Request) {
	for k, v := range receiptRequestHeaders {
		r.Header.Set(k, v)
	}
}

func sendReceiptRequest(firstName, lastName, confirmationCode string) ([]byte, error) {
	body := generateReceiptRequestBody(firstName, lastName, confirmationCode)

	req, err := http.NewRequest("POST", receiptEndpoint, strings.NewReader(body))
	if err != nil {
		return []byte{}, err
	}

	setReceiptRequestHeaders(req)
	res, err := client.Do(req)

	if res.StatusCode != 200 {
		return []byte{}, errors.New("status code was not 200")
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func performReceiptRequest(firstName, lastName, confirmationCode string) (res ReceiptResponse, err error) {
	data, err := sendReceiptRequest(firstName, lastName, confirmationCode)
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return res, err
	}

	if res.ReceiptType == "" {
		return res, errors.New("response not successful")
	}

	return res, nil
}
