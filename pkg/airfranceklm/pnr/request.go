package pnr

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func sendRequest(confirmationCode string, firstName string, lastName string) ([]byte, error) {
	request, err := http.NewRequest("GET", generatePnrUrl(confirmationCode, firstName, lastName), nil)
	if err != nil {
		return []byte{}, err
	}

	token, err := getAuthToken()
	if err != nil {
		return []byte{}, err
	}

	request.Header.Set("Authorization", "Bearer "+token)
	for key, value := range pnrRequestHeaders {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil || response == nil || response.StatusCode != 200 {
		if response != nil {
			log.Println("Status code from AFKL: ", response.StatusCode)
		}
		if err != nil {
			log.Println("Error from AFKL: ", err)
		}

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return []byte{}, errors.New("Invalid response and could not read body")
		}

		log.Println(string(body[:]))

		return []byte{}, errors.New("Invalid response from endpoint.")
	}

	defer response.Body.Close()
	return ioutil.ReadAll(response.Body)
}

func performRequest(confirmationCode string, firstName string, lastName string) (res GetPNRResponse, err error) {
	data, err := sendRequest(confirmationCode, firstName, lastName)
	if err != nil {
		return res, err
	}

	if err := json.Unmarshal(data, &res); err != nil {
		return res, err
	}

	return res, nil
}

func generatePnrUrl(confirmationCode string, firstName string, lastName string) string {
	url := pnrUrl
	url = strings.Replace(url, "{confirmationCode}", confirmationCode, -1)
	url = strings.Replace(url, "{firstName}", firstName, -1)
	url = strings.Replace(url, "{lastName}", lastName, -1)
	return url
}
