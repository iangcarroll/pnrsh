package pnr

type PNR struct {
	Flights    []Flight
	Passengers []Passenger
	Tickets    []Ticket
	SMCalcLink string
}

type Flight struct {
	OriginAirportCode      string
	DestinationAirportCode string
	Distance               string
	CurrentActionCode      string
	PreviousActionCode     string
	Status                 string
	MarketingAirlineCode   string
	OperatingAirlineCode   string
	ClassOfService         string
	Cabin                  string
	UpgradeStatus          string
	ScheduledDeparture     string
	ScheduledArrival       string
	FlightNumber           string

	// On AM, this is a per-segment field... dunno.
	FareBasis string
	FareName  string
}

type Passenger struct {
	Name                string
	Status              string
	IsCEO               bool
	IsInfiniteElite     bool
	IsLifetimeCompanion bool
	IsUnitedClubMember  bool
	IsPresidentialPlus  bool
	SharesPosition      string
}

type SSR struct {
	AirlineCode string
	Type        string
	Remark      string
}

type Ticket struct {
	Number                string
	CouponNumber          string
	Status                string
	PreviousStatus        string
	PassengerName         string
	RelatedDocumentNumber string
	OriginDestination     string
	TotalCost             string
}

type GetPNRResponse struct {
	DeviceID      string `json:"deviceId"`
	SessionID     string `json:"sessionId"`
	Flow          string `json:"flow"`
	RecordLocator string `json:"recordLocator"`
	LastName      string `json:"lastName"`
	Pnr           struct {
		SessionID          string `json:"sessionId"`
		RecordLocator      string `json:"recordLocator"`
		DateCreated        string `json:"dateCreated"`
		Description        string `json:"description"`
		IsActive           bool   `json:"isActive"`
		TicketType         string `json:"ticketType"`
		NumberOfPassengers string `json:"numberOfPassengers"`
		Trips              []struct {
			Index             int         `json:"index"`
			Origin            string      `json:"origin"`
			OriginName        string      `json:"originName"`
			Destination       string      `json:"destination"`
			DestinationName   string      `json:"destinationName"`
			Stops             string      `json:"stops"`
			JourneyTime       string      `json:"journeyTime"`
			JourneyMileage    string      `json:"journeyMileage"`
			DepartureTime     string      `json:"departureTime"`
			ArrivalTime       string      `json:"arrivalTime"`
			DepartureTimeGMT  string      `json:"departureTimeGMT"`
			ArrivalTimeGMT    string      `json:"arrivalTimeGMT"`
			ArrivalOffset     string      `json:"arrivalOffset"`
			ServiceMap        string      `json:"serviceMap"`
			DepartureTimeSort string      `json:"departureTimeSort"`
			CabinType         string      `json:"cabinType"`
			FlightNumber      interface{} `json:"flightNumber"`
			Zeros             []string    `json:"zeros"`
		} `json:"trips"`
		Passengers []struct {
			PassengerName struct {
				Title  string `json:"title"`
				First  string `json:"first"`
				Middle string `json:"middle"`
				Last   string `json:"last"`
				Suffix string `json:"suffix"`
			} `json:"passengerName"`
			KnownTravelerNumber    string `json:"knownTravelerNumber"`
			KtnDisplaySequence     string `json:"ktnDisplaySequence"`
			RedressNumber          string `json:"redressNumber"`
			RedressDisplaySequence string `json:"redressDisplaySequence"`
			SsrDisplaySequence     string `json:"ssrDisplaySequence"`
			SharesPosition         string `json:"sharesPosition"`
			MileagePlus            struct {
				AccountBalance                int         `json:"accountBalance"`
				ActiveStatusCode              string      `json:"activeStatusCode"`
				ActiveStatusDescription       string      `json:"activeStatusDescription"`
				AllianceEliteLevel            int         `json:"allianceEliteLevel"`
				ClosedStatusCode              string      `json:"closedStatusCode"`
				ClosedStatusDescription       string      `json:"closedStatusDescription"`
				CurrentEliteLevel             int         `json:"currentEliteLevel"`
				CurrentEliteLevelDescription  string      `json:"currentEliteLevelDescription"`
				CurrentYearMoneySpent         float64     `json:"currentYearMoneySpent"`
				EliteMileageBalance           int         `json:"eliteMileageBalance"`
				EliteSegmentBalance           int         `json:"eliteSegmentBalance"`
				EliteSegmentDecimalPlaceValue int         `json:"eliteSegmentDecimalPlaceValue"`
				EncryptedPin                  string      `json:"encryptedPin"`
				EnrollDate                    string      `json:"enrollDate"`
				EnrollSourceCode              string      `json:"enrollSourceCode"`
				EnrollSourceDescription       string      `json:"enrollSourceDescription"`
				FlexPqmBalance                int         `json:"flexPqmBalance"`
				FutureEliteLevel              int         `json:"futureEliteLevel"`
				FutureEliteDescription        string      `json:"futureEliteDescription"`
				InstantEliteExpirationDate    string      `json:"instantEliteExpirationDate"`
				IsCEO                         bool        `json:"isCEO"`
				IsClosedPermanently           bool        `json:"isClosedPermanently"`
				IsClosedTemporarily           bool        `json:"isClosedTemporarily"`
				IsCurrentTrialEliteMember     bool        `json:"isCurrentTrialEliteMember"`
				IsFlexPqm                     bool        `json:"isFlexPqm"`
				IsInfiniteElite               bool        `json:"isInfiniteElite"`
				IsLifetimeCompanion           bool        `json:"isLifetimeCompanion"`
				IsLockedOut                   bool        `json:"isLockedOut"`
				IsMergePending                bool        `json:"isMergePending"`
				IsUnitedClubMember            bool        `json:"isUnitedClubMember"`
				IsPresidentialPlus            bool        `json:"isPresidentialPlus"`
				Key                           string      `json:"key"`
				LastActivityDate              string      `json:"lastActivityDate"`
				LastExpiredMile               int         `json:"lastExpiredMile"`
				LastFlightDate                string      `json:"lastFlightDate"`
				LastStatementBalance          int         `json:"lastStatementBalance"`
				LastStatementDate             string      `json:"lastStatementDate"`
				LifetimeEliteLevel            int         `json:"lifetimeEliteLevel"`
				LifetimeEliteMileageBalance   int         `json:"lifetimeEliteMileageBalance"`
				MileageExpirationDate         string      `json:"mileageExpirationDate"`
				NextYearEliteLevel            int         `json:"nextYearEliteLevel"`
				NextYearEliteLevelDescription string      `json:"nextYearEliteLevelDescription"`
				MileagePlusID                 string      `json:"mileagePlusId"`
				MileagePlusPin                string      `json:"mileagePlusPin"`
				PriorUnitedAccountNumber      string      `json:"priorUnitedAccountNumber"`
				StarAllianceEliteLevel        int         `json:"starAllianceEliteLevel"`
				MpCustomerID                  int         `json:"mpCustomerId"`
				VendorCode                    string      `json:"vendorCode"`
				PremierLevelExpirationDate    string      `json:"premierLevelExpirationDate"`
				InstantElite                  interface{} `json:"instantElite"`
				TravelBankAccountNumber       interface{} `json:"travelBankAccountNumber"`
				TravelBankBalance             float64     `json:"travelBankBalance"`
				TravelBankCurrencyCode        interface{} `json:"travelBankCurrencyCode"`
				TravelBankExpirationDate      interface{} `json:"travelBankExpirationDate"`
			} `json:"mileagePlus"`
			OaRewardPrograms     interface{}   `json:"oaRewardPrograms"`
			IsMPMember           bool          `json:"isMPMember"`
			SharesGivenName      string        `json:"sharesGivenName"`
			SelectedSpecialNeeds []interface{} `json:"selectedSpecialNeeds"`
			BirthDate            string        `json:"birthDate"`
			TravelerTypeCode     string        `json:"travelerTypeCode"`
			PricingPaxType       string        `json:"pricingPaxType"`
			Contact              struct {
				Emails []struct {
					Key           string      `json:"key"`
					Channel       interface{} `json:"channel"`
					EmailAddress  string      `json:"emailAddress"`
					IsPrivate     bool        `json:"isPrivate"`
					IsDefault     bool        `json:"isDefault"`
					IsPrimary     bool        `json:"isPrimary"`
					IsDayOfTravel bool        `json:"isDayOfTravel"`
				} `json:"emails"`
				PhoneNumbers []struct {
					AreaNumber             string `json:"areaNumber"`
					Attention              string `json:"attention"`
					ChannelCode            string `json:"channelCode"`
					ChannelCodeDescription string `json:"channelCodeDescription"`
					ChannelTypeCode        string `json:"channelTypeCode"`
					ChannelTypeDescription string `json:"channelTypeDescription"`
					ChannelTypeSeqNumber   int    `json:"channelTypeSeqNumber"`
					CountryCode            string `json:"countryCode"`
					CountryNumber          string `json:"countryNumber"`
					CountryPhoneNumber     string `json:"countryPhoneNumber"`
					CustomerID             int    `json:"customerId"`
					Description            string `json:"description"`
					DiscontinuedDate       string `json:"discontinuedDate"`
					EffectiveDate          string `json:"effectiveDate"`
					ExtensionNumber        string `json:"extensionNumber"`
					IsPrimary              bool   `json:"isPrimary"`
					IsPrivate              bool   `json:"isPrivate"`
					IsProfileOwner         bool   `json:"isProfileOwner"`
					Key                    string `json:"key"`
					LanguageCode           string `json:"languageCode"`
					MileagePlusID          string `json:"mileagePlusId"`
					PagerPinNumber         string `json:"pagerPinNumber"`
					PhoneNumber            string `json:"phoneNumber"`
					SharesCountryCode      string `json:"sharesCountryCode"`
					WrongPhoneDate         string `json:"wrongPhoneDate"`
					DeviceTypeCode         string `json:"deviceTypeCode"`
					DeviceTypeDescription  string `json:"deviceTypeDescription"`
					IsDayOfTravel          bool   `json:"isDayOfTravel"`
					IsSMSEnabled           bool   `json:"isSMSEnabled"`
				} `json:"phoneNumbers"`
			} `json:"contact"`
			EarnedMiles struct {
				FirstName              string `json:"firstName"`
				LastName               string `json:"lastName"`
				AwardMileTotalFormated string `json:"awardMileTotalFormated"`
				PqmTotalFormated       string `json:"pqmTotalFormated"`
				PqsTotalFormated       string `json:"pqsTotalFormated"`
				PqdTotalFormated       string `json:"pqdTotalFormated"`
				PqfTotalFormated       string `json:"pqfTotalFormated"`
				PqpTotalFormated       string `json:"pqpTotalFormated"`
				AwardMileTotalShort    string `json:"awardMileTotalShort"`
				PqmTotalShort          string `json:"pqmTotalShort"`
				PqsTotalShort          string `json:"pqsTotalShort"`
				PqdTotalShort          string `json:"pqdTotalShort"`
				PqfTotalShort          string `json:"pqfTotalShort"`
				PqpTotalShort          string `json:"pqpTotalShort"`
				IsMember               bool   `json:"isMember"`
				MemberLevelText        string `json:"memberLevelText"`
				NonMemberMessage       string `json:"nonMemberMessage"`
				RowData                []struct {
					SegmentText            string `json:"segmentText"`
					TripText               string `json:"tripText"`
					AwardMileText          string `json:"awardMileText"`
					PqmText                string `json:"pqmText"`
					PqsText                string `json:"pqsText"`
					PqdText                string `json:"pqdText"`
					PqfText                string `json:"pqfText"`
					PqpText                string `json:"pqpText"`
					IsElligibleForEarnings bool   `json:"isElligibleForEarnings"`
					IneligibleEarningsText string `json:"ineligibleEarningsText"`
					SecondLineText         string `json:"secondLineText"`
				} `json:"rowData"`
				HasInelligibleSegment bool   `json:"hasInelligibleSegment"`
				EarnedMilesType       string `json:"earnedMilesType"`
			} `json:"earnedMiles"`
			LoyaltyProgramProfile struct {
				LoyaltyProgramCarrierCode           string `json:"LoyaltyProgramCarrierCode"`
				LoyaltyProgramID                    string `json:"LoyaltyProgramID"`
				LoyaltyProgramMemberID              string `json:"LoyaltyProgramMemberID"`
				LoyaltyProgramMemberTierLevel       string `json:"LoyaltyProgramMemberTierLevel"`
				StarEliteLevel                      string `json:"StarEliteLevel"`
				LoyaltyProgramMemberTierDescription string `json:"LoyaltyProgramMemberTierDescription"`
				Balances                            []struct {
					Characteristics []struct {
						Code        string `json:"Code"`
						Value       string `json:"Value"`
						Description string `json:"Description"`
						Status      struct {
							Description string `json:"Description"`
						} `json:"Status,omitempty"`
					} `json:"Characteristics"`
				} `json:"Balances"`
			} `json:"loyaltyProgramProfile"`
			EmployeeProfile interface{} `json:"employeeProfile"`
			PnrCustomerID   string      `json:"pnrCustomerID"`
		} `json:"passengers"`
		Segments []struct {
			OperationoperatingCarrier struct {
				Code         string `json:"code"`
				Name         string `json:"name"`
				FlightNumber string `json:"flightNumber"`
			} `json:"operationoperatingCarrier"`
			Aircraft struct {
				Code      string `json:"code"`
				ShortName string `json:"shortName"`
				LongName  string `json:"longName"`
				ModelCode string `json:"modelCode"`
			} `json:"aircraft"`
			Meal                      string `json:"meal"`
			FlightTime                string `json:"flightTime"`
			GroundTime                string `json:"groundTime"`
			TotalTravelTime           string `json:"totalTravelTime"`
			ActualMileage             string `json:"actualMileage"`
			MileagePlusMileage        string `json:"mileagePlusMileage"`
			Emp                       string `json:"emp"`
			TotalMileagePlusMileage   string `json:"totalMileagePlusMileage"`
			SsrMeals                  string `json:"ssrMeals"`
			ClassOfService            string `json:"classOfService"`
			ClassOfServiceDescription string `json:"classOfServiceDescription"`
			Seats                     []struct {
				Number                  string  `json:"number"`
				SeatRow                 string  `json:"seatRow"`
				SeatLetter              string  `json:"seatLetter"`
				PassengerSHARESPosition string  `json:"passengerSHARESPosition"`
				SeatStatus              string  `json:"seatStatus"`
				SegmentIndex            string  `json:"segmentIndex"`
				Origin                  string  `json:"origin"`
				Destination             string  `json:"destination"`
				EddNumber               string  `json:"eddNumber"`
				EDocID                  string  `json:"eDocId"`
				Price                   float64 `json:"price"`
				Currency                string  `json:"currency"`
				ProgramCode             string  `json:"programCode"`
			} `json:"seats"`
			NumberOfStops                 string        `json:"numberOfStops"`
			Stops                         []interface{} `json:"stops"`
			IsPlusPointSegment            bool          `json:"isPlusPointSegment"`
			HasPreviousSegmentDetails     bool          `json:"hasPreviousSegmentDetails"`
			ScheduledDepartureDateTimeGMT string        `json:"scheduledDepartureDateTimeGMT"`
			ScheduledArrivalDateTimeGMT   string        `json:"scheduledArrivalDateTimeGMT"`
			NoProtection                  string        `json:"noProtection"`
			CodeshareCarrier              struct {
				Code         string `json:"code"`
				Name         string `json:"name"`
				FlightNumber string `json:"flightNumber"`
			} `json:"codeshareCarrier"`
			CodeshareFlightNumber string      `json:"codeshareFlightNumber"`
			ScheduleChangeInfo    interface{} `json:"scheduleChangeInfo"`
			UpgradeVisibility     struct {
				CarrierCode               string `json:"carrierCode"`
				ClassOfService            string `json:"classOfService"`
				DepartureDateTime         string `json:"departureDateTime"`
				Destination               string `json:"destination"`
				FlightNumber              int    `json:"flightNumber"`
				Origin                    string `json:"origin"`
				PreviousSegmentActionCode string `json:"previousSegmentActionCode"`
				SegmentActionCode         string `json:"segmentActionCode"`
				SegmentNumber             int    `json:"segmentNumber"`
				UpgradeRemark             string `json:"upgradeRemark"`
				DecodedUpgradeMessage     string `json:"decodedUpgradeMessage"`
				UpgradeMessage            string `json:"upgradeMessage"`
				UpgradeMessageCode        string `json:"upgradeMessageCode"`
				UpgradeProperties         []struct {
					Key   string `json:"key"`
					Value string `json:"value"`
				} `json:"upgradeProperties"`
				UpgradeStatus      string      `json:"upgradeStatus"`
				UpgradeType        string      `json:"upgradeType"`
				WaitlistSegments   interface{} `json:"waitlistSegments"`
				RemarkAdvisoryType string      `json:"remarkAdvisoryType"`
			} `json:"upgradeVisibility"`
			LowestEliteLevel                   int         `json:"lowestEliteLevel"`
			Upgradeable                        bool        `json:"upgradeable"`
			UpgradeableMessageCode             string      `json:"upgradeableMessageCode"`
			UpgradeableMessage                 string      `json:"upgradeableMessage"`
			UpgradeableRemark                  string      `json:"upgradeableRemark"`
			ComplimentaryInstantUpgradeMessage string      `json:"complimentaryInstantUpgradeMessage"`
			Remove                             bool        `json:"remove"`
			Waitlisted                         bool        `json:"waitlisted"`
			ActionCode                         string      `json:"actionCode"`
			UpgradeEligible                    bool        `json:"upgradeEligible"`
			WaitedCOSDesc                      string      `json:"waitedCOSDesc"`
			LmxProducts                        interface{} `json:"lmxProducts"`
			CabinType                          string      `json:"cabinType"`
			NonPartnerFlight                   bool        `json:"nonPartnerFlight"`
			Bundles                            interface{} `json:"bundles"`
			IsElf                              bool        `json:"isElf"`
			IsIBE                              bool        `json:"isIBE"`
			TripNumber                         string      `json:"tripNumber"`
			SegmentNumber                      int         `json:"segmentNumber"`
			TicketCouponStatus                 string      `json:"ticketCouponStatus"`
			ProductCode                        string      `json:"productCode"`
			FareBasisCode                      string      `json:"fareBasisCode"`
			InCabinPetInfo                     interface{} `json:"inCabinPetInfo"`
			ShowSeatMapLink                    bool        `json:"showSeatMapLink"`
			IsChangeOfGuage                    bool        `json:"isChangeOfGuage"`
			MarketingCarrier                   struct {
				Code         string `json:"code"`
				Name         string `json:"name"`
				FlightNumber string `json:"flightNumber"`
			} `json:"marketingCarrier"`
			FlightNumber string `json:"flightNumber"`
			Departure    struct {
				Code string `json:"code"`
				Name string `json:"name"`
				City string `json:"city"`
			} `json:"departure"`
			Arrival struct {
				Code string `json:"code"`
				Name string `json:"name"`
				City string `json:"city"`
			} `json:"arrival"`
			ScheduledDepartureDateTime          string `json:"scheduledDepartureDateTime"`
			ScheduledArrivalDateTime            string `json:"scheduledArrivalDateTime"`
			ScheduledDepartureTimeGMT           string `json:"scheduledDepartureTimeGMT"`
			ScheduledArrivalTimeGMT             string `json:"scheduledArrivalTimeGMT"`
			FormattedScheduledDepartureDateTime string `json:"formattedScheduledDepartureDateTime"`
			FormattedScheduledArrivalDateTime   string `json:"formattedScheduledArrivalDateTime"`
			FormattedScheduledDepartureDate     string `json:"formattedScheduledDepartureDate"`
			FormattedScheduledArrivalDate       string `json:"formattedScheduledArrivalDate"`
			FormattedScheduledDepartureTime     string `json:"formattedScheduledDepartureTime"`
			FormattedScheduledArrivalTime       string `json:"formattedScheduledArrivalTime"`
		} `json:"segments"`
		LastTripDateDepartureDate       string        `json:"lastTripDateDepartureDate"`
		LastTripDateArrivalDate         string        `json:"lastTripDateArrivalDate"`
		CheckinEligible                 string        `json:"checkinEligible"`
		InfantInLaps                    []interface{} `json:"infantInLaps"`
		AlreadyCheckedin                string        `json:"alreadyCheckedin"`
		NotValid                        string        `json:"notValid"`
		ValidforCheckin                 string        `json:"validforCheckin"`
		PnrCanceled                     string        `json:"pnrCanceled"`
		UaRecordLocator                 string        `json:"uaRecordLocator"`
		CoRecordLocator                 string        `json:"coRecordLocator"`
		PnrOwner                        string        `json:"pnrOwner"`
		MealAccommodationAdvisory       string        `json:"mealAccommodationAdvisory"`
		MealAccommodationAdvisoryHeader string        `json:"mealAccommodationAdvisoryHeader"`
		OaRecordLocators                []interface{} `json:"oaRecordLocators"`
		OaRecordLocatorMessageTitle     string        `json:"oaRecordLocatorMessageTitle"`
		OaRecordLocatorMessage          string        `json:"oaRecordLocatorMessage"`
		IsEligibleToSeatChange          bool          `json:"isEligibleToSeatChange"`
		EmailAddress                    string        `json:"emailAddress"`
		SeatOffer                       struct {
			OfferTitle         string      `json:"offerTitle"`
			Price              float64     `json:"price"`
			CurrencyCode       string      `json:"currencyCode"`
			OfferText1         string      `json:"offerText1"`
			OfferText2         string      `json:"offerText2"`
			OfferText3         string      `json:"offerText3"`
			IsAdvanceSeatOffer bool        `json:"isAdvanceSeatOffer"`
			Miles              int         `json:"miles"`
			DisplayMiles       interface{} `json:"displayMiles"`
		} `json:"seatOffer"`
		IsShuttleOfferEligible    bool `json:"isShuttleOfferEligible"`
		IsETicketed               bool `json:"isETicketed"`
		ShouldDisplayEmailReceipt bool `json:"shouldDisplayEmailReceipt"`
		IsPetAvailable            bool `json:"isPetAvailable"`
		BundleInfo                struct {
			FlightSegments []interface{} `json:"flightSegments"`
			ErrorMessage   string        `json:"errorMessage"`
		} `json:"bundleInfo"`
		ShowOverride24HrFlex                      bool          `json:"showOverride24HrFlex"`
		ShowOverrideATREEligible                  bool          `json:"showOverrideATREEligible"`
		ScheduleChangeInfo                        interface{}   `json:"scheduleChangeInfo"`
		IrropsChangeInfo                          interface{}   `json:"irropsChangeInfo"`
		AssociateMPID                             string        `json:"associateMPId"`
		AssociateMPIDSharesPosition               string        `json:"associateMPIdSharesPosition"`
		AssociateMPIDSharesGivenName              string        `json:"associateMPIdSharesGivenName"`
		AssociateMPIDMessage                      string        `json:"associateMPIdMessage"`
		PetRecordLocators                         []interface{} `json:"petRecordLocators"`
		UpgradeMessage                            string        `json:"upgradeMessage"`
		FarelockExpirationDate                    string        `json:"farelockExpirationDate"`
		FarelockPurchaseMessage                   string        `json:"farelockPurchaseMessage"`
		EarnedMilesHeader                         string        `json:"earnedMilesHeader"`
		EarnedMilesText                           string        `json:"earnedMilesText"`
		IneligibleToEarnCreditMessage             string        `json:"ineligibleToEarnCreditMessage"`
		OaIneligibleToEarnCreditMessage           string        `json:"oaIneligibleToEarnCreditMessage"`
		AwardTravel                               bool          `json:"awardTravel"`
		PsSaTravel                                bool          `json:"psSaTravel"`
		SupressLMX                                bool          `json:"supressLMX"`
		IsELF                                     bool          `json:"isELF"`
		ElfLimitations                            interface{}   `json:"elfLimitations"`
		OverMileageLimitMessage                   string        `json:"overMileageLimitMessage"`
		OverMileageLimitAmount                    string        `json:"overMileageLimitAmount"`
		Lmxtravelers                              interface{}   `json:"lmxtravelers"`
		TripType                                  string        `json:"tripType"`
		IsTPIIncluded                             bool          `json:"isTPIIncluded"`
		IsFareLockOrNRSA                          bool          `json:"isFareLockOrNRSA"`
		HasCheckedBags                            bool          `json:"hasCheckedBags"`
		IsIBELite                                 bool          `json:"isIBELite"`
		IsIBE                                     bool          `json:"isIBE"`
		HasScheduleChanged                        bool          `json:"hasScheduleChanged"`
		StatusMessageItems                        interface{}   `json:"statusMessageItems"`
		ProductCategory                           string        `json:"productCategory"`
		ProductCode                               string        `json:"productCode"`
		MarketType                                string        `json:"marketType"`
		JourneyType                               string        `json:"journeyType"`
		CheckInStatus                             string        `json:"checkInStatus"`
		GetCheckInStatusFromCSLPNRRetrivalService bool          `json:"getCheckInStatusFromCSLPNRRetrivalService"`
		IrrOps                                    bool          `json:"irrOps"`
		IrrOpsViewed                              bool          `json:"irrOpsViewed"`
		IsEnableEditTraveler                      bool          `json:"isEnableEditTraveler"`
		IsUnaccompaniedMinor                      bool          `json:"isUnaccompaniedMinor"`
		IsCanceledWithFutureFlightCredit          bool          `json:"isCanceledWithFutureFlightCredit"`
		Umnr                                      interface{}   `json:"umnr"`
		InCabinPetInfo                            interface{}   `json:"inCabinPetInfo"`
		Futureflightcredit                        interface{}   `json:"futureflightcredit"`
		ShuttleOfferInformation                   interface{}   `json:"shuttleOfferInformation"`
		EarnedMilesInfo                           []struct {
			ID            string `json:"id"`
			CurrentValue  string `json:"currentValue"`
			SaveToPersist bool   `json:"saveToPersist"`
		} `json:"earnedMilesInfo"`
		IsCheckinEligible         bool        `json:"isCheckinEligible"`
		IsAgencyBooking           bool        `json:"isAgencyBooking"`
		AgencyName                interface{} `json:"agencyName"`
		IsPolicyExceptionAlert    bool        `json:"isPolicyExceptionAlert"`
		IsBEChangeEligible        bool        `json:"isBEChangeEligible"`
		ShouldDisplayUpgradeCabin bool        `json:"shouldDisplayUpgradeCabin"`
		IsCorporateBooking        bool        `json:"isCorporateBooking"`
		CorporateVendorName       string      `json:"corporateVendorName"`
		AdvisoryInfo              []struct {
			AdvisoryType  string      `json:"advisoryType"`
			ContentType   string      `json:"contentType"`
			Header        string      `json:"header"`
			SubTitle      interface{} `json:"subTitle"`
			Body          string      `json:"body"`
			Footer        interface{} `json:"footer"`
			Buttontext    string      `json:"buttontext"`
			Buttonlink    interface{} `json:"buttonlink"`
			IsBodyAsHTML  bool        `json:"isBodyAsHtml"`
			ShouldExpand  bool        `json:"shouldExpand"`
			IsDefaultOpen bool        `json:"isDefaultOpen"`
			ButtonItems   interface{} `json:"buttonItems"`
			SubItems      interface{} `json:"subItems"`
		} `json:"advisoryInfo"`
		ConsolidateScheduleChangeMessage bool        `json:"consolidateScheduleChangeMessage"`
		Prices                           interface{} `json:"prices"`
		IsTicketedByUA                   bool        `json:"isTicketedByUA"`
		IsSCChangeEligible               bool        `json:"isSCChangeEligible"`
		IsSCRefundEligible               bool        `json:"isSCRefundEligible"`
		IsSCBulkGroupPWC                 bool        `json:"isSCBulkGroupPWC"`
		IsMilesAndMoney                  bool        `json:"isMilesAndMoney"`
		OneClickEnrollmentEligibility    struct {
			JoinMileagePlusHeader string `json:"joinMileagePlusHeader"`
			JoinMileagePlusText   string `json:"joinMileagePlusText"`
			JoinMileagePlus       string `json:"joinMileagePlus"`
		} `json:"oneClickEnrollmentEligibility"`
		IsATREEligible               bool          `json:"isATREEligible"`
		IsInflightMealsOfferEligible bool          `json:"isInflightMealsOfferEligible"`
		SyncedWithConcur             string        `json:"syncedWithConcur"`
		URLItems                     []interface{} `json:"urlItems"`
		EditTravelerInfo             interface{}   `json:"editTravelerInfo"`
		EncryptPNR                   interface{}   `json:"encryptPNR"`
		EncryptLastName              interface{}   `json:"encryptLastName"`
		Isgroup                      bool          `json:"isgroup"`
		IsBulk                       bool          `json:"isBulk"`
		FareLockMessage              string        `json:"fareLockMessage"`
		FareLockPurchaseButton       string        `json:"fareLockPurchaseButton"`
		FareLockPriceButton          string        `json:"fareLockPriceButton"`
	} `json:"pnr"`
	UaRecordLocator       string      `json:"uaRecordLocator"`
	DotBagRules           interface{} `json:"dotBagRules"`
	DotBaggageInformation interface{} `json:"dotBaggageInformation"`
	CountDownWidgetInfo   struct {
		SectionTitle           string `json:"sectionTitle"`
		SectionDescription     string `json:"sectionDescription"`
		InstructionLinkText    string `json:"instructionLinkText"`
		InstructionPageTitle   string `json:"instructionPageTitle"`
		InstructionPageContent string `json:"instructionPageContent"`
	} `json:"countDownWidgetInfo"`
	PremierAccess  interface{} `json:"premierAccess"`
	RewardPrograms []struct {
		ProgramID   string `json:"programID"`
		Type        string `json:"type"`
		Description string `json:"description"`
	} `json:"rewardPrograms"`
	SpecialNeeds struct {
		SpecialMeals    interface{} `json:"specialMeals"`
		SpecialRequests []struct {
			Value                      string      `json:"value"`
			DisplayDescription         string      `json:"displayDescription"`
			Type                       string      `json:"type"`
			Code                       string      `json:"code"`
			DisplaySequence            interface{} `json:"displaySequence"`
			SubOptions                 interface{} `json:"subOptions"`
			Messages                   interface{} `json:"messages"`
			RegisterServiceDescription string      `json:"registerServiceDescription"`
			SubOptionHeader            interface{} `json:"subOptionHeader"`
		} `json:"specialRequests"`
		ServiceAnimals       interface{} `json:"serviceAnimals"`
		SpecialMealsMessages []struct {
			ID            string `json:"id"`
			CurrentValue  string `json:"currentValue"`
			SaveToPersist bool   `json:"saveToPersist"`
		} `json:"specialMealsMessages"`
		SpecialRequestsMessages   interface{}   `json:"specialRequestsMessages"`
		ServiceAnimalsMessages    interface{}   `json:"serviceAnimalsMessages"`
		HighTouchItems            []interface{} `json:"highTouchItems"`
		MealUnavailable           string        `json:"mealUnavailable"`
		AccommodationsUnavailable string        `json:"accommodationsUnavailable"`
	} `json:"specialNeeds"`
	ShowSeatChange    bool `json:"showSeatChange"`
	ShowPremierAccess bool `json:"showPremierAccess"`
	ShowAddCalendar   bool `json:"showAddCalendar"`
	ShowBaggageInfo   bool `json:"showBaggageInfo"`
	TripInsuranceInfo struct {
		ProductCode                      string  `json:"productCode"`
		ProductName                      string  `json:"productName"`
		DisplayAmount                    string  `json:"displayAmount"`
		FormattedDisplayAmount           string  `json:"formattedDisplayAmount"`
		Amount                           float64 `json:"amount"`
		CoverCost                        string  `json:"coverCost"`
		PageTitle                        string  `json:"pageTitle"`
		Title1                           string  `json:"title1"`
		Title2                           string  `json:"title2"`
		Title3                           string  `json:"title3"`
		QuoteTitle                       string  `json:"quoteTitle"`
		Headline1                        string  `json:"headline1"`
		Headline2                        string  `json:"headline2"`
		Body1                            string  `json:"body1"`
		Body2                            string  `json:"body2"`
		Body3                            string  `json:"body3"`
		LineItemText                     string  `json:"lineItemText"`
		TNC                              string  `json:"tNC"`
		Image                            string  `json:"image"`
		ProductID                        string  `json:"productId"`
		PaymentContent                   string  `json:"paymentContent"`
		PkDispenserPublicKey             string  `json:"pkDispenserPublicKey"`
		Confirmation1                    string  `json:"confirmation1"`
		Confirmation2                    string  `json:"confirmation2"`
		TpiAIGReturnedMessageContentList []struct {
			ID            string `json:"id"`
			CurrentValue  string `json:"currentValue"`
			SaveToPersist bool   `json:"saveToPersist"`
		} `json:"tpiAIGReturnedMessageContentList"`
	} `json:"tripInsuranceInfo"`
	Ancillary struct {
		AwardAccelerators struct {
			OfferTile struct {
				Title                   string      `json:"title"`
				Text1                   string      `json:"text1"`
				Text2                   string      `json:"text2"`
				Text3                   string      `json:"text3"`
				Price                   float64     `json:"price"`
				CurrencyCode            string      `json:"currencyCode"`
				Miles                   int         `json:"miles"`
				DisplayMiles            interface{} `json:"displayMiles"`
				ShowUpliftPerMonthPrice bool        `json:"showUpliftPerMonthPrice"`
			} `json:"offerTile"`
			EligibleTravelers      interface{} `json:"eligibleTravelers"`
			OptionsTnCs            interface{} `json:"optionsTnCs"`
			PaymentTnCs            interface{} `json:"paymentTnCs"`
			Captions               interface{} `json:"captions"`
			AwardAcceleratorTnCs   interface{} `json:"awardAcceleratorTnCs"`
			PremierAcceleratorTnCs interface{} `json:"premierAcceleratorTnCs"`
			TermsAndConditions     interface{} `json:"termsAndConditions"`
		} `json:"awardAccelerators"`
		PremiumCabinUpgrade struct {
			ProductCode string `json:"productCode"`
			ProductName string `json:"productName"`
			OfferTile   struct {
				Title                   string      `json:"title"`
				Text1                   string      `json:"text1"`
				Text2                   string      `json:"text2"`
				Text3                   string      `json:"text3"`
				Price                   float64     `json:"price"`
				CurrencyCode            string      `json:"currencyCode"`
				Miles                   int         `json:"miles"`
				DisplayMiles            interface{} `json:"displayMiles"`
				ShowUpliftPerMonthPrice bool        `json:"showUpliftPerMonthPrice"`
			} `json:"offerTile"`
			PcuOptions struct {
				EligibleTravelers []string `json:"eligibleTravelers"`
				EligibleSegments  []struct {
					SegmentNumber             int      `json:"segmentNumber"`
					FlightDescription         string   `json:"flightDescription"`
					FormattedPrice            string   `json:"formattedPrice"`
					UpgradeDescription        string   `json:"upgradeDescription"`
					Price                     float64  `json:"price"`
					SegmentDescription        string   `json:"segmentDescription"`
					CabinDescription          string   `json:"cabinDescription"`
					NoOfTravelersText         string   `json:"noOfTravelersText"`
					TotalPriceForAllTravelers float64  `json:"totalPriceForAllTravelers"`
					ProductIds                []string `json:"productIds"`
					IsUpgradeFailed           bool     `json:"isUpgradeFailed"`
					UpgradeOptions            []struct {
						OptionID                       string   `json:"optionId"`
						UpgradeOptionDescription       string   `json:"upgradeOptionDescription"`
						CabinDescriptionForPaymentPage string   `json:"cabinDescriptionForPaymentPage"`
						FormattedPrice                 string   `json:"formattedPrice"`
						Price                          float64  `json:"price"`
						TotalPriceForAllTravelers      float64  `json:"totalPriceForAllTravelers"`
						ProductIds                     []string `json:"productIds"`
					} `json:"upgradeOptions"`
					Origin       string `json:"origin"`
					Destination  string `json:"destination"`
					FlightNumber string `json:"flightNumber"`
				} `json:"eligibleSegments"`
				CompareOptions []struct {
					ImageURL string `json:"imageUrl"`
					Product  string `json:"product"`
					Header   string `json:"header"`
					Body     string `json:"body"`
				} `json:"compareOptions"`
				CurrencyCode string `json:"currencyCode"`
			} `json:"pcuOptions"`
			Captions []struct {
				ID            string `json:"id"`
				CurrentValue  string `json:"currentValue"`
				SaveToPersist bool   `json:"saveToPersist"`
			} `json:"captions"`
			MobileCmsContentMessages []struct {
				ContentFull  string `json:"contentFull"`
				ContentShort string `json:"contentShort"`
				HeadLine     string `json:"headLine"`
				LocationCode string `json:"locationCode"`
				Title        string `json:"title"`
			} `json:"mobileCmsContentMessages"`
			PkDispenserPublicKey string `json:"pkDispenserPublicKey"`
			CartID               string `json:"cartId"`
		} `json:"premiumCabinUpgrade"`
		PriorityBoarding struct {
			ProductCode     string `json:"productCode"`
			ProductName     string `json:"productName"`
			PbOfferTileInfo struct {
				Title                   string      `json:"title"`
				Text1                   string      `json:"text1"`
				Text2                   string      `json:"text2"`
				Text3                   string      `json:"text3"`
				Price                   float64     `json:"price"`
				CurrencyCode            string      `json:"currencyCode"`
				Miles                   int         `json:"miles"`
				DisplayMiles            interface{} `json:"displayMiles"`
				ShowUpliftPerMonthPrice bool        `json:"showUpliftPerMonthPrice"`
			} `json:"pbOfferTileInfo"`
			PbOfferDetails []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"pbOfferDetails"`
			TAndC []struct {
				Key   string `json:"key"`
				Value string `json:"value"`
			} `json:"tAndC"`
			Segments []struct {
				Origin        string      `json:"origin"`
				Destination   string      `json:"destination"`
				FlightDate    string      `json:"flightDate"`
				Customers     interface{} `json:"customers"`
				PbSegmentType string      `json:"pbSegmentType"`
				SegmentID     string      `json:"segmentId"`
				Message       string      `json:"message"`
				Fee           int         `json:"fee"`
				IsChecked     bool        `json:"isChecked"`
				CustPrice     int         `json:"custPrice"`
				OfferIds      []string    `json:"offerIds"`
			} `json:"segments"`
			PbDetailsConfirmationTxt string `json:"pbDetailsConfirmationTxt"`
			PbAddedTravelerTxt       string `json:"pbAddedTravelerTxt"`
		} `json:"priorityBoarding"`
		PlacePass           interface{} `json:"placePass"`
		TravelOptionsBundle interface{} `json:"travelOptionsBundle"`
		BasicEconomyBuyOut  interface{} `json:"basicEconomyBuyOut"`
	} `json:"ancillary"`
	ShowJoinOneClickEnrollment bool        `json:"showJoinOneClickEnrollment"`
	TransactionID              string      `json:"transactionId"`
	LanguageCode               string      `json:"languageCode"`
	MachineName                string      `json:"machineName"`
	CallDuration               int         `json:"callDuration"`
	Exception                  interface{} `json:"exception"`
}
