package pnr

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	reqEndpoint      = `https://www.united.com/api/myTrips/lookup`
	reqBody          = `{"confirmationNumber":"{conf}","lastName":"{lname}","mpUserName":"","region":"US","isDecryptNeeded":false,"encryptedTripDetails":"","partnerCode":null}`
	defaultAuthToken = `xvt5PZjkBYc716K2zwbB97wizAgIw5HBM9L6nx2zxa4qwHrSgJnxqArFSy4f76GHUD+W1/Z4BX3C7/zV+6Fb8qejhpjotgHDdLi19axZDjojnrdA7AsT5C2VYvahyRWhv/2IRTewO7JyDVSYG3e4fXNR0//tS0B3OvbGJXu+4RlM8S7ujCjAzR8x+2chkhir`
)

var (
	requestHeaders = map[string]string{
		"Accept":          "*/*",
		"Content-Type":    "application/json",
		"Accept-Language": "en-US,en;q=0.9",
		"User-Agent":      "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36",
	}

	client = http.Client{}
)

func getAuthToken() string {
	if token := os.Getenv("UNITED_AUTH_TOKEN"); token != "" {
		return token
	}

	return defaultAuthToken
}

func generateReqBody(lastName, confirmationCode string) string {
	body := reqBody
	body = strings.Replace(body, "{conf}", confirmationCode, -1)
	body = strings.Replace(body, "{lname}", lastName, -1)
	return body
}

func setRequestHeaders(r *http.Request) {
	for k, v := range requestHeaders {
		r.Header.Set(k, v)
	}

	r.Header.Set(`x-authorization-api`, `bearer `+getAuthToken())
}

func sendRequest(lastName, confirmationCode string) ([]byte, error) {
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

	if res.ContentMessage.Success != "true" {
		return res, errors.New("response not successful")
	}

	return res, nil
}

func convertResponse(res GetPNRResponse) (pnr PNR) {
	convertFlights(res, &pnr)
	convertPassengers(res, &pnr)
	convertRemarks(res, &pnr)

	return pnr
}
