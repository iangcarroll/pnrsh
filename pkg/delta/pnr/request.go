package pnr

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	DeltaEndpoint          = `https://api.delta.com/api2/mobile/getPnr`
	VirginAtlanticEndpoint = `https://api.virginatlantic.com/api2/mobile/getPnr`
	reqBody                = `<retrievePnrRequest><confirmationNumber>{conf}</confirmationNumber><disableMerchandiseSearch>false</disableMerchandiseSearch><firstName>{fname}</firstName><isCheckForUpsell>Y</isCheckForUpsell><lastName>{lname}</lastName><requestInfo><appId>mobile</appId><applicationVersion>5.14</applicationVersion><buildNumber>19902</buildNumber><channel></channel><channelId>mobile</channelId><deviceName>IPHONE</deviceName><deviceType>2x</deviceType><osName>iOS</osName><osVersion>15.3</osVersion><responseType>xml</responseType><transactionId></transactionId></requestInfo><retrieveBy>confirmationNumber</retrieveBy><vacationsSearch>true</vacationsSearch></retrievePnrRequest>`
)

var (
	requestHeaders = map[string]string{
		"X-Device-Resolution-Class": "2x",
		"Tlioscnx":                  "Wi-Fi",
		"Accept":                    "application/xml",
		"Tliosloc":                  "gn=my_trips:find&ch=find",
		"Tliosid":                   "15.3, iPad Pro (12.9 inch) 3G (Wi-Fi), 5.14",
		"Content-Type":              "application/xml; charset=utf-8",
		"response-xml":              "true",
		"Accept-Language":           "en-us",
		"Accept-Encoding":           "gzip;q=1.0, compress;q=0.5",
		"User-Agent":                "Fly Delta iPhone, iOS 15.3, 5.14, Phone",
	}

	client = http.Client{
		Timeout: time.Second * 15,

		Transport: &http.Transport{
			Proxy: func(r *http.Request) (*url.URL, error) {
				if p := os.Getenv("HTTP_PROXY"); p != "" {
					return url.Parse(p)
				}

				return nil, nil
			},
		},
	}
)

func generateRequestBody(firstName, lastName, confirmationCode string) string {
	body := reqBody
	body = strings.Replace(body, "{fname}", firstName, -1)
	body = strings.Replace(body, "{lname}", lastName, -1)
	body = strings.Replace(body, "{conf}", confirmationCode, -1)
	return body
}

func setRequestHeaders(r *http.Request) {
	for k, v := range requestHeaders {
		r.Header.Set(k, v)
	}
}

func sendRequest(apiServer, firstName, lastName, confirmationCode string) ([]byte, error) {
	body := generateRequestBody(firstName, lastName, confirmationCode)

	req, err := http.NewRequest("POST", apiServer, strings.NewReader(body))
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

func performRequest(apiServer, firstName, lastName, confirmationCode string) (res RetrievePnrResponse, err error) {
	data, err := sendRequest(apiServer, firstName, lastName, confirmationCode)
	if err != nil {
		return res, err
	}

	if err := xml.Unmarshal(data, &res); err != nil {
		return res, err
	}

	if res.Status != "SUCCESS" {
		return res, errors.New("response not successful")
	}

	return res, nil
}

func convertResponse(res RetrievePnrResponse, firstName, lastName, confirmationCode string) (pnr PNR) {
	var wg sync.WaitGroup
	wg.Add(1)

	// Kick off the receipt request async while we do the other processing.
	go func() {
		defer wg.Done()

		receiptRes, err := performReceiptRequest(firstName, lastName, confirmationCode)
		if err == nil {
			res.receiptResponse = &receiptRes
		} else {
			log.Println("Got an error retrieving receipts")
		}
	}()

	convertRemarks(res, &pnr)
	convertFlights(res, &pnr)
	convertPassengers(res, &pnr)
	convertFlags(res, &pnr)
	convertTickets(res, &pnr)
	convertFare(res, &pnr)

	// Wait for the receipt request and process it if we succeeded.
	wg.Wait()
	if res.receiptResponse != nil {
		convertReceiptResponse(res, &pnr)
	}

	return pnr
}
