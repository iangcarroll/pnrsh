package pnr

import "log"

func Retrieve(lastName, confirmationCode string) (PNR, error) {
	res, err := performRequest(lastName, confirmationCode)
	if err != nil {
		log.Println(err)
		return PNR{}, err
	}

	return convertResponse(res), nil
}
