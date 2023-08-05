package pnr

import "time"

type PNR struct {
	Flights    []Flight
	Passengers []Passenger
	Tickets    []Ticket
	Remarks    []Remark
	SSRs       []SSR
	SMCalcLink string
}

type Flight struct {
	OriginAirportCode      string
	DestinationAirportCode string
	Distance               string
	CurrentActionCode      string
	Status                 string
	MarketingAirlineCode   string
	OperatingAirlineCode   string
	ClassOfService         string
	Cabin                  string
	UpgradeStatus          string
	ScheduledDeparture     string
	ScheduledArrival       string
	FlightNumber           string
}

type Remark struct {
	Description     string
	DisplaySequence string
}

type Passenger struct {
	Name   string
	Status string
}

type SSR struct {
	Key         string
	Code        string
	Description string
	Comments    string
}

type Ticket struct {
	DocumentID         string
	IssueDate          string
	TicketValidityDate string
	Coupons            []Coupon
}

type Coupon struct {
	Status               string
	DepartureAirport     string
	ArrivalAirport       string
	FlightNumber         string
	OperatingAirlineCode string
}

type GetPNRResponse struct {
	CartID         string `json:"CartID"`
	ContentMessage struct {
		Body    interface{} `json:"Body"`
		Success string      `json:"Success"`
		Status  string      `json:"Status"`
		Results []struct {
			ComponentSchema string `json:"ComponentSchema"`
			ComponentTitle  string `json:"ComponentTitle"`
			ContentFull     string `json:"ContentFull"`
			LookupCode      string `json:"LookupCode"`
		} `json:"Results"`
	} `json:"ContentMessage"`
	Detail struct {
		EncryptedData []struct {
			FirstName          string `json:"FirstName"`
			LastName           string `json:"LastName"`
			EncryptFirstName   string `json:"EncryptFirstName"`
			EncryptLastName    string `json:"EncryptLastName"`
			ConfirmationNumber string `json:"ConfirmationNumber"`
			ItenNumber         string `json:"ItenNumber"`
		} `json:"EncryptedData"`
		Characteristic []struct {
			Code        string      `json:"Code"`
			Description interface{} `json:"Description"`
			Value       string      `json:"Value"`
			Status      interface{} `json:"Status"`
		} `json:"Characteristic"`
		ConfirmationID string `json:"ConfirmationID"`
		Channel        string `json:"Channel"`
		CreateDate     string `json:"CreateDate"`
		EmailAddress   []struct {
			Address string `json:"Address"`
		} `json:"EmailAddress"`
		FlightSegments []struct {
			BookingClass struct {
				Cabin struct {
					Name         string      `json:"Name"`
					Key          interface{} `json:"Key"`
					BoardingInfo interface{} `json:"BoardingInfo"`
					Description  string      `json:"Description"`
					Layout       interface{} `json:"Layout"`
					SeatRows     interface{} `json:"SeatRows"`
					TotalSeats   interface{} `json:"TotalSeats"`
				} `json:"Cabin"`
				Code string `json:"Code"`
			} `json:"BookingClass"`
			Characteristic []struct {
				Code        string      `json:"Code"`
				Description interface{} `json:"Description"`
				Value       string      `json:"Value"`
				Status      interface{} `json:"Status"`
			} `json:"Characteristic"`
			CurrentSeats []struct {
				ReservationNameIndex string      `json:"ReservationNameIndex"`
				EDocID               interface{} `json:"EDocID"`
				Seat                 struct {
					Identifier      string      `json:"Identifier"`
					SeatType        string      `json:"SeatType"`
					Description     string      `json:"Description"`
					Characteristics interface{} `json:"Characteristics"`
				} `json:"Seat"`
			} `json:"CurrentSeats"`
			FlightSegment struct {
				ActualArrivalTime   interface{} `json:"ActualArrivalTime"`
				IsGroupBooking      string      `json:"IsGroupBooking"`
				ActualDepartureTime interface{} `json:"ActualDepartureTime"`
				Amenities           interface{} `json:"Amenities"`
				ArrivalAirport      struct {
					Address struct {
						City          interface{} `json:"City"`
						StateProvince struct {
							Name string `json:"Name"`
						} `json:"StateProvince"`
					} `json:"Address"`
					IATACode        string `json:"IATACode"`
					ICAOCode        string `json:"ICAOCode"`
					IATACountryCode struct {
						CountryCode string `json:"CountryCode"`
					} `json:"IATACountryCode"`
					Name      string `json:"Name"`
					ShortName string `json:"ShortName"`
				} `json:"ArrivalAirport"`
				ArrivalDelayMinutes interface{} `json:"ArrivalDelayMinutes"`
				ArrivalBagClaimUnit interface{} `json:"ArrivalBagClaimUnit"`
				ArrivalDateTime     string      `json:"ArrivalDateTime"`
				ArrivalGate         interface{} `json:"ArrivalGate"`
				ArrivalTermimal     interface{} `json:"ArrivalTermimal"`
				ArrivalUTCDateTime  string      `json:"ArrivalUTCDateTime"`
				BoardTime           interface{} `json:"BoardTime"`
				BookingClasses      []struct {
					Cabin struct {
						Name         string      `json:"Name"`
						Key          interface{} `json:"Key"`
						BoardingInfo interface{} `json:"BoardingInfo"`
						Description  string      `json:"Description"`
						Layout       interface{} `json:"Layout"`
						SeatRows     interface{} `json:"SeatRows"`
						TotalSeats   interface{} `json:"TotalSeats"`
					} `json:"Cabin"`
					Code string `json:"Code"`
				} `json:"BookingClasses"`
				CabinCode      interface{} `json:"CabinCode"`
				ClassOfService interface{} `json:"ClassOfService"`
				Characteristic []struct {
					Code        string `json:"Code,omitempty"`
					Value       string `json:"Value"`
					Description string `json:"Description,omitempty"`
					Status      struct {
						Description string `json:"Description"`
						Code        string `json:"Code"`
					} `json:"Status,omitempty"`
				} `json:"Characteristic"`
				CurrentSeats     interface{} `json:"CurrentSeats"`
				DepartureAirport struct {
					Address struct {
						City          interface{} `json:"City"`
						StateProvince interface{} `json:"StateProvince"`
					} `json:"Address"`
					IATACode        string `json:"IATACode"`
					ICAOCode        string `json:"ICAOCode"`
					IATACountryCode struct {
						CountryCode string `json:"CountryCode"`
					} `json:"IATACountryCode"`
					Name      string `json:"Name"`
					ShortName string `json:"ShortName"`
				} `json:"DepartureAirport"`
				DepartureDateTime     string      `json:"DepartureDateTime"`
				DepartureDelayMinutes interface{} `json:"DepartureDelayMinutes"`
				DepartureGate         interface{} `json:"DepartureGate"`
				DepartureTerminal     interface{} `json:"DepartureTerminal"`
				DepartureUTCDateTime  string      `json:"DepartureUTCDateTime"`
				Equipment             struct {
					PseudoTailNumber  interface{} `json:"PseudoTailNumber"`
					CruiseSpeed       interface{} `json:"CruiseSpeed"`
					Engine            interface{} `json:"Engine"`
					SeatmapLegendList interface{} `json:"SeatmapLegendList"`
					Wingspan          interface{} `json:"Wingspan"`
					CabinCount        float64     `json:"CabinCount"`
					Cabins            interface{} `json:"Cabins"`
					Model             struct {
						Description string      `json:"Description"`
						Fleet       string      `json:"Fleet"`
						Genre       interface{} `json:"Genre"`
						Key         string      `json:"Key"`
						SubFleet    interface{} `json:"SubFleet"`
						VendorName  interface{} `json:"VendorName"`
					} `json:"Model"`
					NoseNumber       interface{} `json:"NoseNumber"`
					OwnerAirlineCode interface{} `json:"OwnerAirlineCode"`
					ShipNumber       interface{} `json:"ShipNumber"`
					TailNumber       interface{} `json:"TailNumber"`
				} `json:"Equipment"`
				EstimatedArrivalDelayMinutes   interface{} `json:"EstimatedArrivalDelayMinutes"`
				EstimatedArrivalTime           interface{} `json:"EstimatedArrivalTime"`
				EstimatedArrivalUTCTime        interface{} `json:"EstimatedArrivalUTCTime"`
				EstimatedDepartureDelayMinutes interface{} `json:"EstimatedDepartureDelayMinutes"`
				EstimatedDepartureTime         interface{} `json:"EstimatedDepartureTime"`
				EstimatedDepartureUTCTime      interface{} `json:"EstimatedDepartureUTCTime"`
				FlightNumber                   string      `json:"FlightNumber"`
				FlightSegmentType              string      `json:"FlightSegmentType"`
				FlightStatuses                 []struct {
					Code        string `json:"Code"`
					Description string `json:"Description"`
					StatusType  string `json:"StatusType"`
				} `json:"FlightStatuses"`
				PreviousSegmentDetails interface{} `json:"PreviousSegmentDetails"`
				ID                     interface{} `json:"ID"`
				IsActive               interface{} `json:"IsActive"`
				IsChangeOfGauge        string      `json:"IsChangeOfGauge"`
				IsCodeShare            string      `json:"IsCodeShare"`
				InboundFlightSegment   interface{} `json:"InboundFlightSegment"`
				InTime                 interface{} `json:"InTime"`
				InUTCTime              interface{} `json:"InUTCTime"`
				InstantUpgradable      bool        `json:"InstantUpgradable"`
				IsInternational        string      `json:"IsInternational"`
				InstanceUpgradable     interface{} `json:"InstanceUpgradable"`
				IsSeatMapPricing       interface{} `json:"IsSeatMapPricing"`
				JourneyDuration        string      `json:"JourneyDuration"`
				LOFNumber              int         `json:"LOFNumber"`
				MarketedFlightSegment  []struct {
					FlightNumber         string `json:"FlightNumber"`
					MarketingAirlineCode string `json:"MarketingAirlineCode"`
				} `json:"MarketedFlightSegment"`
				NoProtection            interface{} `json:"NoProtection"`
				Message                 interface{} `json:"Message"`
				OperatingAirline        interface{} `json:"OperatingAirline"`
				OperatingAirlineCode    string      `json:"OperatingAirlineCode"`
				OperatingAirlineName    string      `json:"OperatingAirlineName"`
				OffTime                 interface{} `json:"OffTime"`
				OffUTCTime              interface{} `json:"OffUTCTime"`
				OnTime                  interface{} `json:"OnTime"`
				OnUTCTime               interface{} `json:"OnUTCTime"`
				OutTime                 interface{} `json:"OutTime"`
				OutUTCTime              interface{} `json:"OutUTCTime"`
				PlannedTaxiIn           interface{} `json:"PlannedTaxiIn"`
				PlannedTaxiOut          interface{} `json:"PlannedTaxiOut"`
				ReasonStatuses          interface{} `json:"ReasonStatuses"`
				SegmentNumber           string      `json:"SegmentNumber"`
				Ship                    interface{} `json:"Ship"`
				SpclTerminalMsgsDetails interface{} `json:"SpclTerminalMsgsDetails"`
				StateProvince           interface{} `json:"StateProvince"`
				Status                  interface{} `json:"Status"`
				TypeOfSegment           string      `json:"TypeOfSegment"`
				UpgradeProperties       []struct {
					Code        string      `json:"Code"`
					Description interface{} `json:"Description"`
					Value       string      `json:"Value"`
				} `json:"UpgradeProperties"`
				UpgradeVisibilityType        string      `json:"UpgradeVisibilityType"`
				UpgradeEligibilityStatus     string      `json:"UpgradeEligibilityStatus"`
				FareBasisCode                string      `json:"FareBasisCode"`
				TicketDesignator             interface{} `json:"TicketDesignator"`
				NotValidAfter                interface{} `json:"NotValidAfter"`
				OperatingBookingClasses      interface{} `json:"OperatingBookingClasses"`
				OperatingAirlineFlightNumber string      `json:"OperatingAirlineFlightNumber"`
				Legs                         interface{} `json:"Legs"`
				FareComponentSequenceNumber  int         `json:"FareComponentSequenceNumber"`
				IsRulesEligible              string      `json:"IsRulesEligible"`
			} `json:"FlightSegment"`
			Legs         interface{} `json:"Legs"`
			IsConnection string      `json:"IsConnection"`
			Seat         struct {
				Identifier      string      `json:"Identifier"`
				SeatType        string      `json:"SeatType"`
				Description     string      `json:"Description"`
				Characteristics interface{} `json:"Characteristics"`
			} `json:"Seat"`
			SegmentNumber string `json:"SegmentNumber"`
			Status        struct {
				Key interface{} `json:"Key"`
			} `json:"Status"`
			TripNumber string `json:"TripNumber"`
			LOFNumber  string `json:"LOFNumber"`
		} `json:"FlightSegments"`
		Payment struct {
			AccountNumber string `json:"AccountNumber"`
			Amount        string `json:"Amount"`
			Currency      struct {
				Code string `json:"Code"`
			} `json:"Currency"`
			Type struct {
				Description string `json:"Description"`
				Key         string `json:"Key"`
				Value       string `json:"Value"`
			} `json:"Type"`
			WalletCategory string `json:"WalletCategory"`
		} `json:"Payment"`
		Pets             interface{} `json:"Pets"`
		MethodsofPayment []struct {
			CreditCard struct {
				Description                 int     `json:"Description"`
				CreditCardTypeCode          int     `json:"CreditCardTypeCode"`
				PinEntryCapability          int     `json:"PinEntryCapability"`
				InputCapability             int     `json:"InputCapability"`
				TerminalAttended            int     `json:"TerminalAttended"`
				TerminalCardholderActivated int     `json:"TerminalCardholderActivated"`
				Amount                      float64 `json:"Amount"`
				Date                        string  `json:"Date"`
				AccountNumber               string  `json:"AccountNumber"`
				Currency                    struct {
					Code string `json:"Code"`
				} `json:"Currency"`
				WalletCategory         int `json:"WalletCategory"`
				PaymentCharacteristics []struct {
					Code  string `json:"Code"`
					Value string `json:"Value"`
				} `json:"PaymentCharacteristics"`
				GroupPaymentType int `json:"GroupPaymentType"`
			} `json:"CreditCard"`
		} `json:"MethodsofPayment"`
		BookingIndicators struct {
		} `json:"BookingIndicators"`
		BookingAgency interface{} `json:"BookingAgency"`
		PointOfSale   struct {
			Country struct {
				CountryCode      string      `json:"CountryCode"`
				DefaultCurrency  interface{} `json:"DefaultCurrency"`
				DefaultLanguage  interface{} `json:"DefaultLanguage"`
				ISOAlpha3Code    interface{} `json:"ISOAlpha3Code"`
				Name             interface{} `json:"Name"`
				PhoneCountryCode interface{} `json:"PhoneCountryCode"`
				ShortName        interface{} `json:"ShortName"`
				Status           interface{} `json:"Status"`
			} `json:"Country"`
			CurrencyCode interface{} `json:"CurrencyCode"`
		} `json:"PointOfSale"`
		Prices []struct {
			BasePrice []struct {
				Code     string `json:"Code"`
				Currency struct {
					Code         string `json:"Code"`
					DecimalPlace int    `json:"DecimalPlace"`
				} `json:"Currency"`
				Amount float64 `json:"Amount"`
			} `json:"BasePrice"`
			Taxes []struct {
				Code     string `json:"Code"`
				Name     string `json:"Name"`
				Currency struct {
					Code         string `json:"Code"`
					DecimalPlace int    `json:"DecimalPlace"`
				} `json:"Currency"`
				Amount float64 `json:"Amount"`
				Status string  `json:"Status"`
			} `json:"Taxes"`
			Surcharges []interface{} `json:"Surcharges"`
			Fees       []interface{} `json:"Fees"`
			Totals     []struct {
				Name     string `json:"Name"`
				Currency struct {
					Code         string `json:"Code"`
					DecimalPlace int    `json:"DecimalPlace"`
				} `json:"Currency,omitempty"`
				Amount float64 `json:"Amount"`
			} `json:"Totals"`
			BasePriceEquivalent struct {
				Code     string `json:"Code"`
				Currency struct {
					Code         string `json:"Code"`
					DecimalPlace int    `json:"DecimalPlace"`
				} `json:"Currency"`
				Amount float64 `json:"Amount"`
			} `json:"BasePriceEquivalent"`
			Type struct {
				Value string `json:"Value"`
			} `json:"Type"`
			PinCode string `json:"PinCode"`
			Rules   []struct {
				Description string `json:"Description"`
			} `json:"Rules"`
			FareCalculation string `json:"FareCalculation"`
			FareComponents  []struct {
				AirlineCode         string   `json:"AirlineCode"`
				OriginCity          string   `json:"OriginCity"`
				DestinationCity     string   `json:"DestinationCity"`
				BasisCodes          []string `json:"BasisCodes"`
				BasePriceEquivalent struct {
					Code     string `json:"Code"`
					Currency struct {
						Code string `json:"Code"`
					} `json:"Currency"`
					Amount float64 `json:"Amount"`
				} `json:"BasePriceEquivalent"`
				BasePrice []struct {
					Code     string `json:"Code"`
					Currency struct {
						Code         string `json:"Code"`
						DecimalPlace int    `json:"DecimalPlace"`
					} `json:"Currency"`
					Amount float64 `json:"Amount"`
				} `json:"BasePrice"`
				DynamicFareCode        string        `json:"DynamicFareCode"`
				Taxes                  []interface{} `json:"Taxes"`
				DestinationAirportCode string        `json:"DestinationAirportCode"`
				OriginAirportCode      string        `json:"OriginAirportCode"`
				SequenceNumber         int           `json:"SequenceNumber"`
				SegmentReferenceIDs    []string      `json:"SegmentReferenceIDs"`
				Status                 string        `json:"Status"`
				DepartureDateTime      string        `json:"DepartureDateTime"`
				IsRulesEligible        string        `json:"IsRulesEligible"`
				TripNumber             string        `json:"TripNumber,omitempty"`
			} `json:"FareComponents"`
			Count struct {
				CountType string `json:"CountType"`
				Value     string `json:"Value"`
			} `json:"Count"`
			FareType            int `json:"FareType"`
			PriceFlightSegments []struct {
				FareBasisCode               string `json:"FareBasisCode"`
				LOFNumber                   int    `json:"LOFNumber"`
				FareComponentSequenceNumber int    `json:"FareComponentSequenceNumber"`
				NotValidAfter               string `json:"NotValidAfter"`
				NotValidBefore              string `json:"NotValidBefore"`
				GlobalIndicator             string `json:"GlobalIndicator"`
				DepartureAirport            struct {
					IATACode     string `json:"IATACode"`
					ShortName    string `json:"ShortName"`
					IATACityCode struct {
						CityCode string `json:"CityCode"`
					} `json:"IATACityCode"`
				} `json:"DepartureAirport"`
				ArrivalAirport struct {
					IATACode     string `json:"IATACode"`
					ShortName    string `json:"ShortName"`
					IATACityCode struct {
						CityCode string `json:"CityCode"`
					} `json:"IATACityCode"`
				} `json:"ArrivalAirport"`
				DepartureDateTime     string `json:"DepartureDateTime"`
				ArrivalDateTime       string `json:"ArrivalDateTime"`
				FlightNumber          string `json:"FlightNumber"`
				MarketedFlightSegment []struct {
					MarketingAirlineCode string `json:"MarketingAirlineCode"`
					FlightNumber         string `json:"FlightNumber"`
				} `json:"MarketedFlightSegment"`
				MarriageGroup  string `json:"MarriageGroup"`
				SegmentNumber  int    `json:"SegmentNumber"`
				BookingClasses []struct {
					Code string `json:"Code"`
				} `json:"BookingClasses"`
				Legs []struct {
					DepartureAirport struct {
						IATACode string `json:"IATACode"`
					} `json:"DepartureAirport"`
					ArrivalAirport struct {
						IATACode string `json:"IATACode"`
					} `json:"ArrivalAirport"`
					Equipment struct {
						Model struct {
							Fleet string `json:"Fleet"`
						} `json:"Model"`
						SeatmapLegendList interface{} `json:"SeatmapLegendList"`
					} `json:"Equipment"`
					CabinCount int `json:"CabinCount"`
				} `json:"Legs"`
				FlightStatuses []struct {
					StatusType  string `json:"StatusType"`
					Description string `json:"Description"`
					Code        string `json:"Code"`
				} `json:"FlightStatuses"`
				UpgradeEligibilityStatus int    `json:"UpgradeEligibilityStatus"`
				UpgradeVisibilityType    int    `json:"UpgradeVisibilityType"`
				InstantUpgradable        bool   `json:"InstantUpgradable"`
				ArrivalUTCDateTime       string `json:"ArrivalUTCDateTime,omitempty"`
				DepartureUTCDateTime     string `json:"DepartureUTCDateTime,omitempty"`
			} `json:"PriceFlightSegments"`
			PassengerTypeCode string `json:"PassengerTypeCode"`
			EquivalentTaxes   []struct {
				Code     string `json:"Code"`
				Currency struct {
					Code string `json:"Code"`
				} `json:"Currency"`
				Amount float64 `json:"Amount"`
			} `json:"EquivalentTaxes"`
			FareMode        string `json:"FareMode"`
			Characteristics []struct {
				Value       string `json:"Value"`
				Description string `json:"Description"`
			} `json:"Characteristics"`
			PassengerIDs struct {
				Key         string `json:"Key"`
				Description string `json:"Description"`
			} `json:"PassengerIDs"`
			POSBaseFare struct {
				Code     string `json:"Code"`
				Currency struct {
					Code         string `json:"Code"`
					DecimalPlace int    `json:"DecimalPlace"`
				} `json:"Currency"`
				Amount float64 `json:"Amount"`
			} `json:"POSBaseFare"`
		} `json:"Prices"`
		Services []struct {
			Key            string `json:"Key"`
			Description    string `json:"Description"`
			SequenceNumber int    `json:"SequenceNumber"`
			Comments       string `json:"Comments"`
			Type           struct {
				Key              string `json:"Key"`
				DefaultIndicator string `json:"DefaultIndicator"`
				DisplaySequence  int    `json:"DisplaySequence"`
			} `json:"Type,omitempty"`
			DisplaySequence int `json:"DisplaySequence"`
			Status          struct {
				Description string `json:"Description"`
				Key         string `json:"Key"`
				Code        string `json:"Code"`
			} `json:"Status,omitempty"`
			Code              string `json:"Code"`
			TravelerNameIndex string `json:"TravelerNameIndex"`
			SubCategory       string `json:"SubCategory,omitempty"`
			Category          string `json:"Category,omitempty"`
			CommentType       string `json:"CommentType,omitempty"`
			IsActive          bool   `json:"IsActive,omitempty"`
			SegmentNumber     []int  `json:"SegmentNumber,omitempty"`
			CarrierCode       struct {
				Name         string `json:"Name"`
				IsCassMember string `json:"IsCassMember"`
				IsMustRide   string `json:"IsMustRide"`
			} `json:"CarrierCode,omitempty"`
		} `json:"Services"`
		Remarks []struct {
			Description     string `json:"Description"`
			DisplaySequence string `json:"DisplaySequence"`
		} `json:"Remarks"`
		TelephoneNumbers []struct {
			AreaCityCode      string `json:"AreaCityCode"`
			CountryAccessCode string `json:"CountryAccessCode"`
			Description       string `json:"Description"`
			PhoneNumber       string `json:"PhoneNumber"`
		} `json:"TelephoneNumbers"`
		TimeLimitExpiration string `json:"TimeLimitExpiration"`
		Status              struct {
			DefaultIndicator string `json:"DefaultIndicator"`
			Description      string `json:"Description"`
			Key              string `json:"Key"`
		} `json:"Status"`
		Travelers []struct {
			CustomerID            interface{} `json:"CustomerID"`
			GivenName             interface{} `json:"GivenName"`
			Surname               interface{} `json:"Surname"`
			ID                    interface{} `json:"ID"`
			PassengerTypeCode     interface{} `json:"PassengerTypeCode"`
			EmployeeProfile       interface{} `json:"EmployeeProfile"`
			LoyaltyProgramProfile struct {
				LoyaltyProgramCarrierCode           string `json:"LoyaltyProgramCarrierCode"`
				LoyaltyProgramID                    string `json:"LoyaltyProgramID"`
				LoyaltyProgramMemberID              string `json:"LoyaltyProgramMemberID"`
				LoyaltyProgramMemberTierLevel       string `json:"LoyaltyProgramMemberTierLevel"`
				StarEliteLevel                      string `json:"StarEliteLevel"`
				LoyaltyProgramMemberTierDescription int    `json:"LoyaltyProgramMemberTierDescription"`
				Balances                            []struct {
					Characteristics []struct {
						Code        string `json:"Code"`
						Value       string `json:"Value"`
						Description string `json:"Description"`
						Status      struct {
							Description string `json:"Description"`
							Key         string `json:"Key"`
						} `json:"Status,omitempty"`
					} `json:"Characteristics"`
				} `json:"Balances"`
			} `json:"LoyaltyProgramProfile"`
			Tickets []struct {
				TicketingCarrierCode interface{} `json:"TicketingCarrierCode"`
				IssueDate            string      `json:"IssueDate"`
				DocumentID           string      `json:"DocumentID"`
				Type                 string      `json:"Type"`
				TicketValidityDate   string      `json:"TicketValidityDate"`
				IsConjunctiveTicket  string      `json:"IsConjunctiveTicket"`
				FlightCoupons        []struct {
					Status struct {
						Code string `json:"Code"`
					} `json:"Status"`
					FlightSegment struct {
						ActualArrivalTime   interface{} `json:"ActualArrivalTime"`
						IsGroupBooking      interface{} `json:"IsGroupBooking"`
						ActualDepartureTime interface{} `json:"ActualDepartureTime"`
						Amenities           interface{} `json:"Amenities"`
						ArrivalAirport      struct {
							Address         interface{} `json:"Address"`
							IATACode        string      `json:"IATACode"`
							ICAOCode        interface{} `json:"ICAOCode"`
							IATACountryCode interface{} `json:"IATACountryCode"`
							Name            interface{} `json:"Name"`
							ShortName       interface{} `json:"ShortName"`
						} `json:"ArrivalAirport"`
						ArrivalDelayMinutes interface{} `json:"ArrivalDelayMinutes"`
						ArrivalBagClaimUnit interface{} `json:"ArrivalBagClaimUnit"`
						ArrivalDateTime     interface{} `json:"ArrivalDateTime"`
						ArrivalGate         interface{} `json:"ArrivalGate"`
						ArrivalTermimal     interface{} `json:"ArrivalTermimal"`
						ArrivalUTCDateTime  interface{} `json:"ArrivalUTCDateTime"`
						BoardTime           interface{} `json:"BoardTime"`
						BookingClasses      interface{} `json:"BookingClasses"`
						CabinCode           interface{} `json:"CabinCode"`
						ClassOfService      interface{} `json:"ClassOfService"`
						Characteristic      interface{} `json:"Characteristic"`
						CurrentSeats        interface{} `json:"CurrentSeats"`
						DepartureAirport    struct {
							Address         interface{} `json:"Address"`
							IATACode        string      `json:"IATACode"`
							ICAOCode        interface{} `json:"ICAOCode"`
							IATACountryCode interface{} `json:"IATACountryCode"`
							Name            interface{} `json:"Name"`
							ShortName       interface{} `json:"ShortName"`
						} `json:"DepartureAirport"`
						DepartureDateTime              string      `json:"DepartureDateTime"`
						DepartureDelayMinutes          interface{} `json:"DepartureDelayMinutes"`
						DepartureGate                  interface{} `json:"DepartureGate"`
						DepartureTerminal              interface{} `json:"DepartureTerminal"`
						DepartureUTCDateTime           interface{} `json:"DepartureUTCDateTime"`
						Equipment                      interface{} `json:"Equipment"`
						EstimatedArrivalDelayMinutes   interface{} `json:"EstimatedArrivalDelayMinutes"`
						EstimatedArrivalTime           interface{} `json:"EstimatedArrivalTime"`
						EstimatedArrivalUTCTime        interface{} `json:"EstimatedArrivalUTCTime"`
						EstimatedDepartureDelayMinutes interface{} `json:"EstimatedDepartureDelayMinutes"`
						EstimatedDepartureTime         interface{} `json:"EstimatedDepartureTime"`
						EstimatedDepartureUTCTime      interface{} `json:"EstimatedDepartureUTCTime"`
						FlightNumber                   string      `json:"FlightNumber"`
						FlightSegmentType              interface{} `json:"FlightSegmentType"`
						FlightStatuses                 interface{} `json:"FlightStatuses"`
						PreviousSegmentDetails         interface{} `json:"PreviousSegmentDetails"`
						ID                             interface{} `json:"ID"`
						IsActive                       interface{} `json:"IsActive"`
						IsChangeOfGauge                interface{} `json:"IsChangeOfGauge"`
						IsCodeShare                    interface{} `json:"IsCodeShare"`
						InboundFlightSegment           interface{} `json:"InboundFlightSegment"`
						InTime                         interface{} `json:"InTime"`
						InUTCTime                      interface{} `json:"InUTCTime"`
						InstantUpgradable              bool        `json:"InstantUpgradable"`
						IsInternational                interface{} `json:"IsInternational"`
						InstanceUpgradable             interface{} `json:"InstanceUpgradable"`
						IsSeatMapPricing               interface{} `json:"IsSeatMapPricing"`
						JourneyDuration                interface{} `json:"JourneyDuration"`
						LOFNumber                      int         `json:"LOFNumber"`
						MarketedFlightSegment          interface{} `json:"MarketedFlightSegment"`
						NoProtection                   interface{} `json:"NoProtection"`
						Message                        interface{} `json:"Message"`
						OperatingAirline               interface{} `json:"OperatingAirline"`
						OperatingAirlineCode           string      `json:"OperatingAirlineCode"`
						OperatingAirlineName           interface{} `json:"OperatingAirlineName"`
						OffTime                        interface{} `json:"OffTime"`
						OffUTCTime                     interface{} `json:"OffUTCTime"`
						OnTime                         interface{} `json:"OnTime"`
						OnUTCTime                      interface{} `json:"OnUTCTime"`
						OutTime                        interface{} `json:"OutTime"`
						OutUTCTime                     interface{} `json:"OutUTCTime"`
						PlannedTaxiIn                  interface{} `json:"PlannedTaxiIn"`
						PlannedTaxiOut                 interface{} `json:"PlannedTaxiOut"`
						ReasonStatuses                 interface{} `json:"ReasonStatuses"`
						SegmentNumber                  int         `json:"SegmentNumber"`
						Ship                           interface{} `json:"Ship"`
						SpclTerminalMsgsDetails        interface{} `json:"SpclTerminalMsgsDetails"`
						StateProvince                  interface{} `json:"StateProvince"`
						Status                         interface{} `json:"Status"`
						TypeOfSegment                  interface{} `json:"TypeOfSegment"`
						UpgradeProperties              interface{} `json:"UpgradeProperties"`
						UpgradeVisibilityType          int         `json:"UpgradeVisibilityType"`
						UpgradeEligibilityStatus       int         `json:"UpgradeEligibilityStatus"`
						FareBasisCode                  interface{} `json:"FareBasisCode"`
						TicketDesignator               interface{} `json:"TicketDesignator"`
						NotValidAfter                  interface{} `json:"NotValidAfter"`
						OperatingBookingClasses        interface{} `json:"OperatingBookingClasses"`
						OperatingAirlineFlightNumber   interface{} `json:"OperatingAirlineFlightNumber"`
						Legs                           interface{} `json:"Legs"`
						FareComponentSequenceNumber    int         `json:"FareComponentSequenceNumber"`
						IsRulesEligible                interface{} `json:"IsRulesEligible"`
					} `json:"FlightSegment"`
				} `json:"FlightCoupons"`
				TourCode interface{} `json:"TourCode"`
				Payments []struct {
					AccountNumber               interface{} `json:"AccountNumber"`
					AccountNumberEncrypted      interface{} `json:"AccountNumberEncrypted"`
					AccountNumberHMAC           interface{} `json:"AccountNumberHMAC"`
					AccountNumberLastFourDigits interface{} `json:"AccountNumberLastFourDigits"`
					AccountNumberMasked         interface{} `json:"AccountNumberMasked"`
					AccountNumberToken          interface{} `json:"AccountNumberToken"`
					Amount                      float64     `json:"Amount"`
					Authorization               interface{} `json:"Authorization"`
					BillingAddress              struct {
						AddressLines   []string    `json:"AddressLines"`
						Characteristic interface{} `json:"Characteristic"`
						City           string      `json:"City"`
						Country        struct {
							CountryCode      string      `json:"CountryCode"`
							DefaultCurrency  interface{} `json:"DefaultCurrency"`
							DefaultLanguage  interface{} `json:"DefaultLanguage"`
							ISOAlpha3Code    interface{} `json:"ISOAlpha3Code"`
							Name             interface{} `json:"Name"`
							PhoneCountryCode interface{} `json:"PhoneCountryCode"`
							ShortName        interface{} `json:"ShortName"`
							StateCount       int         `json:"StateCount"`
							Status           interface{} `json:"Status"`
						} `json:"Country"`
						DefaultIndicator interface{} `json:"DefaultIndicator"`
						Description      interface{} `json:"Description"`
						DisplaySequence  int         `json:"DisplaySequence"`
						Genre            interface{} `json:"Genre"`
						Key              interface{} `json:"Key"`
						Name             interface{} `json:"Name"`
						PostalCode       string      `json:"PostalCode"`
						Region           interface{} `json:"Region"`
						StateProvince    struct {
							CountryCode       interface{} `json:"CountryCode"`
							Name              interface{} `json:"Name"`
							ShortName         interface{} `json:"ShortName"`
							StateProvinceCode string      `json:"StateProvinceCode"`
						} `json:"StateProvince"`
						Status interface{} `json:"Status"`
					} `json:"BillingAddress"`
					Category               interface{} `json:"Category"`
					Currency               interface{} `json:"Currency"`
					Date                   interface{} `json:"Date"`
					ExchangeRate           interface{} `json:"ExchangeRate"`
					GroupID                interface{} `json:"GroupID"`
					IsLocallyBilled        interface{} `json:"IsLocallyBilled"`
					OperationID            interface{} `json:"OperationID"`
					PaymentCharacteristics interface{} `json:"PaymentCharacteristics"`
					Payor                  interface{} `json:"Payor"`
					Pin                    interface{} `json:"Pin"`
					Type                   interface{} `json:"Type"`
					WalletCategory         int         `json:"WalletCategory"`
				} `json:"Payments"`
			} `json:"Tickets"`
			Person struct {
				GivenName        string `json:"GivenName"`
				Surname          string `json:"Surname"`
				EncryptFirstName string `json:"EncryptFirstName"`
				EncryptLastName  string `json:"EncryptLastName"`
				MiddleName       string `json:"MiddleName"`
				DateOfBirth      string `json:"DateOfBirth"`
				Key              string `json:"Key"`
				ChildIndicator   string `json:"ChildIndicator"`
				InfantIndicator  string `json:"InfantIndicator"`
				PricingPaxType   string `json:"PricingPaxType"`
				Type             string `json:"Type"`
				Sex              string `json:"Sex"`
				Contact          struct {
					Emails []struct {
						Address string `json:"Address"`
					} `json:"Emails"`
					PhoneNumbers []struct {
						PhoneNumber       string `json:"PhoneNumber"`
						CountryAccessCode string `json:"CountryAccessCode"`
						Genre             struct {
							Description string `json:"Description"`
						} `json:"Genre"`
					} `json:"PhoneNumbers"`
				} `json:"Contact"`
				Charges   interface{} `json:"Charges"`
				Documents []struct {
					Genre               interface{} `json:"Genre"`
					RedressNumber       interface{} `json:"RedressNumber"`
					KnownTravelerNumber interface{} `json:"KnownTravelerNumber"`
					Type                string      `json:"Type"`
				} `json:"Documents"`
				CustomerID string `json:"CustomerID"`
				Title      string `json:"Title"`
			} `json:"Person"`
			TicketNumber                 interface{} `json:"TicketNumber"`
			TicketingDate                interface{} `json:"TicketingDate"`
			TravelerNameIndex            interface{} `json:"TravelerNameIndex"`
			ProductLoyaltyProgramProfile interface{} `json:"ProductLoyaltyProgramProfile"`
			Characteristics              []struct {
				Code        string      `json:"Code"`
				Description interface{} `json:"Description"`
				Value       string      `json:"Value"`
				Status      interface{} `json:"Status"`
			} `json:"Characteristics"`
			IsUnaccompaniedMinor interface{} `json:"IsUnaccompaniedMinor"`
		} `json:"Travelers"`
		Sponsor interface{} `json:"Sponsor"`
		Type    []struct {
			Description string      `json:"Description"`
			Key         string      `json:"Key"`
			Value       interface{} `json:"Value"`
		} `json:"Type"`
		UMNRInformation interface{} `json:"UMNRInformation"`
		PNRHistory      interface{} `json:"PNRHistory"`
		CartID          interface{} `json:"CartId"`
	} `json:"Detail"`
	MyInsuranceOffers    interface{} `json:"MyInsuranceOffers"`
	MerchProductOffers   interface{} `json:"MerchProductOffers"`
	PNRChangeEligibility struct {
		IsATREEligible   string        `json:"IsATREEligible"`
		IsPolicyEligible string        `json:"IsPolicyEligible"`
		Policies         []interface{} `json:"Policies"`
	} `json:"PNRChangeEligibility"`
	IsOfferCallNeeded         string `json:"IsOfferCallNeeded"`
	IsTripInsuranceCallNeeded bool   `json:"IsTripInsuranceCallNeeded"`
}

type TokenResponse struct {
	Data struct {
		Token struct {
			Hash      string    `json:"hash"`
			ExpiresAt time.Time `json:"expiresAt"`
		} `json:"token"`
	} `json:"data"`
	Valid           interface{} `json:"valid"`
	TokenExpiration interface{} `json:"TokenExpiration"`
}
