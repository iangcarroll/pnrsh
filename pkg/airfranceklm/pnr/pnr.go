package pnr

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

const (
	authUrl   = "https://api-mobile.airfranceklm.com/mobile/cid/token?client_id=w53y7kv97nbcdmgmkr9watkj"
	pnrUrl    = "https://api-mobile.airfranceklm.com/mobile/v3/reservations/?bookingCode={confirmationCode}&firstName={firstName}&lastName={lastName}"
	userAgent = "KL-STR/13.9.2 (com.klm.mobile.iphone.klmmobile; build:2023.08.04.111; iOS 16.6.0) Alamofire/5.7.1"
)

var (
	tokenRequestHeaders = map[string]string{
		"Accept":        "*/*",
		"Content-Type":  "application/json",
		"User-Agent":    userAgent,
		"Authorization": "Basic dzUzeTdrdjk3bmJjZG1nbWtyOXdhdGtqOkdNWDFwaWtPNzR8QjBENENDMEUtRkE5NS00NzY2LUFDMTMtNTQxMTA1NzczREJDOzMwRTgwRjkxLUQ5NzMtNDVFQi05OUZDLTNERjk1QkY2MzZDMnx2WnR5YmloMFVLcjhRVGJFSitBTEVSSjJhRHlUVEtrbU5ERVU0VE8vMlNRPQ==",
	}

	pnrRequestHeaders = map[string]string{
		"Accept":              "*/*",
		"AFKL-TRAVEL-Host":    "KL",
		"AFKL-TRAVEL-Country": "us",
		"User-Agent":          userAgent,
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

func Retrieve(confirmationCode string, firstName string, lastName string) (PNR, error) {
	res, err := performRequest(confirmationCode, firstName, lastName)
	if err != nil {
		log.Println(err)
		return PNR{}, err
	}

	return convertResponse(res), nil
}
