package pnr


func Retrieve(lastName, confirmationCode string) (PNR, error) {

	res, raw, err := performRequest(lastName, confirmationCode)
	if err != nil {
		return PNR{}, err
	}

 	return PNR{Data:res, RawResponse: string(raw)}, nil
}
