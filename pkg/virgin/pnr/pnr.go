package pnr

func Retrieve(firstName string, lastName string, confirmationCode string) (PNR, error) {
	res, err := performRequest(firstName, lastName, confirmationCode)
	if err != nil {
		return PNR{}, err
	}

	return convertResponse(res), nil
}
