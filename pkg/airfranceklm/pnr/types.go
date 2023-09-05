package pnr

type PNR struct {
	Flags           []Flag
	Tickets         []Ticket
	ReissueRequired bool
	CheckedIn       bool
	SMCalcLink      string
	QMCalcLink      string
	HasDLFlights    bool
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
	MarketingAirlineName   string
	OperatingAirlineName   string
	ClassOfService         string
	UpgradeStatus          string
	ScheduledDeparture     string
	ScheduledArrival       string
	FlightNumber           string
	FareClass              string
	SkyPriority            bool
	AircraftType           string
}

type Passenger struct {
	Name                 string
	CheckedIn            bool
	Status               string
	FrequentFlyerNumber  string
	FrequentFlyerProgram string
}

type Flag struct {
	Name  string
	Value string
}

type Ticket struct {
	Number                 string
	ExpirationDate         string
	IssueDate              string
	Status                 string
	PassengerName          string
	NumCoupons             int
	ValidatedAgainstCoupon bool
	Flights                []Flight
	Passenger              Passenger
	Fare                   Fare
}

type Fare struct {
	BaseCurrencyCode            string
	BaseFare                    string
	TotalTax                    string
	TotalTaxCurrencyCode        string
	TotalFare                   string
	TotalCurrencyCode           string
	CarrierImposedSurcharge     string
	CarrierImposedSurchargeCode string
	TaxRows                     []TaxRow
	FareBasisCode               string
	EstimatedMQD                string
}

type TaxRow struct {
	TaxType           string
	Amount            string
	Currency          string
	CarrierImposedFee bool
}

type SeatProduct struct {
	SegmentID   int `json:"segmentId"`
	ProductName struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"productName"`
	ProductType struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"productType"`
	Seat struct {
		Row struct {
			Number int `json:"number"`
		} `json:"row"`
		Column struct {
			Code string `json:"code"`
		} `json:"column"`
		Characteristics []struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"characteristics"`
		ReservationStatus struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"reservationStatus"`
	} `json:"seat"`
	Status struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"status"`
	ProductClass string `json:"productClass"`
}

type BaggageProduct struct {
	ConnectionID int `json:"connectionId"`
	Amount       int `json:"amount"`
	Unit         struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"unit"`
	Type struct {
		Code  string `json:"code"`
		Name  string `json:"name"`
		Cabin bool   `json:"cabin"`
	} `json:"type"`
	Status struct {
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"status"`
}

type Segment struct {
	ID            int    `json:"id"`
	Type          string `json:"type"`
	Informational bool   `json:"informational"`
	DepartFrom    struct {
		IATACode string `json:"IATACode"`
		Name     string `json:"name"`
		City     struct {
			IATACode string `json:"IATACode"`
			Name     string `json:"name"`
			Country  struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"country"`
		} `json:"city"`
	} `json:"departFrom"`
	ArriveOn struct {
		IATACode string `json:"IATACode"`
		Name     string `json:"name"`
		City     struct {
			IATACode string `json:"IATACode"`
			Name     string `json:"name"`
			Country  struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"country"`
		} `json:"city"`
	} `json:"arriveOn"`
	MarketingFlight struct {
		Number  string `json:"number"`
		Carrier struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"carrier"`
		Status struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"status"`
		NumberOfStops int `json:"numberOfStops"`
		SellingClass  struct {
			Code string `json:"code"`
		} `json:"sellingClass"`
		OperatingFlight struct {
			Number  string `json:"number"`
			Carrier struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"carrier"`
			Duration                string `json:"duration"`
			Flown                   bool   `json:"flown"`
			CheckInStart            string `json:"checkInStart"`
			LocalCheckInStart       string `json:"localCheckInStart"`
			CheckInEnd              string `json:"checkInEnd"`
			LocalCheckInEnd         string `json:"localCheckInEnd"`
			ScheduledArrival        string `json:"scheduledArrival"`
			LocalScheduledArrival   string `json:"localScheduledArrival"`
			ScheduledDeparture      string `json:"scheduledDeparture"`
			LocalScheduledDeparture string `json:"localScheduledDeparture"`
			ArrivalTerminal         struct {
				Name string `json:"name"`
			} `json:"arrivalTerminal"`
			DepartureTerminal struct {
				Name string `json:"name"`
			} `json:"departureTerminal"`
			AircraftOwner string `json:"aircraftOwner"`
			Equipment     struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"equipment"`
			InFlightServices []struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"inFlightServices"`
			ScheduledDepartureUTC string `json:"scheduledDepartureUTC"`
			ScheduledArrivalUTC   string `json:"scheduledArrivalUTC"`
		} `json:"operatingFlight"`
	} `json:"marketingFlight"`
	ReplacedBySegments []any  `json:"replacedBySegments"`
	HaulType           string `json:"haulType"`
	Notifications      []any  `json:"notifications"`
	SequenceNumber     int    `json:"sequenceNumber"`
	Ordinal            int    `json:"ordinal"`
}

type FrequentTravelerMembership struct {
	Number         string `json:"number"`
	LoyaltyProgram struct {
		Carrier struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"carrier"`
		Code string `json:"code"`
		Name string `json:"name"`
	} `json:"loyaltyProgram"`
	Links struct {
		Update struct {
			Href string `json:"href"`
		} `json:"update"`
	} `json:"_links"`
}

type GetPNRResponse struct {
	BookingCode       string `json:"bookingCode"`
	ReservationTypes  []any  `json:"reservationTypes"`
	ContactDetails    []any  `json:"contactDetails"`
	EmergencyContacts []struct {
		Name        string `json:"name"`
		PhoneNumber string `json:"phoneNumber"`
		Country     struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"country"`
		RelationType struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"relationType"`
		Refusal     bool  `json:"refusal"`
		Memberships []any `json:"memberships"`
		Links       struct {
			Update struct {
				Href string `json:"href"`
			} `json:"update"`
			Delete struct {
				Href string `json:"href"`
			} `json:"delete"`
		} `json:"_links"`
	} `json:"emergencyContacts"`
	Itinerary struct {
		Type        string `json:"type"`
		Connections []struct {
			ID       int    `json:"id"`
			Duration string `json:"duration"`
			Origin   struct {
				IATACode string `json:"IATACode"`
				Name     string `json:"name"`
				City     struct {
					IATACode string `json:"IATACode"`
					Name     string `json:"name"`
					Country  struct {
						Code string `json:"code"`
						Name string `json:"name"`
					} `json:"country"`
				} `json:"city"`
			} `json:"origin"`
			Destination struct {
				IATACode string `json:"IATACode"`
				Name     string `json:"name"`
				City     struct {
					IATACode string `json:"IATACode"`
					Name     string `json:"name"`
					Country  struct {
						Code string `json:"code"`
						Name string `json:"name"`
					} `json:"country"`
				} `json:"city"`
			} `json:"destination"`
			OAndDNumber int       `json:"oAndDNumber"`
			Segments    []Segment `json:"segments"`
			Explanation struct {
				RelName     string `json:"relName"`
				Description string `json:"description"`
			} `json:"explanation"`
			Notifications []any `json:"notifications"`
			Links         struct {
			} `json:"links"`
			ConnectionDeparted bool `json:"connectionDeparted"`
		} `json:"connections"`
		Links struct {
			RebookingEligibility struct {
				Href string `json:"href"`
			} `json:"rebookingEligibility"`
		} `json:"_links"`
	} `json:"itinerary"`
	Passengers []struct {
		ID        int    `json:"id"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
		Title     struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"title"`
		Type struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"type"`
		AlteaPassengerType struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"alteaPassengerType"`
		IATAPassengerType struct {
			Code string `json:"code"`
			Name string `json:"name"`
		} `json:"IATAPassengerType"`
		CurrentlyLoggedOn bool `json:"currentlyLoggedOn"`
		EarnQuote         struct {
			TotalMiles       int `json:"totalMiles"`
			Xp               int `json:"xp"`
			Uxp              int `json:"uxp"`
			EarnQuoteDetails []struct {
				Miles           int    `json:"miles"`
				ProductCategory string `json:"productCategory"`
				Xp              int    `json:"xp"`
			} `json:"earnQuoteDetails"`
		} `json:"earnQuote"`
		SupplementaryDocumentRequirements struct {
			ESTARequired bool `json:"ESTARequired"`
			ETARequired  bool `json:"ETARequired"`
		} `json:"supplementaryDocumentRequirements"`
		SocialMediaIdentity []any `json:"socialMediaIdentity"`
		Memberships         []struct {
			FrequentTravelerMembership FrequentTravelerMembership `json:"frequentTravelerMembership,omitempty"`
		} `json:"memberships"`
		AdvancePassengerInformation struct {
			APIDocuments []struct {
				SupplementaryDocument struct {
					PrefillFromProfile        bool   `json:"prefillFromProfile"`
					Number                    string `json:"number"`
					SupplementaryDocumentType struct {
						Code string `json:"code"`
						Name string `json:"name"`
					} `json:"supplementaryDocumentType"`
					CountryOfIssue struct {
						Code string `json:"code"`
						Name string `json:"name"`
					} `json:"countryOfIssue"`
					Carrier struct {
						Code string `json:"code"`
						Name string `json:"name"`
					} `json:"carrier"`
					Status struct {
						Code string `json:"code"`
						Name string `json:"name"`
					} `json:"status"`
					Links struct {
						Update struct {
							Href string `json:"href"`
						} `json:"update"`
						Delete struct {
							Href string `json:"href"`
						} `json:"delete"`
					} `json:"_links"`
				} `json:"supplementaryDocument,omitempty"`
				TravelDocument struct {
					PrefillFromProfile bool   `json:"prefillFromProfile"`
					Number             string `json:"number"`
					CountryOfIssue     struct {
						Code string `json:"code"`
						Name string `json:"name"`
					} `json:"countryOfIssue"`
					ExpiryDate         string `json:"expiryDate"`
					TravelDocumentType struct {
						Code string `json:"code"`
						Name string `json:"name"`
					} `json:"travelDocumentType"`
					Surname    string `json:"surname"`
					GivenNames string `json:"givenNames"`
					Gender     struct {
						Code string `json:"code"`
						Name string `json:"name"`
					} `json:"gender"`
					DateOfBirth string `json:"dateOfBirth"`
					Nationality struct {
						Code string `json:"code"`
						Name string `json:"name"`
					} `json:"nationality"`
					Holder  bool `json:"holder"`
					Carrier struct {
						Code string `json:"code"`
						Name string `json:"name"`
					} `json:"carrier"`
					Status struct {
						Code string `json:"code"`
						Name string `json:"name"`
					} `json:"status"`
					Links struct {
						Update struct {
							Href string `json:"href"`
						} `json:"update"`
						Delete struct {
							Href string `json:"href"`
						} `json:"delete"`
					} `json:"_links"`
				} `json:"travelDocument,omitempty"`
			} `json:"apiDocuments"`
			Addresses []struct {
				Carrier struct {
					Code string `json:"code"`
					Name string `json:"name"`
				} `json:"carrier"`
				Status struct {
					Code string `json:"code"`
					Name string `json:"name"`
				} `json:"status"`
				Country struct {
					Code string `json:"code"`
					Name string `json:"name"`
				} `json:"country"`
				PostalAddressType struct {
					Code string `json:"code"`
					Name string `json:"name"`
				} `json:"postalAddressType"`
				PrefillFromProfile bool `json:"prefillFromProfile"`
				SegmentID          int  `json:"segmentId"`
				Links              struct {
					Update struct {
						Href string `json:"href"`
					} `json:"update"`
					Delete struct {
						Href string `json:"href"`
					} `json:"delete"`
				} `json:"_links"`
			} `json:"addresses"`
		} `json:"advancePassengerInformation"`
		PassengerFlightDetails []struct {
			CheckinStatus struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"checkinStatus"`
			SegmentID int `json:"segmentId"`
			Cabin     struct {
				Code string `json:"code"`
				Name string `json:"name"`
			} `json:"cabin"`
			SkyPriority          bool   `json:"skyPriority"`
			SimplifiedCabinClass string `json:"simplifiedCabinClass"`
			Notifications        []struct {
				Type string `json:"type"`
				Name string `json:"name"`
				Code string `json:"code"`
				Text string `json:"text"`
			} `json:"notifications,omitempty"`
			LoungeAccess struct {
				EligibilityStatus string `json:"eligibilityStatus"`
				GuestPassenger    bool   `json:"guestPassenger"`
				InvitationAllowed bool   `json:"invitationAllowed"`
				Lounge            struct {
					LoungeName  string `json:"loungeName"`
					DisplayName string `json:"displayName"`
					Airport     struct {
						IATACode string `json:"IATACode"`
						Name     string `json:"name"`
					} `json:"airport"`
				} `json:"lounge"`
				LoungeInvitations []any `json:"loungeInvitations"`
			} `json:"loungeAccess,omitempty"`
		} `json:"passengerFlightDetails"`
		ContactDetails []struct {
			EmailAddress struct {
				EmailAddress string `json:"emailAddress"`
				Links        struct {
					Update struct {
						Href string `json:"href"`
					} `json:"update"`
					Delete struct {
						Href string `json:"href"`
					} `json:"delete"`
				} `json:"_links"`
			} `json:"emailAddress,omitempty"`
			PhoneNumber struct {
				PhoneNumber string `json:"phoneNumber"`
				Type        struct {
					Code string `json:"code"`
					Name string `json:"name"`
				} `json:"type"`
				Links struct {
					Update struct {
						Href string `json:"href"`
					} `json:"update"`
					Delete struct {
						Href string `json:"href"`
					} `json:"delete"`
				} `json:"_links"`
			} `json:"phoneNumber,omitempty"`
		} `json:"contactDetails"`
		EmergencyContacts []any `json:"emergencyContacts"`
		Tickets           []struct {
			Number        string `json:"number"`
			TicketingDate string `json:"ticketingDate"`
			BaseFare      struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"baseFare"`
			TotalFare struct {
				Amount   float64 `json:"amount"`
				Currency string  `json:"currency"`
			} `json:"totalFare"`
			FormOfPayments []struct {
				Type struct {
					Code string `json:"code"`
					Name string `json:"name"`
				} `json:"type"`
				AmountPaid struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"amountPaid"`
			} `json:"formOfPayments"`
			TicketBooklets []struct {
				Number        string `json:"number"`
				TicketCoupons []struct {
					Number    int    `json:"number"`
					SegmentID int    `json:"segmentId"`
					FareBasis string `json:"fareBasis"`
				} `json:"ticketCoupons"`
				InvalidTicketCoupons []struct {
					Number    int    `json:"number"`
					SegmentID int    `json:"segmentId"`
					FareBasis string `json:"fareBasis"`
				} `json:"invalidTicketCoupons"`
			} `json:"ticketBooklets"`
			AirTransportationCharges struct {
				TicketPrice struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"ticketPrice"`
				Charges []struct {
					Code   string `json:"code"`
					Name   string `json:"name"`
					Amount struct {
						Amount   float64 `json:"amount"`
						Currency string  `json:"currency"`
					} `json:"amount"`
					CostType struct {
						Code string `json:"code"`
						Name string `json:"name"`
					} `json:"costType"`
				} `json:"charges"`
			} `json:"airTransportationCharges"`
			Taxes []struct {
				Code   string `json:"code"`
				Name   string `json:"name"`
				Amount struct {
					Amount   float64 `json:"amount"`
					Currency string  `json:"currency"`
				} `json:"amount"`
				CostType struct {
					Code string `json:"code"`
					Name string `json:"name"`
				} `json:"costType"`
			} `json:"taxes"`
			IssuingCity         string `json:"issuingCity"`
			IssuingCountry      string `json:"issuingCountry"`
			IssuingCompany      string `json:"issuingCompany"`
			IssuingCurrency     string `json:"issuingCurrency"`
			IssuingIATAOfficeID string `json:"issuingIATAOfficeId"`
			TicketingOfficeID   string `json:"ticketingOfficeId"`
		} `json:"tickets"`
		OrderedProducts []struct {
			SeatProduct    SeatProduct    `json:"seatProduct,omitempty"`
			BaggageProduct BaggageProduct `json:"baggageProduct,omitempty"`
		} `json:"orderedProducts"`
		PenaltyFees         []any `json:"penaltyFees"`
		ResidualVouchers    []any `json:"residualVouchers"`
		ContractInformation []any `json:"contractInformation"`
		Links               struct {
			AddEmergencyContactDetails struct {
				Href string `json:"href"`
			} `json:"addEmergencyContactDetails"`
			AddContactDetails struct {
				Href string `json:"href"`
			} `json:"addContactDetails"`
			AddAddressAdvancePassengerInformation struct {
				Href string `json:"href"`
			} `json:"addAddressAdvancePassengerInformation"`
			AddAPIDocumentAdvancePassengerInformation struct {
				Href string `json:"href"`
			} `json:"addApiDocumentAdvancePassengerInformation"`
			Invoice struct {
				Href string `json:"href"`
			} `json:"invoice"`
			PushSubscriptionPreference struct {
				Href string `json:"href"`
			} `json:"pushSubscriptionPreference"`
			SocialMediaNotificationPreference struct {
				Href string `json:"href"`
			} `json:"socialMediaNotificationPreference"`
			PassengerClearance struct {
				Href string `json:"href"`
			} `json:"passengerClearance"`
		} `json:"_links"`
	} `json:"passengers"`
	Notifications []struct {
		Name string `json:"name"`
		Code string `json:"code"`
		Text string `json:"text"`
	} `json:"notifications"`
	ThirdPartyOrderedProducts []any `json:"thirdPartyOrderedProducts"`
	TicketTimeLimits          []any `json:"ticketTimeLimits"`
	TotalPricePerCurrency     []struct {
		Amount   float64 `json:"amount"`
		Currency string  `json:"currency"`
	} `json:"totalPricePerCurrency"`
	Links struct {
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Refresh struct {
			Href string `json:"href"`
		} `json:"refresh"`
		Shop struct {
			Href string `json:"href"`
		} `json:"shop"`
		AddEmergencyContactDetails struct {
			Href string `json:"href"`
		} `json:"addEmergencyContactDetails"`
		AddContactDetails struct {
			Href string `json:"href"`
		} `json:"addContactDetails"`
		Invoice struct {
			Href string `json:"href"`
		} `json:"invoice"`
		PushSubscriptionPreference struct {
			Href string `json:"href"`
		} `json:"pushSubscriptionPreference"`
		AddMembership struct {
			Href string `json:"href"`
		} `json:"addMembership"`
		NextBestActions struct {
			Href string `json:"href"`
		} `json:"nextBestActions"`
		Enumerations struct {
			Href string `json:"href"`
		} `json:"enumerations"`
		AdvancePassengerInformationValidations struct {
			Href string `json:"href"`
		} `json:"advancePassengerInformationValidations"`
		TripSummary struct {
			Href string `json:"href"`
		} `json:"tripSummary"`
		RewardUpgradeEligibility struct {
			Href string `json:"href"`
		} `json:"rewardUpgradeEligibility"`
		RefundEligibility struct {
			Href string `json:"href"`
		} `json:"refundEligibility"`
		UpgradeInAdvanceEligibility struct {
			Href string `json:"href"`
		} `json:"upgradeInAdvanceEligibility"`
		Ticketconditions struct {
			Href string `json:"href"`
		} `json:"ticketconditions"`
	} `json:"_links"`
	ReservationResourceID string `json:"reservationResourceId"`
	PosCountry            string `json:"posCountry"`
	PosCity               string `json:"posCity"`
	DoHopReservation      bool   `json:"doHopReservation"`
	CorRequired           bool   `json:"corRequired"`
}

type TokenResponse struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
	TokenType   string `json:"token_type"`
}
