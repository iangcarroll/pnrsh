package pnr

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	reqEndpoint = `https://www.aeromexico.com/api/v1/app/manage/pnr?store=us&pos=MYB-MOBILE&language=`
)

var (
	requestHeaders = map[string]string{
		"Accept":          "*/*",
		"x-am-app":        "YW0tYXBwLW1vYmlsZQ==",
		"Accept-Language": "en-us",
		"Authorization":   "Atmosphere realm=http://atmosphere,atmosphere_app_id=WorkAndCoApp, atmosphere_signature_method=NONE",
		"User-Agent":      "aeromexico/4187 CFNetwork/1329 Darwin/21.3.0",
	}

	client = http.Client{}
)

func generateAuthHeader(lastName, confirmationCode string) string {
	return fmt.Sprintf(`PNR name="%s",reservation="%s"`, lastName, confirmationCode)
}

func setRequestHeaders(r *http.Request, lastName string, confirmationCode string) {
	for k, v := range requestHeaders {
		r.Header.Set(k, v)
	}

	r.Header.Set(`x-am-user-auth`, generateAuthHeader(lastName, confirmationCode))
}

func sendRequest(lastName, confirmationCode string) ([]byte, error) {
	req, err := http.NewRequest("GET", reqEndpoint, nil)
	if err != nil {
		return []byte{}, err
	}

	setRequestHeaders(req, lastName, confirmationCode)
	res, err := client.Do(req)

	if res.StatusCode != 200 {
		return []byte{}, errors.New("status code was not 200")
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}

func performRequest(lastName, confirmationCode string) (res ManagePnrResponse, err error) {
	data, err := sendRequest(lastName, confirmationCode)
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return res, err
	}

	if res.Meta.Class != "ManagePNRCollection" {
		return res, errors.New("response not successful")
	}

	return res, nil
}

func convertResponse(res ManagePnrResponse) (pnr PNR) {
	convertRemarks(res, &pnr)
	convertFlights(res, &pnr)
	convertPassengers(res, &pnr)
	convertTickets(res, &pnr)
	convertEarnings(res, &pnr)

	return pnr
}
