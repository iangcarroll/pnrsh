package pnr

import (
	"github.com/aws/aws-sdk-go/aws"
    "github.com/aws/aws-sdk-go/aws/credentials"
    "github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
    "github.com/aws/aws-sdk-go/service/cognitoidentity"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
	"strconv"
	"io"
)

const (
	graphqlQuery = `query RetrievePnr(\n  $language: String!\n  $bookingReferenceID: String!\n  $surname: String!\n  $context: String\n) {\n  retrievePNR(\n    language: $language\n    bookingReferenceID: $bookingReferenceID\n    surname: $surname\n    context: $context\n  ) {\n    __typename\n    bookingInfo {\n      __typename\n      language\n      currency\n      pointOfSale\n      bookingReference\n      source\n      tripType\n      employeeBooking\n      lastModified\n      isDisrupted\n      acoManageUrl\n      acBidUpgrade {\n        __typename\n        isEligible\n        isModify\n        url\n      }\n      eUpgrade {\n        __typename\n        isEligible\n      }\n      mixedFare\n      seatsInfo {\n        __typename\n        allSeatsSelected\n        allAcSeatsSelected\n        acoSeatManageURL\n      }\n      otherReferences {\n        __typename\n        type\n        bookingReference\n        airlineCode\n        airlineName\n      }\n      ticketInfo {\n        __typename\n        statusCode\n        ticketed\n        type\n        requestedDate\n        ticketPermissions {\n          __typename\n          code\n          name\n          query\n        }\n      }\n    }\n    ac2uErrorsWarnings {\n      __typename\n      code\n      type\n      message\n    }\n    error {\n      __typename\n      lang\n      context\n      systemService\n      systemErrorType\n      systemErrorCode\n      systemErrorMessage\n      friendlyCode\n      friendlyTitle\n      friendlyMessage\n      action {\n        __typename\n        number\n        buttonLabel\n        action\n      }\n    }\n    priceChange {\n      __typename\n      previousAmount\n      updatedAmount\n      changeTitle\n      changeMessage\n      actions {\n        __typename\n        number\n        buttonLabel\n        action\n      }\n    }\n    bounds {\n      __typename\n      number\n      origin\n      destination\n      departureDateTimeLocal\n      departureDateTimeGMT\n      arrivalDateTimeLocal\n      arrivalDateTimeGMT\n      durationTotal\n      type\n      flown\n      segmentCount\n      connectionCount\n      containsDirect\n      operatingDisclosure\n      market\n      segments {\n        __typename\n        number\n        origin\n        destination\n        departureDateTimeLocal\n        departureDateTimeGMT\n        arrivalDateTimeLocal\n        arrivalDateTimeGMT\n        duration\n        flow\n        schedChangePending\n        status\n        departsEarly\n        stopCount\n        waitlistInfo {\n          __typename\n          waitlisted\n          status\n          bookingClass\n          cabinName\n        }\n        upgradeInfo {\n          __typename\n          upgraded\n          status\n          bookingClass\n          cabinName\n        }\n        marketingFlightInfo {\n          __typename\n          flightNumber\n          carrierCode\n          carrierName\n        }\n        operatingFlightInfo {\n          __typename\n          acOperated\n          flightNumber\n          carrierCode\n          carrierName\n          website\n        }\n        stops {\n          __typename\n          airportStop\n          disembarkationRequired\n        }\n        aircraft {\n          __typename\n          code\n          name\n          cabins {\n            __typename\n            cabinCode\n            cabinName\n            shortCabin\n            seatType\n            attributes {\n              __typename\n              number\n              name\n              description\n              icon\n              shortlist\n              values {\n                __typename\n                value\n              }\n              categories {\n                __typename\n                categoryName\n              }\n            }\n          }\n        }\n        seats {\n          __typename\n          allSeatsSelected\n          useAcSeatSelection\n          selectedSeats {\n            __typename\n            passengerNumber\n            firstName\n            middleName\n            lastName\n            seatNumber\n            preferredSeat\n            seatAttributes {\n              __typename\n              number\n              name\n              description\n              icon\n              shortlist\n              values {\n                __typename\n                value\n              }\n              categories {\n                __typename\n                categoryName\n              }\n            }\n          }\n        }\n        meals {\n          __typename\n          code\n          name\n        }\n      }\n      connections {\n        __typename\n        number\n        airport\n        startDateTimeLocal\n        startDateTimeGMT\n        endDateTimeLocal\n        endDateTimeGMT\n        duration\n        overnight\n        previousFlight {\n          __typename\n          marketingCarrierCode\n          marketingFlightNumber\n        }\n        nextFlight {\n          __typename\n          marketingCarrierCode\n          marketingFlightNumber\n        }\n      }\n      checkInInformation {\n        __typename\n        checkInOpen\n        scheduledCheckInStartLocal\n        scheduledCheckInStartGMT\n        scheduledCheckInEndLocal\n        scheduledCheckInEndGMT\n        checkInWithAirCanada\n        checkInAirline\n        acCheckInSystem\n        checkInURL\n        usCheckIn\n      }\n      fare {\n        __typename\n        fareFamily {\n          __typename\n          fareFamily\n          shortFareFamily\n          bookingClass {\n            __typename\n            marketingCode\n            flightNumber\n            bookingClassCode\n            fareBasisCode\n            selectionID\n          }\n          mixedCabin {\n            __typename\n            number\n            marketingCode\n            flightNumber\n            origin\n            destination\n            actualCabinCode\n            actualCabinName\n          }\n        }\n        cabin {\n          __typename\n          cabinCode\n          cabinName\n          shortCabinName\n        }\n        benefits {\n          __typename\n          ticketAttributes {\n            __typename\n            fareFamily\n            marketingCarrierCode\n            marketingFlightNumber\n            attributes {\n              __typename\n              number\n              name\n              description\n              icon\n              shortlist\n              assessment\n              values {\n                __typename\n                value\n              }\n              categories {\n                __typename\n                categoryName\n              }\n            }\n          }\n        }\n      }\n    }\n    passengers {\n      __typename\n      number\n      type\n      accompaniedByInfant\n      title\n      firstName\n      middleName\n      lastName\n      hasEmailAddress\n      emailAddress\n      hasPhoneNumber\n      phoneNumber\n      hasSecureFlight\n      hasTravelDocument\n      checkInStatus {\n        __typename\n        carrierCode\n        flightNumber\n        checkedIn\n        boardingPass\n      }\n      mealPreference {\n        __typename\n        code\n        name\n      }\n      loyalty {\n        __typename\n        fqtvProgramCode\n        fqtvProgramName\n        fqtvNumber\n        acTierName\n        acTierColour\n        saTierName\n        saTierColour\n      }\n      ticketNumber {\n        __typename\n        number\n      }\n      addresses {\n        __typename\n        addressLine\n        cityName\n        postalCode\n        stateProvince\n        country\n        remark\n      }\n    }\n    specialServiceRequest {\n      __typename\n      number\n      code\n      name\n      status\n      text\n      description\n      flights {\n        __typename\n        marketingCarrierCode\n        marketingFlightNumber\n        passengers {\n          __typename\n          firstName\n          lastName\n        }\n      }\n    }\n    selectedTravelerOption {\n      __typename\n      number\n      name\n      description {\n        __typename\n        lineItems {\n          __typename\n          itemNumber\n          itemText\n        }\n        images {\n          __typename\n          url\n        }\n      }\n      purchased {\n        __typename\n        type\n        origin\n        destination\n        included\n        price {\n          __typename\n          type\n          totalPrice\n        }\n        passenger {\n          __typename\n          firstName\n          lastName\n        }\n      }\n    }\n    paymentInfo {\n      __typename\n      transactionNumber\n      transactionType\n      amount\n      currency\n      method {\n        __typename\n        isCash\n        cardCode\n        cardType\n        cardNumber\n        cardExpiry\n      }\n    }\n    fareBreakdown {\n      __typename\n      fareSummary {\n        __typename\n        airChargesTotal\n        travelOptionsTotal\n        seatTotal\n        taxesFeesTotal\n        grandTotal\n      }\n      fareBreakdown {\n        __typename\n        airCharges {\n          __typename\n          baseFare {\n            __typename\n            origin\n            destination\n            fareFamily\n            cabinName\n            passengers {\n              __typename\n              type\n              quantity\n              amount\n            }\n          }\n          surcharges {\n            __typename\n            code\n            name\n            passengers {\n              __typename\n              type\n              quantity\n              amount\n            }\n          }\n          taxes {\n            __typename\n            code\n            name\n            details\n            passengers {\n              __typename\n              type\n              quantity\n              amount\n            }\n          }\n          fees {\n            __typename\n            code\n            name\n            details\n            passengers {\n              __typename\n              type\n              quantity\n              amount\n            }\n          }\n          beforeOptionsTotal {\n            __typename\n            type\n            quantity\n            amountTotal\n            allPassengersTotal\n          }\n        }\n        travelOptions {\n          __typename\n          purchased {\n            __typename\n            type\n            origin\n            destination\n            ancillaryOption {\n              __typename\n              code\n              name\n              totalPrice\n              passengers {\n                __typename\n                type\n                identifier\n                quantity\n                ancillaryPrice\n                amount\n              }\n            }\n          }\n          taxes {\n            __typename\n            code\n            name\n            details\n            amount\n          }\n          totalTravelOptions {\n            __typename\n            ancillaryTotal\n          }\n        }\n        seats {\n          __typename\n          selections {\n            __typename\n            passengerNumber\n            flightNumber\n            flightOrigin\n            flightDestination\n            seatNumber\n            seatCharacteristic\n            seatPosition\n            seatAmount\n          }\n          taxes {\n            __typename\n            code\n            name\n            details\n            amount\n          }\n        }\n      }\n    }\n  }\n}\n`
	reqBody     = `{"query": "{query}", "variables":{"language":"en","bookingReferenceID":"{conf}","surname":"{lname}","context":""}}`
	reqRegion  = "us-east-2"
	reqUserAgent = "aws-sdk-android/2.22.6 Linux/4.4.124+ Dalvik/2.1.0/0 en_US"
	reqEndpoint = "https://sfutz575ibfi5iloc3gwhasiqq.appsync-api.us-east-2.amazonaws.com/graphql"
)

var (
	client = http.Client{}
)

func getSigningCredentials() (*cognitoidentity.Credentials, error) {
	mySession := session.Must(session.NewSession())
	svc := cognitoidentity.New(mySession, aws.NewConfig().WithRegion(reqRegion))
	
	idRes, err := svc.GetId(&cognitoidentity.GetIdInput{
		IdentityPoolId: aws.String("us-east-2:6659d286-8971-4231-bf50-ba720e7ba3e3"),
	})

	if err != nil {
		return nil, errors.New("Couldn't get Identity ID")
	}

	creds, err := svc.GetCredentialsForIdentity(&cognitoidentity.GetCredentialsForIdentityInput{
		IdentityId: idRes.IdentityId,
	})

	if err != nil {
		return creds.Credentials, errors.New("Couldn't get identity Credentials")
	}

	return creds.Credentials, nil
}

func generateReqBody(lastName, confirmationCode string) string {
	body := reqBody
	body = strings.Replace(body, "{conf}", confirmationCode, -1)
	body = strings.Replace(body, "{lname}", lastName, -1)
	body = strings.Replace(body, "{query}", graphqlQuery, -1)
	return body
}

func buildRequest(endpoint, body string) (*http.Request, io.ReadSeeker) {
	reader := strings.NewReader(body)
	return buildRequestWithBodyReader(endpoint, reader)
}

func buildRequestWithBodyReader(endpoint string, body io.Reader) (*http.Request, io.ReadSeeker) {
	var bodyLen int

	type lenner interface {
		Len() int
	}
	if lr, ok := body.(lenner); ok {
		bodyLen = lr.Len()
	}

	req, _ := http.NewRequest("POST", endpoint, body)
	req.Header.Set("Content-Type", "application/x-amz-json-1.0")
	req.Header.Set("User-Agent", reqUserAgent)

	if bodyLen > 0 {
		req.Header.Set("Content-Length", strconv.Itoa(bodyLen))
	}

	var seeker io.ReadSeeker
	if sr, ok := body.(io.ReadSeeker); ok {
		seeker = sr
	} else {
		seeker = aws.ReadSeekCloser(body)
	}

	return req, seeker
}

func sendRequest(lastName, confirmationCode string) ([]byte, error) {
	req, body := buildRequest(reqEndpoint, generateReqBody(lastName, confirmationCode))

	creds, err := getSigningCredentials()
	if err != nil {
		return []byte{}, err
	}

	signer := v4.NewSigner(credentials.NewStaticCredentials(
		*creds.AccessKeyId,
		*creds.SecretKey,
		*creds.SessionToken,
	))

	signer.Sign(req, body, "appsync", reqRegion, time.Now())

	res, err := client.Do(req)

	if res.StatusCode != 200 {
		return []byte{}, errors.New("status code was not 200")
	}

	defer res.Body.Close()
	return ioutil.ReadAll(res.Body)
}


func performRequest(lastName , confirmationCode string) (res AirCanadaPNR, raw []byte,  err error) {

	data, err := sendRequest(lastName, confirmationCode)

	if err != nil {
		return res, data, err
	}

    var jsonresponse RetrievePnrResponse
	if err := json.Unmarshal(data, &jsonresponse); err != nil {
		return jsonresponse.Data.RetrievePNR, data, err
	}

	return jsonresponse.Data.RetrievePNR, data, nil
}