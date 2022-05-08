package pnr

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sync/atomic"
)

var (
	tokenErrorCount     uint64
	tokenErrorThreshold = uint64(2)
)

func getAuthToken() (string, error) {
	req, err := http.NewRequest(`GET`, `https://www.united.com/api/token/anonymous`, nil)
	if err != nil {
		return "", errors.New("could not build token request")
	}

	req.Header.Set(`User-Agent`, userAgent)
	res, err := client.Do(req)

	if err != nil {
		return "", errors.New("error performing token request")
	}

	if res.StatusCode != 200 {
		if atomic.AddUint64(&tokenErrorCount, 1) > tokenErrorThreshold {
			log.Println("Restarting because we cannot get any more tokens")
			os.Exit(1)
		}

		return "", errors.New("token status code not 200")
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", errors.New("could not read body")
	}

	var parsedResponse TokenResponse
	return parsedResponse.Data.Token.Hash, json.Unmarshal(body, &parsedResponse)
}
