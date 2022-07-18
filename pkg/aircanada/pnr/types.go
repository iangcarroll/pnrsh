package pnr

type PNR struct {
	Data AirCanadaPNR

	RawResponse string
}

type RetrievePnrResponse struct {
	Data struct {
		RetrievePNR AirCanadaPNR `json:"retrievePNR"`
	} `json:"data"`
}

type AirCanadaAC2UErrorsWarnings struct {
	Code    string `json:code`
	Type    string `json:type`
	Message string `json:message`
}

type AirCanadaACBidUpgrade struct {
	IsEligible bool   `json:isEligible`
	IsModify   bool   `json:isModify`
	Url        string `json:url`
}

type AirCanadaAction struct {
	Number      string `json:number`
	ButtonLabel string `json:buttonLabel`
	Action      string `json:action`
}

type AirCanadaAddresses struct {
	AddressLine   []string `json:addressLine`
	CityName      string   `json:cityName`
	PostalCode    string   `json:postalCode`
	StateProvince string   `json:stateProvince`
	Country       string   `json:country`
	Remark        string   `json:remark`
}

type AirCanadaAirCharges struct {
	BaseFare           []AirCanadaBaseFare           `json:baseFare`
	Surcharges         []AirCanadaSurcharge          `json:surcharges`
	Taxes              []AirCanadaTax                `json:taxes`
	Fees               []AirCanadaFee                `json:fees`
	BeforeOptionsTotal []AirCanadaBeforeOptionsTotal `json:beforeOptionsTotal`
}

type AirCanadaAirChargesPassenger struct {
	Type     string `json:type`
	Quantity string `json:quantity`
	Amount   string `json:amount`
}

type AirCanadaAirCraft struct {
	Code   string         `json:code`
	Name   string         `json:name`
	Cabins AirCanadaCabin `json:cabins`
}

type AirCanadaAncillaryOption struct {
	Code       string                        `json:code`
	Name       string                        `json:name`
	TotalPrice string                        `json:totalPrice`
	Passengers []AirCanadaAncillaryPassenger `json:passengers`
}

type AirCanadaAncillaryPassenger struct {
	Type           string `json:type`
	Identifier     string `json:identifier`
	Quantity       string `json:quantity`
	AncillaryPrice string `json:ancillaryPrice`
	Amount         string `json:amount`
}

type AirCanadaBaseFare struct {
	Origin      string                         `json:origin`
	Destination string                         `json:destination`
	FareFamily  string                         `json:fareFamily`
	CabinName   string                         `json:cabinName`
	Passengers  []AirCanadaAirChargesPassenger `json:passengers`
}

type AirCanadaBeforeOptionsTotal struct {
	Type               string `json:type`
	Quantity           string `json:quantity`
	AmountTotal        string `json:amountTotal`
	AllPassengersTotal string `json:allPassengersTotal`
}

type AirCanadaBenefits struct {
	TicketAttributes []AirCanadaTicketAttribute `json:ticketAttributes`
}

type AirCanadaBookingClass struct {
	MarketingCode    string `json:marketingCode`
	FlightNumber     string `json:flightNumber`
	BookingClassCode string `json:bookingClassCode`
	FareBasisCode    string `json:fareBasisCode`
	SelectionID      string `json:selectionID`
}

type AirCanadaBookingInfo struct {
	Language         string                    `json:language`
	Currency         string                    `json:currency`
	PointOfSale      string                    `json:pointOfSale`
	BookingReference string                    `json:bookingReference`
	Source           string                    `json:source`
	TripType         string                    `json:tripType`
	EmployeeBooking  bool                      `json:employeeBooking`
	LastModified     string                    `json:lastModified`
	IsDisrupted      bool                      `json:isDisrupted`
	AcoManageUrl     string                    `json:acoManageUrl`
	AcBidUpgrade     AirCanadaACBidUpgrade     `json:acBidUpgrade`
	EUpgrade         AirCanadaEUpgrade         `json:eUpgrade`
	MixedFare        bool                      `json:mixedFare`
	SeatsInfo        AirCanadaBookingInfoSeats `json:seatsInfo`
	OtherReferences  []AirCanadaOtherReference `json:otherReferences`
	TicketInfo       AirCanadaTicketInfo       `json:ticketInfo`
}

type AirCanadaBookingInfoSeats struct {
	AllSeatsSelected   bool   `json:allSeatsSelected`
	AllAcSeatsSelected bool   `json:allAcSeatsSelected`
	AcoSeatManageURL   string `json:acoSeatManageURL`
}

type AirCanadaBound struct {
	Number                 string                      `json:number`
	Origin                 string                      `json:origin`
	Destination            string                      `json:destination`
	DepartureDateTimeLocal string                      `json:departureDateTimeLocal`
	DepartureDateTimeGMT   string                      `json:departureDateTimeGMT`
	ArrivalDateTimeLocal   string                      `json:arrivalDateTimeLocal`
	ArrivalDateTimeGMT     string                      `json:arrivalDateTimeGMT`
	DurationTotal          string                      `json:durationTotal`
	Type                   string                      `json:type`
	Flown                  bool                        `json:flown`
	SegmentCount           string                      `json:segmentCount`
	ConnectionCount        string                      `json:connectionCount`
	ContainsDirect         bool                        `json:containsDirect`
	OperatingDisclosure    string                      `json:operatingDisclosure`
	Market                 string                      `json:market`
	Segments               []AirCanadaBoundSegment     `json:segments`
	Connections            []AirCanadaConnection       `json:connections`
	CheckInInformation     AirCanadaCheckInInformation `json:checkInInformation`
	Fare                   AirCanadaFare               `json:fare`
}

type AirCanadaBoundSegment struct {
	Number                 string                       `json:number`
	Origin                 string                       `json:origin`
	Destination            string                       `json:destination`
	DepartureDateTimeLocal string                       `json:departureDateTimeLocal`
	DepartureDateTimeGMT   string                       `json:departureDateTimeGMT`
	ArrivalDateTimeLocal   string                       `json:arrivalDateTimeLocal`
	ArrivalDateTimeGMT     string                       `json:arrivalDateTimeGMT`
	Duration               string                       `json:duration`
	Flow                   bool                         `json:flow`
	SchedChangePending     bool                         `json:schedChangePending`
	Status                 string                       `json:status`
	DepartsEarly           bool                         `json:departsEarly`
	StopCount              string                       `json:stopCount`
	WaitlistInfo           AirCanadaWaitlistInfo        `json:waitlistInfo`
	UpgradeInfo            AirCanadaUpgradeInfo         `json:upgradeInfo`
	MarketingFlightInfo    AirCanadaMarketingFlightInfo `json:marketingFlightInfo`
	OperatingFlightInfo    AirCanadaOperatingFlightInfo `json:operatingFlightInfo`
	Stops                  []AirCanadaStop              `json:stops`
	Aircraft               AirCanadaAirCraft            `json:aircraft`
	Seats                  AirCanadaSegmentSeats        `json:seats`
	Meals                  []AirCanadaMeal              `json:meals`
}

type AirCanadaCabin struct {
	CabinCode  string                    `json:cabinCode`
	CabinName  string                    `json:cabinName`
	ShortCabin string                    `json:shortCabin`
	SeatType   string                    `json:seatType`
	Attributes []AirCanadaCabinAttribute `json:attributes`
}

type AirCanadaCabinAttribute struct {
	Number      string              `json:number`
	Name        string              `json:name`
	Description string              `json:description`
	Icon        string              `json:icon`
	Shortlist   bool                `json:shortlist`
	Values      []AirCanadaValue    `json:values`
	Categories  []AirCanadaCategory `json:categories`
}

type AirCanadaCategory struct {
	CategoryName string `json:categoryName`
}

type AirCanadaCheckInInformation struct {
	CheckInOpen                bool   `json:checkInOpen`
	ScheduledCheckInStartLocal string `json:scheduledCheckInStartLocal`
	ScheduledCheckInStartGMT   string `json:scheduledCheckInStartGMT`
	ScheduledCheckInEndLocal   string `json:scheduledCheckInEndLocal`
	ScheduledCheckInEndGMT     string `json:scheduledCheckInEndGMT`
	CheckInWithAirCanada       bool   `json:checkInWithAirCanada`
	CheckInAirline             string `json:checkInAirline`
	AcCheckInSystem            string `json:acCheckInSystem`
	CheckInURL                 string `json:checkInURL`
	UsCheckIn                  bool   `json:usCheckIn`
}

type AirCanadaCheckInStatus struct {
	CarrierCode  string `json:carrierCode`
	FlightNumber string `json:flightNumber`
	CheckedIn    bool   `json:checkedIn`
	BoardingPass bool   `json:boardingPass`
}

type AirCanadaConnection struct {
	Number             string          `json:number`
	Airport            string          `json:airport`
	StartDateTimeLocal string          `json:startDateTimeLocal`
	StartDateTimeGMT   string          `json:startDateTimeGMT`
	EndDateTimeLocal   string          `json:endDateTimeLocal`
	EndDateTimeGMT     string          `json:endDateTimeGMT`
	Duration           string          `json:duration`
	Overnight          bool            `json:overnight`
	PreviousFlight     AirCanadaFlight `json:previousFlight`
	NextFlight         AirCanadaFlight `json:nextFlight`
}

type AirCanadaDescription struct {
	LineItems []AirCanadaLineItem `json:lineItems`
	Images    []AirCanadaImage    `json:images`
}

type AirCanadaEUpgrade struct {
	IsEligible bool `json:isEligible`
}

type AirCanadaError struct {
	Lang               string                 `json:lang`
	Context            string                 `json:context`
	SystemService      string                 `json:systemService`
	SystemErrorType    string                 `json:systemErrorType`
	SystemErrorCode    string                 `json:systemErrorCode`
	SystemErrorMessage string                 `json:systemErrorMessage`
	FriendlyCode       string                 `json:friendlyCode`
	FriendlyTitle      string                 `json:friendlyTitle`
	FriendlyMessage    string                 `json:friendlyMessage`
	Action             []AirCanadaErrorAction `json:action`
}

type AirCanadaErrorAction struct {
	Number      string `json:number`
	ButtonLabel string `json:buttonLabel`
	Action      string `json:action`
}

type AirCanadaFare struct {
	FareFamily []AirCanadaFareFamily `json:fareFamily`
	Cabin      AirCanadaFareCabin    `json:cabin`
	Benefits   AirCanadaBenefits     `json:benefits`
}

type AirCanadaFareBreakdown struct {
	FareSummary   AirCanadaFareSummary        `json:fareSummary`
	FareBreakdown AirCanadaFareBreakdownChild `json:fareBreakdown`
}

type AirCanadaFareBreakdownChild struct {
	AirCharges    AirCanadaAirCharges    `json:airCharges`
	TravelOptions AirCanadaTravelOptions `json:travelOptions`
	Seats         AirCanadaFareSeats     `json:seats`
}

type AirCanadaFareCabin struct {
	CabinCode      string `json:cabinCode`
	CabinName      string `json:cabinName`
	ShortCabinName string `json:shortCabinName`
}

type AirCanadaFareFamily struct {
	FareFamily      string                  `json:fareFamily`
	ShortFareFamily string                  `json:shortFareFamily`
	BookingClass    []AirCanadaBookingClass `json:bookingClass`
	MixedCabin      []AirCanadaMixedCabin   `json:mixedCabin`
}

type AirCanadaFareSeatTax struct {
	Code    string `json:code`
	Name    string `json:name`
	Details string `json:details`
	Amount  string `json:amount`
}

type AirCanadaFareSeats struct {
	Selections []AirCanadaSeatSelection `json:selections`
	Taxes      []AirCanadaFareSeatTax   `json:taxes`
}

type AirCanadaFareSummary struct {
	AirChargesTotal    string `json:airChargesTotal`
	TravelOptionsTotal string `json:travelOptionsTotal`
	SeatTotal          string `json:seatTotal`
	TaxesFeesTotal     string `json:taxesFeesTotal`
	GrandTotal         string `json:grandTotal`
}

type AirCanadaFee struct {
	Code       string                         `json:code`
	Name       string                         `json:name`
	Details    string                         `json:details`
	Passengers []AirCanadaAirChargesPassenger `json:passengers`
}

type AirCanadaFlight struct {
	MarketingCarrierCode  string `json:marketingCarrierCode`
	MarketingFlightNumber string `json:marketingFlightNumber`
}

type AirCanadaFlightPassenger struct {
	FirstName string `json:firstName`
	LastName  string `json:lastName`
}

type AirCanadaImage struct {
	Url string `json:url`
}

type AirCanadaLineItem struct {
	ItemNumber string `json:itemNumber`
	ItemText   string `json:itemText`
}

type AirCanadaLoyalty struct {
	FqtvProgramCode string `json:fqtvProgramCode`
	FqtvProgramName string `json:fqtvProgramName`
	FqtvNumber      string `json:fqtvNumber`
	AcTierName      string `json:acTierName`
	AcTierColour    string `json:acTierColour`
	SaTierName      string `json:saTierName`
	SaTierColour    string `json:saTierColour`
}

type AirCanadaMarketingFlightInfo struct {
	FlightNumber string `json:flightNumber`
	CarrierCode  string `json:carrierCode`
	CarrierName  string `json:carrierName`
}

type AirCanadaMeal struct {
	Code string `json:code`
	Name string `json:name`
}

type AirCanadaMealPreference struct {
	Code string `json:code`
	Name string `json:name`
}

type AirCanadaMethod struct {
	IsCash     bool   `json:isCash`
	CardCode   string `json:cardCode`
	CardType   string `json:cardType`
	CardNumber string `json:cardNumber`
	CardExpiry string `json:cardExpiry`
}

type AirCanadaMixedCabin struct {
	Number          string `json:number`
	MarketingCode   string `json:marketingCode`
	FlightNumber    string `json:flightNumber`
	Origin          string `json:origin`
	Destination     string `json:destination`
	ActualCabinCode string `json:actualCabinCode`
	ActualCabinName string `json:actualCabinName`
}

type AirCanadaOperatingFlightInfo struct {
	AcOperated   bool   `json:acOperated`
	FlightNumber string `json:flightNumber`
	CarrierCode  string `json:carrierCode`
	CarrierName  string `json:carrierName`
	Website      string `json:website`
}

type AirCanadaOtherReference struct {
	Type             string `json:type`
	BookingReference string `json:bookingReference`
	AirlineCode      string `json:airlineCode`
	AirlineName      string `json:airlineName`
}

type AirCanadaPassenger struct {
	Number              string                   `json:number`
	Type                string                   `json:type`
	AccompaniedByInfant bool                     `json:accompaniedByInfant`
	Title               string                   `json:title`
	FirstName           string                   `json:firstName`
	MiddleName          string                   `json:middleName`
	LastName            string                   `json:lastName`
	HasEmailAddress     bool                     `json:hasEmailAddress`
	EmailAddress        string                   `json:emailAddress`
	HasPhoneNumber      bool                     `json:hasPhoneNumber`
	PhoneNumber         string                   `json:phoneNumber`
	HasSecureFlight     bool                     `json:hasSecureFlight`
	HasTravelDocument   bool                     `json:hasTravelDocument`
	CheckInStatus       []AirCanadaCheckInStatus `json:checkInStatus`
	MealPreference      AirCanadaMealPreference  `json:mealPreference`
	Loyalty             AirCanadaLoyalty         `json:loyalty`
	TicketNumber        []AirCanadaTicketNumber  `json:ticketNumber`
	Addresses           []AirCanadaAddresses     `json:addresses`
}

type AirCanadaPaymentInfo struct {
	TransactionNumber string          `json:transactionNumber`
	TransactionType   string          `json:transactionType`
	Amount            string          `json:amount`
	Currency          string          `json:currency`
	Method            AirCanadaMethod `json:method`
}

type AirCanadaPrice struct {
	Type       string `json:type`
	TotalPrice string `json:totalPrice`
}

type AirCanadaPriceChange struct {
	PreviousAmount string            `json:previousAmount`
	UpdatedAmount  string            `json:updatedAmount`
	ChangeTitle    string            `json:changeTitle`
	ChangeMessage  string            `json:changeMessage`
	Actions        []AirCanadaAction `json:actions`
}

type AirCanadaPurchased struct {
	Type        string                     `json:type`
	Origin      string                     `json:origin`
	Destination string                     `json:destination`
	Included    bool                       `json:included`
	Price       []AirCanadaPrice           `json:price`
	Passenger   []AirCanadaFlightPassenger `json:passenger`
}

type AirCanadaPurchasedOptions struct {
	Type            string                     `json:type`
	Origin          string                     `json:origin`
	Destination     string                     `json:destination`
	AncillaryOption []AirCanadaAncillaryOption `json:ancillaryOption`
}

type AirCanadaSSRFlight struct {
	MarketingCarrierCode  string                     `json:marketingCarrierCode`
	MarketingFlightNumber string                     `json:marketingFlightNumber`
	Passengers            []AirCanadaFlightPassenger `json:passengers`
}

type AirCanadaSeatAttributes struct {
	Number      string              `json:number`
	Name        string              `json:name`
	Description string              `json:description`
	Icon        string              `json:icon`
	Shortlist   bool                `json:shortlist`
	Values      []AirCanadaValue    `json:values`
	Categories  []AirCanadaCategory `json:categories`
}

type AirCanadaSeatSelection struct {
	PassengerNumber    string `json:passengerNumber`
	FlightNumber       string `json:flightNumber`
	FlightOrigin       string `json:flightOrigin`
	FlightDestination  string `json:flightDestination`
	SeatNumber         string `json:seatNumber`
	SeatCharacteristic string `json:seatCharacteristic`
	SeatPosition       string `json:seatPosition`
	SeatAmount         string `json:seatAmount`
}

type AirCanadaSegmentSeats struct {
	AllSeatsSelected   bool                            `json:allSeatsSelected`
	UseAcSeatSelection bool                            `json:useAcSeatSelection`
	SelectedSeats      []AirCanadaSegmentSelectedSeats `json:selectedSeats`
}

type AirCanadaSegmentSelectedSeats struct {
	PassengerNumber string                    `json:passengerNumber`
	FirstName       string                    `json:firstName`
	MiddleName      string                    `json:middleName`
	LastName        string                    `json:lastName`
	SeatNumber      string                    `json:seatNumber`
	PreferredSeat   bool                      `json:preferredSeat`
	SeatAttributes  []AirCanadaSeatAttributes `json:seatAttributes`
}

type AirCanadaSelectedTravelerOption struct {
	Number      string               `json:number`
	Name        string               `json:name`
	Description AirCanadaDescription `json:description`
	Purchased   []AirCanadaPurchased `json:purchased`
}

type AirCanadaSpecialServiceRequest struct {
	Number      string               `json:number`
	Code        string               `json:code`
	Name        string               `json:name`
	Status      string               `json:status`
	Text        string               `json:text`
	Description string               `json:description`
	Flights     []AirCanadaSSRFlight `json:flights`
}

type AirCanadaStop struct {
	AirportStop            string `json:airportStop`
	DisembarkationRequired bool   `json:disembarkationRequired`
}

type AirCanadaSurcharge struct {
	Code       string                         `json:code`
	Name       string                         `json:name`
	Passengers []AirCanadaAirChargesPassenger `json:passengers`
}

type AirCanadaTavelOptionsTax struct {
	Code    string `json:code`
	Name    string `json:name`
	Details string `json:details`
	Amount  string `json:amount`
}

type AirCanadaTax struct {
	Code       string                         `json:code`
	Name       string                         `json:name`
	Details    string                         `json:details`
	Passengers []AirCanadaAirChargesPassenger `json:passengers`
}

type AirCanadaTicketAttribute struct {
	FareFamily            string                   `json:fareFamily`
	MarketingCarrierCode  string                   `json:marketingCarrierCode`
	MarketingFlightNumber string                   `json:marketingFlightNumber`
	Attributes            []AirCanadaTicketBenefit `json:attributes`
}

type AirCanadaTicketBenefit struct {
	Number      string              `json:number`
	Name        string              `json:name`
	Description string              `json:description`
	Icon        string              `json:icon`
	Shortlist   bool                `json:shortlist`
	Assessment  string              `json:assessment`
	Values      []AirCanadaValue    `json:values`
	Categories  []AirCanadaCategory `json:categories`
}

type AirCanadaTicketInfo struct {
	StatusCode        string                       `json:statusCode`
	Ticketed          bool                         `json:ticketed`
	Type              string                       `json:type`
	RequestedDate     string                       `json:requestedDate`
	TicketPermissions []AirCanadaTicketPermissions `json:ticketPermissions`
}

type AirCanadaTicketNumber struct {
	Number string `json:number`
}

type AirCanadaTicketPermissions struct {
	Code  string `json:code`
	Name  string `json:name`
	Query string `json:query`
}

type AirCanadaTotalTravelOptions struct {
	AncillaryTotal string `json:ancillaryTotal`
}

type AirCanadaTravelOptions struct {
	Purchased          []AirCanadaPurchasedOptions `json:purchased`
	Taxes              []AirCanadaTavelOptionsTax  `json:taxes`
	TotalTravelOptions AirCanadaTotalTravelOptions `json:totalTravelOptions`
}

type AirCanadaUpgradeInfo struct {
	Upgraded     bool   `json:upgraded`
	Status       string `json:status`
	BookingClass string `json:bookingClass`
	CabinName    string `json:cabinName`
}

type AirCanadaValue struct {
	Value string `json:value`
}

type AirCanadaWaitlistInfo struct {
	Waitlisted   bool   `json:waitlisted`
	Status       string `json:status`
	BookingClass string `json:bookingClass`
	CabinName    string `json:cabinName`
}

type AirCanadaPNR struct {
	BookingInfo            AirCanadaBookingInfo              `json:bookingInfo`
	Ac2uErrorsWarnings     []AirCanadaAC2UErrorsWarnings     `json:ac2uErrorsWarnings`
	Error                  AirCanadaError                    `json:error`
	PriceChange            AirCanadaPriceChange              `json:priceChange`
	Bounds                 []AirCanadaBound                  `json:bounds`
	Passengers             []AirCanadaPassenger              `json:passengers`
	SpecialServiceRequest  []AirCanadaSpecialServiceRequest  `json:specialServiceRequest`
	SelectedTravelerOption []AirCanadaSelectedTravelerOption `json:selectedTravelerOption`
	PaymentInfo            []AirCanadaPaymentInfo            `json:paymentInfo`
	FareBreakdown          AirCanadaFareBreakdown            `json:fareBreakdown`
}
