package pnr

import (
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	reqEndpoint = `https://www.united.com/api/myTrips/lookup`
	reqBody     = `{"confirmationNumber":"{conf}","lastName":"{lname}","mpUserName":"","region":"US","isDecryptNeeded":false,"encryptedTripDetails":"","partnerCode":null}`

	userAgent = `Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/101.0.4951.54 Safari/537.36`
)

var (
	requestHeaders = map[string]string{
		"Accept":          "application/json",
		"Accept-Language": "en-US",
		"Content-Type":    "application/json",
		"DNT":             "1",
		"Origin":          "https://www.united.com",
		"Referer":         "https://www.united.com/en/us/awardaccelerator",
		"User-Agent":      userAgent,
	}

	client = http.Client{
		Timeout: time.Second * 15,

		Transport: &http.Transport{
			// When using an HTTP proxy, we need to disable TLS verification.
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: os.Getenv("HTTP_PROXY") != "",
			},

			Proxy: func(r *http.Request) (*url.URL, error) {
				if p := os.Getenv("HTTP_PROXY"); p != "" {
					return url.Parse(p)
				}

				return nil, nil
			},
		},
	}
)

func generateReqBody(lastName, confirmationCode string) string {
	body := reqBody
	body = strings.Replace(body, "{conf}", confirmationCode, -1)
	body = strings.Replace(body, "{lname}", lastName, -1)
	return body
}

func setRequestHeaders(r *http.Request, token string) {
	for k, v := range requestHeaders {
		r.Header.Set(k, v)
	}

	r.Header.Set(`x-authorization-api`, `bearer `+token)
}

func sendRequest(lastName, confirmationCode string) ([]byte, error) {
	req, err := http.NewRequest("POST", reqEndpoint, strings.NewReader(generateReqBody(lastName, confirmationCode)))
	if err != nil {
		return []byte{}, err
	}

	token, err := getAuthToken()
	if err != nil {
		return []byte{}, err
	}

	setRequestHeaders(req, token)
	res, err := client.Do(req)

	if res.StatusCode != 200 {
		return []byte{}, errors.New(fmt.Sprintf("status code was %d", res.StatusCode))
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
	convertTickets(res, &pnr)
	convertSsrs(res, &pnr)

	return pnr
}
