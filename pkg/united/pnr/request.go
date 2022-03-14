package pnr

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	reqEndpoint = `https://smartphone.united.com/UnitedMobileDataServices/api/ManageReservation/GetPNRByRecordLocator`
	reqBody     = `{"deviceId":"E816208D-B7C5-5367-99A8-48D893371813","isOTFConversion":false,"mileagePlusNumber":"","recordLocator":"{conf}","languageCode":"en-US","sessionId":"","accessCode":"ACCESSCODE","application":{"id":1,"name":"iOS","isProduction":false,"version":{"minor":"4.1.44","major":"4.1.44","displayText":"","build":""}},"transactionId":"E816208D-B7C5-5367-99A8-48D893371813|F7C7F1DA-E289-4BBF-AC85-717FE32D204D","flow":"VIEWRES","lastName":"{lname}"}`
)

var (
	requestHeaders = map[string]string{
		"Accept":          "*/*",
		"Content-Type":    "application/json",
		"Accept-Language": "en-US,en;q=0.9",
		"User-Agent":      "UnitedCustomerFacingIPhone/4.1.44.1 CFNetwork/1329 Darwin/21.3.0",
	}

	client = http.Client{}
)

func generateReqBody(lastName string, confirmationCode string) string {
	body := reqBody
	body = strings.Replace(body, "{conf}", confirmationCode, -1)
	body = strings.Replace(body, "{lname}", lastName, -1)
	return body
}

func setRequestHeaders(r *http.Request) {
	for k, v := range requestHeaders {
		r.Header.Set(k, v)
	}
}

func sendRequest(lastName string, confirmationCode string) ([]byte, error) {
	req, err := http.NewRequest("POST", reqEndpoint, strings.NewReader(generateReqBody(lastName, confirmationCode)))
	if err != nil {
		return []byte{}, err
	}

	setRequestHeaders(req)
	res, err := client.Do(req)

	if res.StatusCode != 200 {
		return []byte{}, errors.New("status code was not 200")
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func performRequest(lastName string, confirmationCode string) (res GetPNRResponse, err error) {
	data, err := sendRequest(lastName, confirmationCode)
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return res, err
	}

	if res.Exception != nil {
		return res, errors.New("response not successful")
	}

	return res, nil
}

func convertResponse(res GetPNRResponse) (pnr PNR) {
	convertFlights(res, &pnr)
	convertPassengers(res, &pnr)

	return pnr
}
