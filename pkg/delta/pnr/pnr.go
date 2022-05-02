package pnr

func Retrieve(apiServer, firstName, lastName, confirmationCode string) (PNR, error) {
	res, err := performRequest(apiServer, firstName, lastName, confirmationCode)
	if err != nil {
		return PNR{}, err
	}

	return convertResponse(res, firstName, lastName, confirmationCode), nil
}
