package pnr

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func getAuthToken() (string, error) {

	request, err := http.NewRequest("POST", authUrl, strings.NewReader(`{"grant_type":"client_credentials"}`))
	if err != nil {
		return "", errors.New("Could not build token request.")
	}

	for key, value := range tokenRequestHeaders {
		request.Header.Set(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		return "", errors.New("Error performing token request")
	}

	if response.StatusCode != 200 {

		body, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return "", errors.New("Token status code not 200 and could not read body")
		}

		log.Println(string(body[:]))

		return "", errors.New("Token status code not 200")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", errors.New("Could not read body")
	}

	var parsedResponse TokenResponse
	return parsedResponse.AccessToken, json.Unmarshal(body, &parsedResponse)
}
