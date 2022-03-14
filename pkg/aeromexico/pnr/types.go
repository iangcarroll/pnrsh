package pnr

type PNR struct {
	Remarks    []Remark
	Flights    []Flight
	Passengers []Passenger
	Tickets    []Ticket
	SMCalcLink string
}

type Remark struct {
	FreeFormText string
	RemarkType   string
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
	Name                   string
	Status                 string
	OverbookingEligible    bool
	SkyPriority            bool
	BenefitCodes           string
	FeeForSeatSelection    bool
	FeeForAmPlusUpgrade    bool
	FeeForPreferredUpgrade bool
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

type ManagePnrResponse struct {
	Meta struct {
		Class string `json:"class"`
	} `json:"_meta"`
	Collection []struct {
		Cart struct {
			Meta struct {
				Class                       string `json:"class"`
				CartID                      string `json:"cartId"`
				CartExpirationTimeInSeconds int    `json:"cartExpirationTimeInSeconds"`
				Total                       struct {
					Currency struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						CurrencyCode string `json:"currencyCode"`
						Base         int    `json:"base"`
						TaxDetails   struct {
						} `json:"taxDetails"`
						TotalTax             int `json:"totalTax"`
						TaxDetailsRefundable struct {
						} `json:"taxDetailsRefundable"`
						TotalTaxRefundable int `json:"totalTaxRefundable"`
						SubTotal           int `json:"subTotal"`
						Total              int `json:"total"`
						TotalCPPoints      int `json:"totalCPPoints"`
					} `json:"currency"`
					Points int `json:"points"`
				} `json:"total"`
			} `json:"_meta"`
			TravelersCode string `json:"travelersCode"`
			TravelerInfo  struct {
				Meta struct {
					Class string `json:"class"`
				} `json:"_meta"`
				Collection []struct {
					Meta struct {
						Class string `json:"class"`
					} `json:"_meta"`
					CheckinStatus                bool          `json:"checkinStatus"`
					CheckinStatusBySegment       interface{}   `json:"checkinStatusBySegment"`
					ShowEmailCapture             bool          `json:"showEmailCapture"`
					IsEligibleToCheckin          bool          `json:"isEligibleToCheckin"`
					IsSelectedToCheckin          bool          `json:"isSelectedToCheckin"`
					MissingCheckinRequiredFields []interface{} `json:"missingCheckinRequiredFields"`
					Infant                       interface{}   `json:"infant"`
					PaxType                      string        `json:"paxType"`
					CouponCode                   interface{}   `json:"couponCode"`
					TicketNumber                 string        `json:"ticketNumber"`
					Token                        interface{}   `json:"token"`
					AllTicketNumbers             []string      `json:"allTicketNumbers"`
					AmountTicket                 struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						CurrencyCode string `json:"currencyCode"`
						Base         int    `json:"base"`
						TaxDetails   struct {
						} `json:"taxDetails"`
						TotalTax             int `json:"totalTax"`
						TaxDetailsRefundable struct {
						} `json:"taxDetailsRefundable"`
						TotalTaxRefundable int     `json:"totalTaxRefundable"`
						SubTotal           int     `json:"subTotal"`
						Total              float64 `json:"total"`
						TotalCPPoints      int     `json:"totalCPPoints"`
					} `json:"AmountTicket"`
					TicketNumbers []struct {
						Number                string `json:"number"`
						RelatedDocumentNumber string `json:"relatedDocumentNumber"`
						Coupon                string `json:"coupon"`
						Status                string `json:"status"`
						PreviousStatus        string `json:"previousStatus"`
						FlightNumber          string `json:"flightNumber"`
						StartDate             string `json:"startDate"`
						StartLocation         string `json:"startLocation"`
						EndLocation           string `json:"endLocation"`
					} `json:"ticketNumbers"`
					BenefitRedeemedCollection struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						Collection []interface{} `json:"_collection"`
					} `json:"benefitRedeemedCollection"`
					Ancillaries struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						Collection []interface{} `json:"_collection"`
					} `json:"ancillaries"`
					SegmentChoices struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						Collection []struct {
							PaidSeat          bool        `json:"paidSeat"`
							SeatCost          interface{} `json:"seatCost"`
							ChangeSeatAllowed bool        `json:"changeSeatAllowed"`
							IsGenericSeat     bool        `json:"isGenericSeat"`
							Meta              struct {
								Class string `json:"class"`
							} `json:"_meta"`
							Seat struct {
								SubjectToSeatSelectionFee bool `json:"subjectToSeatSelectionFee"`
								Meta                      struct {
									Class string `json:"class"`
								} `json:"_meta"`
								Code                string        `json:"code"`
								SeatCharacteristics []interface{} `json:"seatCharacteristics"`
								SeatFacilities      interface{}   `json:"seatFacilities"`
							} `json:"seat"`
							FareType    interface{} `json:"fareType"`
							Status      interface{} `json:"status"`
							SegmentCode string      `json:"segmentCode"`
						} `json:"_collection"`
					} `json:"segmentChoices"`
					BookingClasses struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						Collection []struct {
							SegmentCode  string `json:"segmentCode"`
							BookingClass string `json:"bookingClass"`
							BookingCabin string `json:"bookingCabin"`
							PassengerID  string `json:"passengerId"`
							Flight       string `json:"flight"`
						} `json:"_collection"`
					} `json:"bookingClasses"`
					ItineraryBookingClasses struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						Collection []struct {
							SegmentCode  string `json:"segmentCode"`
							BookingClass string `json:"bookingClass"`
							BookingCabin string `json:"bookingCabin"`
							PassengerID  string `json:"passengerId"`
							Flight       string `json:"flight"`
						} `json:"_collection"`
					} `json:"itineraryBookingClasses"`
					FareBasisCodes struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						Collection []struct {
							SegmentCode   string `json:"segmentCode"`
							FareBasisCode string `json:"fareBasisCode"`
						} `json:"_collection"`
					} `json:"fareBasisCodes"`
					CouponNumbers struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						Collection []struct {
							SegmentCode  string `json:"segmentCode"`
							CouponNumber int    `json:"couponNumber"`
						} `json:"_collection"`
					} `json:"couponNumbers"`
					EditCodesBySegment          interface{}   `json:"editCodesBySegment"`
					IneligibleReasons           []interface{} `json:"ineligibleReasons"`
					NameRefNumber               string        `json:"nameRefNumber"`
					DigitalSignatureStatus      interface{}   `json:"digitalSignatureStatus"`
					PassengerType               string        `json:"passengerType"`
					CheckedBags                 interface{}   `json:"checkedBags"`
					CheckinNumber               interface{}   `json:"checkinNumber"`
					Selectee                    bool          `json:"selectee"`
					OverBookingInfo             interface{}   `json:"overBookingInfo"`
					PriorityCode                interface{}   `json:"priorityCode"`
					UpgradeCode                 interface{}   `json:"upgradeCode"`
					PriorityBySegmentCollection interface{}   `json:"priorityBySegmentCollection"`
					CandidateForUpgradeList     bool          `json:"candidateForUpgradeList"`
					IsTitular                   bool          `json:"isTitular"`
					IsCompanion                 bool          `json:"isCompanion"`
					OnStandByList               bool          `json:"onStandByList"`
					OnUpgradeList               bool          `json:"onUpgradeList"`
					ChangeSeatAllowed           bool          `json:"changeSeatAllowed"`
					MealSelectionAllowed        interface{}   `json:"mealSelectionAllowed"`
					MealSelected                interface{}   `json:"mealSelected"`
					ManageUserProfileAccount    struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						Account struct {
							Meta struct {
								Class string `json:"class"`
							} `json:"_meta"`
							Balance              interface{} `json:"balance"`
							Tier                 interface{} `json:"tier"`
							TierExpirationDate   interface{} `json:"tierExpirationDate"`
							FrequentFlyerNumber  interface{} `json:"frequentFlyerNumber"`
							FrequentFlyerProgram interface{} `json:"frequentFlyerProgram"`
							Login                bool        `json:"login"`
						} `json:"account"`
					} `json:"ManageUserProfileAccount"`
					ChubbInsuranceUSA []struct {
						State        string `json:"state"`
						AviableChubb bool   `json:"aviableChubb"`
					} `json:"ChubbInsuranceUSA"`
					ClassificationPaxType string      `json:"classificationPaxType"`
					DisplayName           string      `json:"displayName"`
					FirstName             string      `json:"firstName"`
					MiddleName            interface{} `json:"middleName"`
					LastName              string      `json:"lastName"`
					Prefix                interface{} `json:"prefix"`
					ID                    string      `json:"id"`
					Phones                struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						Collection []interface{} `json:"_collection"`
					} `json:"phones"`
					Email       string   `json:"email"`
					SavedEmails []string `json:"savedEmails"`
					SsrRemarks  struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						SsrCodes []interface{} `json:"ssrCodes"`
					} `json:"ssrRemarks"`
					FreeBaggageAllowance struct {
						Personal          int     `json:"personal"`
						CarryOn           int     `json:"carryOn"`
						CarryOnWeightKG   int     `json:"carryOnWeightKG"`
						CarryOnWeightLB   float64 `json:"carryOnWeightLB"`
						Checked           int     `json:"checked"`
						CheckedMaxWeight  int     `json:"checkedMaxWeight"`
						CheckedWeightUnit string  `json:"checkedWeightUnit"`
					} `json:"freeBaggageAllowance"`
					FreeBaggageAllowancePerLegsArray []struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						LegCode              string `json:"legCode"`
						FreeBaggageAllowance struct {
							Personal                int     `json:"personal"`
							CarryOn                 int     `json:"carryOn"`
							CarryOnWeightKG         int     `json:"carryOnWeightKG"`
							CarryOnWeightLB         float64 `json:"carryOnWeightLB"`
							Checked                 int     `json:"checked"`
							CheckedWeightCollection []struct {
								CheckedFreeWeight float64 `json:"checkedFreeWeight"`
								CheckedMaxWeight  float64 `json:"checkedMaxWeight"`
								CheckedWeightUnit string  `json:"checkedWeightUnit"`
							} `json:"checkedWeightCollection"`
							CarryOnWeightCollection []struct {
								CarryOnWeight     float64 `json:"carryOnWeight"`
								CarryOnWeightUnit string  `json:"carryOnWeightUnit"`
							} `json:"carryOnWeightCollection"`
							AllowSaleBags bool `json:"allowSaleBags"`
						} `json:"freeBaggageAllowance"`
					} `json:"freeBaggageAllowancePerLegsArray"`
					Gender                        string      `json:"gender"`
					DateOfBirth                   string      `json:"dateOfBirth"`
					FrequentFlyerProgram          string      `json:"frequentFlyerProgram"`
					FrequentFlyerNumber           string      `json:"frequentFlyerNumber"`
					TierLevel                     interface{} `json:"tierLevel"`
					CorporateFrequentFlyerProgram interface{} `json:"corporateFrequentFlyerProgram"`
					CorporateFrequentFlyerNumber  interface{} `json:"corporateFrequentFlyerNumber"`
					Loyalty                       []struct {
						SupplierCode    string      `json:"supplierCode"`
						Number          string      `json:"number"`
						TierLevelNumber int         `json:"tierLevelNumber"`
						TierTag         string      `json:"tierTag"`
						TierPriority    interface{} `json:"tierPriority"`
					} `json:"loyalty"`
					RedressNumber             interface{} `json:"redressNumber"`
					Ktn                       interface{} `json:"ktn"`
					Ctn                       interface{} `json:"ctn"`
					TravelDocument            interface{} `json:"travelDocument"`
					TimaticInfo               interface{} `json:"timaticInfo"`
					TimaticProvided           bool        `json:"timaticProvided"`
					VisaInfo                  interface{} `json:"visaInfo"`
					PermanentResidentCardInfo interface{} `json:"permanentResidentCardInfo"`
					VisaInfoAlreadySent       bool        `json:"visaInfoAlreadySent"`
					EtaAlreadySent            bool        `json:"etaAlreadySent"`
					ReturningDateAlreadySent  bool        `json:"returningDateAlreadySent"`
					CountryOfResidence        interface{} `json:"countryOfResidence"`
					DestinationAddress        interface{} `json:"destinationAddress"`
					DeclinedEmergencyContact  bool        `json:"declinedEmergencyContact"`
					EmergencyContact          interface{} `json:"emergencyContact"`
					OnwardTravelDate          interface{} `json:"onwardTravelDate"`
					Benefit                   struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						Collection []struct {
							Meta struct {
								Class string `json:"class"`
							} `json:"_meta"`
							Code         string `json:"code"`
							SeatBenefits struct {
								Meta struct {
									Class string `json:"class"`
								} `json:"_meta"`
								Reason      string `json:"reason"`
								FeeRequired struct {
									SeatSelection    bool `json:"seatSelection"`
									AmPlusUgrade     bool `json:"amPlusUgrade"`
									PrefferedUpgrade bool `json:"prefferedUpgrade"`
								} `json:"feeRequired"`
								FeeDiscounted struct {
									SeatSelection    float64 `json:"seatSelection"`
									AmPlusUgrade     float64 `json:"amPlusUgrade"`
									PreferredUpgrade float64 `json:"preferredUpgrade"`
								} `json:"feeDiscounted"`
							} `json:"seatBenefits"`
							BagBenefit        interface{} `json:"bagBenefit"`
							BagBenefitCobrand interface{} `json:"bagBenefitCobrand"`
							LegCodeBenefits   interface{} `json:"legCodeBenefits"`
						} `json:"_collection"`
					} `json:"benefit"`
					IsOverBookingEligible       bool        `json:"isOverBookingEligible"`
					IsSkyPriority               bool        `json:"isSkyPriority"`
					RequiresUpdateProfile       bool        `json:"requiresUpdateProfile"`
					RequiresAddProfile          bool        `json:"requiresAddProfile"`
					Nickname                    interface{} `json:"nickname"`
					BenefitCorporateRecognition struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						PreferredSeats bool        `json:"preferredSeats"`
						Clid           interface{} `json:"clid"`
						SeatCode       interface{} `json:"seatCode"`
						PreferredBags  bool        `json:"preferredBags"`
					} `json:"benefitCorporateRecognition"`
					CovidRestrictionConfirmed bool        `json:"covidRestrictionConfirmed"`
					StateCode                 interface{} `json:"stateCode"`
					CountryCode               interface{} `json:"countryCode"`
					IsLiveUSA                 bool        `json:"isLiveUSA"`
					BaggageRedemption         interface{} `json:"baggageRedemption"`
				} `json:"_collection"`
			} `json:"travelerInfo"`
			CorporateFrequentFlyerNumber interface{} `json:"corporateFrequentFlyerNumber"`
			FingerprintSessionID         interface{} `json:"fingerprintSessionId"`
			Warnings                     struct {
				Meta struct {
					Class string `json:"class"`
				} `json:"_meta"`
				Collection []interface{} `json:"_collection"`
			} `json:"_warnings"`
			EmdVoucherRs interface{} `json:"emdVoucherRs"`
			UpsellOffers struct {
				Meta struct {
					Class string `json:"class"`
				} `json:"_meta"`
				Collection []struct {
					SegmentCode      string      `json:"segmentCode"`
					UpgradeAvailable bool        `json:"upgradeAvailable"`
					Currency         interface{} `json:"currency"`
					SeatsRemaining   interface{} `json:"seatsRemaining"`
					Ancillary        interface{} `json:"ancillary"`
					Quantity         int         `json:"quantity"`
					LegCode          string      `json:"legCode"`
					SegmentCodeAux   interface{} `json:"segmentCodeAux"`
					Meta             struct {
						Class string `json:"class"`
					} `json:"_meta"`
				} `json:"_collection"`
			} `json:"upsellOffers"`
			RedemptionAncilaryList []interface{} `json:"redemptionAncilaryList"`
		} `json:"cart"`
		ReservationStatus        string      `json:"reservationStatus"`
		PricePerPassengerTypeMap interface{} `json:"pricePerPassengerTypeMap"`
		HoursToExpire            int         `json:"hoursToExpire"`
		ManageStatus             string      `json:"manageStatus"`
		CheckinStatus            string      `json:"checkinStatus"`
		Remarks                  []string    `json:"remarks"`
		Clid                     struct {
			Meta struct {
				Class string `json:"class"`
			} `json:"_meta"`
			CompanyCode           interface{} `json:"companyCode"`
			CompanyNumber         interface{} `json:"companyNumber"`
			ApplyCorporateBenefit bool        `json:"applyCorporateBenefit"`
			Message               string      `json:"message"`
			IsPreferentBag        bool        `json:"isPreferentBag"`
		} `json:"clid"`
		ElegibilityEMD struct {
			Meta struct {
				Class string `json:"class"`
			} `json:"_meta"`
			EmdVoucherFull     bool `json:"emdVoucherFull"`
			EmdVoucherExcedeed bool `json:"emdVoucherExcedeed"`
			EmdVoucherMinimun  bool `json:"emdVoucherMinimun"`
		} `json:"elegibilityEMD"`
		Meta struct {
			Class string `json:"class"`
		} `json:"_meta"`
		Pnr          string      `json:"pnr"`
		CodeSharePnr interface{} `json:"codeSharePnr"`
		Legs         struct {
			Meta struct {
				Class string `json:"class"`
			} `json:"_meta"`
			Collection []struct {
				Meta struct {
					Class string `json:"class"`
				} `json:"_meta"`
				EstimatedDepartureTime string      `json:"estimatedDepartureTime"`
				ScheduledDepartureTime string      `json:"scheduledDepartureTime"`
				BoardingTime           interface{} `json:"boardingTime"`
				RemainingTimeToBoard   int         `json:"remainingTimeToBoard"`
				BoardingTerminal       interface{} `json:"boardingTerminal"`
				BoardingGate           interface{} `json:"boardingGate"`
				Cabin                  interface{} `json:"cabin"`
				IsChangeAvailable      bool        `json:"isChangeAvailable"`
				FlightStatus           interface{} `json:"flightStatus"`
				CheckinStatus          string      `json:"checkinStatus"`
				OriginalCheckinStatus  interface{} `json:"originalCheckinStatus"`
				ManageStatus           string      `json:"manageStatus"`
				Segments               struct {
					Meta struct {
						Class string `json:"class"`
					} `json:"_meta"`
					Collection []struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						CheckinIneligibleReasons []interface{} `json:"checkinIneligibleReasons"`
						Segment                  struct {
							Meta struct {
								Class string `json:"class"`
							} `json:"_meta"`
							SegmentCode                    string        `json:"segmentCode"`
							FareBasisCode                  string        `json:"fareBasisCode"`
							FareFamily                     string        `json:"fareFamily"`
							FareAndCabinName               string        `json:"fareAndCabinName"`
							DepartureAirport               string        `json:"departureAirport"`
							DepartureTerminal              string        `json:"departureTerminal"`
							DepartureDateTime              string        `json:"departureDateTime"`
							TimeZoneName                   string        `json:"timeZoneName"`
							TimeZoneID                     string        `json:"timeZoneId"`
							DstOffset                      int           `json:"dstOffset"`
							RawOffset                      int           `json:"rawOffset"`
							ArrivalTimeZoneName            string        `json:"arrivalTimeZoneName"`
							ArrivalTimeZoneID              string        `json:"arrivalTimeZoneId"`
							ArrivalDstOffset               int           `json:"arrivalDstOffset"`
							ArrivalRawOffset               int           `json:"arrivalRawOffset"`
							ArrivalAirport                 string        `json:"arrivalAirport"`
							ArrivalTerminal                string        `json:"arrivalTerminal"`
							ArrivalDateTime                string        `json:"arrivalDateTime"`
							Stops                          []interface{} `json:"stops"`
							FlightDurationInMinutes        int           `json:"flightDurationInMinutes"`
							LayoverToNextSegmentsInMinutes int           `json:"layoverToNextSegmentsInMinutes"`
							AircraftType                   string        `json:"aircraftType"`
							OperatingCarrier               string        `json:"operatingCarrier"`
							OperatingFlightCode            string        `json:"operatingFlightCode"`
							MarketingCarrier               string        `json:"marketingCarrier"`
							MarketingFlightCode            string        `json:"marketingFlightCode"`
							AmenitiesPremiumCabin          struct {
								ComplimentaryFeatures []string      `json:"complimentaryFeatures"`
								PaidFeatures          []interface{} `json:"paidFeatures"`
							} `json:"amenitiesPremiumCabin"`
							AmenitiesMainCabin struct {
								ComplimentaryFeatures []string      `json:"complimentaryFeatures"`
								PaidFeatures          []interface{} `json:"paidFeatures"`
							} `json:"amenitiesMainCabin"`
							IsPremierAvailable     bool        `json:"isPremierAvailable"`
							SegmentStatus          string      `json:"segmentStatus"`
							BoardingTime           interface{} `json:"boardingTime"`
							BoardingTerminal       interface{} `json:"boardingTerminal"`
							BoardingGate           interface{} `json:"boardingGate"`
							Capacity               interface{} `json:"capacity"`
							BookingClass           string      `json:"bookingClass"`
							Cabin                  string      `json:"cabin"`
							FlightStatus           interface{} `json:"flightStatus"`
							CheckinStatus          interface{} `json:"checkinStatus"`
							EstimatedDepartureTime interface{} `json:"estimatedDepartureTime"`
							ScheduledDepartureTime interface{} `json:"scheduledDepartureTime"`
							EstimatedArrivalTime   interface{} `json:"estimatedArrivalTime"`
							ScheduledArrivalTime   interface{} `json:"scheduledArrivalTime"`
							Domestic               bool        `json:"domestic"`
							CodeShare              bool        `json:"codeShare"`
							DaysBeforeDeparture    interface{} `json:"daysBeforeDeparture"`
							PassFlight             bool        `json:"passFlight"`
							IsPremierLight         bool        `json:"isPremierLight"`
						} `json:"segment"`
						FlagFlightCanada bool `json:"flagFlightCanada"`
					} `json:"_collection"`
					LegCode                  string      `json:"legCode"`
					TripID                   interface{} `json:"tripId"`
					LegType                  string      `json:"legType"`
					TotalFlightTimeInMinutes int         `json:"totalFlightTimeInMinutes"`
					IsLeavingUk              bool        `json:"isLeavingUk"`
				} `json:"segments"`
				Domestic                             bool        `json:"domestic"`
				CarryOn                              bool        `json:"carryOn"`
				StandByReservation                   bool        `json:"standByReservation"`
				UpgradeReservation                   bool        `json:"upgradeReservation"`
				ApplyForUpgrade                      bool        `json:"applyForUpgrade"`
				ShowUpgradeList                      bool        `json:"showUpgradeList"`
				GroundHandled                        bool        `json:"groundHandled"`
				ClosedForEarlyChk                    bool        `json:"closedForEarlyChk"`
				PriorityReservation                  bool        `json:"priorityReservation"`
				OnStandByList                        bool        `json:"onStandByList"`
				OnUpgradeList                        bool        `json:"onUpgradeList"`
				PaxOnStandByList                     int         `json:"paxOnStandByList"`
				PaxOnUpgradeList                     int         `json:"paxOnUpgradeList"`
				StandByListPaxLimit                  int         `json:"standByListPaxLimit"`
				CheckInRestrictedByStandbyConstraint bool        `json:"checkInRestrictedByStandbyConstraint"`
				RegionKey                            string      `json:"regionKey"`
				MarketFlight                         string      `json:"marketFlight"`
				SeatSelectionType                    string      `json:"seatSelectionType"`
				FirstBagPolicyType                   interface{} `json:"firstBagPolicyType"`
				ItineraryPartType                    string      `json:"itineraryPartType"`
				LookUpTime                           int         `json:"lookUpTime"`
				IsGroundedReservation                bool        `json:"isGroundedReservation"`
				SeamlessCheckin                      bool        `json:"seamlessCheckin"`
				OperatingHandler                     interface{} `json:"operatingHandler"`
				OvsReservation                       bool        `json:"ovsReservation"`
			} `json:"_collection"`
		} `json:"legs"`
		LegsOrigin struct {
			Meta struct {
				Class string `json:"class"`
			} `json:"_meta"`
			Collection []struct {
				Meta struct {
					Class string `json:"class"`
				} `json:"_meta"`
				EstimatedDepartureTime string      `json:"estimatedDepartureTime"`
				ScheduledDepartureTime string      `json:"scheduledDepartureTime"`
				BoardingTime           interface{} `json:"boardingTime"`
				RemainingTimeToBoard   int         `json:"remainingTimeToBoard"`
				BoardingTerminal       interface{} `json:"boardingTerminal"`
				BoardingGate           interface{} `json:"boardingGate"`
				Cabin                  interface{} `json:"cabin"`
				IsChangeAvailable      bool        `json:"isChangeAvailable"`
				FlightStatus           interface{} `json:"flightStatus"`
				CheckinStatus          string      `json:"checkinStatus"`
				OriginalCheckinStatus  interface{} `json:"originalCheckinStatus"`
				ManageStatus           string      `json:"manageStatus"`
				Segments               struct {
					Meta struct {
						Class string `json:"class"`
					} `json:"_meta"`
					Collection []struct {
						Meta struct {
							Class string `json:"class"`
						} `json:"_meta"`
						CheckinIneligibleReasons []interface{} `json:"checkinIneligibleReasons"`
						Segment                  struct {
							Meta struct {
								Class string `json:"class"`
							} `json:"_meta"`
							SegmentCode                    string        `json:"segmentCode"`
							FareBasisCode                  string        `json:"fareBasisCode"`
							FareFamily                     string        `json:"fareFamily"`
							FareAndCabinName               string        `json:"fareAndCabinName"`
							DepartureAirport               string        `json:"departureAirport"`
							DepartureTerminal              string        `json:"departureTerminal"`
							DepartureDateTime              string        `json:"departureDateTime"`
							TimeZoneName                   string        `json:"timeZoneName"`
							TimeZoneID                     string        `json:"timeZoneId"`
							DstOffset                      int           `json:"dstOffset"`
							RawOffset                      int           `json:"rawOffset"`
							ArrivalTimeZoneName            string        `json:"arrivalTimeZoneName"`
							ArrivalTimeZoneID              string        `json:"arrivalTimeZoneId"`
							ArrivalDstOffset               int           `json:"arrivalDstOffset"`
							ArrivalRawOffset               int           `json:"arrivalRawOffset"`
							ArrivalAirport                 string        `json:"arrivalAirport"`
							ArrivalTerminal                string        `json:"arrivalTerminal"`
							ArrivalDateTime                string        `json:"arrivalDateTime"`
							Stops                          []interface{} `json:"stops"`
							FlightDurationInMinutes        int           `json:"flightDurationInMinutes"`
							LayoverToNextSegmentsInMinutes int           `json:"layoverToNextSegmentsInMinutes"`
							AircraftType                   string        `json:"aircraftType"`
							OperatingCarrier               string        `json:"operatingCarrier"`
							OperatingFlightCode            string        `json:"operatingFlightCode"`
							MarketingCarrier               string        `json:"marketingCarrier"`
							MarketingFlightCode            string        `json:"marketingFlightCode"`
							AmenitiesPremiumCabin          struct {
								ComplimentaryFeatures []string      `json:"complimentaryFeatures"`
								PaidFeatures          []interface{} `json:"paidFeatures"`
							} `json:"amenitiesPremiumCabin"`
							AmenitiesMainCabin struct {
								ComplimentaryFeatures []string      `json:"complimentaryFeatures"`
								PaidFeatures          []interface{} `json:"paidFeatures"`
							} `json:"amenitiesMainCabin"`
							IsPremierAvailable     bool        `json:"isPremierAvailable"`
							SegmentStatus          string      `json:"segmentStatus"`
							BoardingTime           interface{} `json:"boardingTime"`
							BoardingTerminal       interface{} `json:"boardingTerminal"`
							BoardingGate           interface{} `json:"boardingGate"`
							Capacity               interface{} `json:"capacity"`
							BookingClass           string      `json:"bookingClass"`
							Cabin                  string      `json:"cabin"`
							FlightStatus           interface{} `json:"flightStatus"`
							CheckinStatus          interface{} `json:"checkinStatus"`
							EstimatedDepartureTime interface{} `json:"estimatedDepartureTime"`
							ScheduledDepartureTime interface{} `json:"scheduledDepartureTime"`
							EstimatedArrivalTime   interface{} `json:"estimatedArrivalTime"`
							ScheduledArrivalTime   interface{} `json:"scheduledArrivalTime"`
							Domestic               bool        `json:"domestic"`
							CodeShare              bool        `json:"codeShare"`
							DaysBeforeDeparture    interface{} `json:"daysBeforeDeparture"`
							PassFlight             bool        `json:"passFlight"`
							IsPremierLight         bool        `json:"isPremierLight"`
						} `json:"segment"`
						FlagFlightCanada bool `json:"flagFlightCanada"`
					} `json:"_collection"`
					LegCode                  string      `json:"legCode"`
					TripID                   interface{} `json:"tripId"`
					LegType                  string      `json:"legType"`
					TotalFlightTimeInMinutes int         `json:"totalFlightTimeInMinutes"`
					IsLeavingUk              bool        `json:"isLeavingUk"`
				} `json:"segments"`
				Domestic                             bool        `json:"domestic"`
				CarryOn                              bool        `json:"carryOn"`
				StandByReservation                   bool        `json:"standByReservation"`
				UpgradeReservation                   bool        `json:"upgradeReservation"`
				ApplyForUpgrade                      bool        `json:"applyForUpgrade"`
				ShowUpgradeList                      bool        `json:"showUpgradeList"`
				GroundHandled                        bool        `json:"groundHandled"`
				ClosedForEarlyChk                    bool        `json:"closedForEarlyChk"`
				PriorityReservation                  bool        `json:"priorityReservation"`
				OnStandByList                        bool        `json:"onStandByList"`
				OnUpgradeList                        bool        `json:"onUpgradeList"`
				PaxOnStandByList                     int         `json:"paxOnStandByList"`
				PaxOnUpgradeList                     int         `json:"paxOnUpgradeList"`
				StandByListPaxLimit                  int         `json:"standByListPaxLimit"`
				CheckInRestrictedByStandbyConstraint bool        `json:"checkInRestrictedByStandbyConstraint"`
				RegionKey                            string      `json:"regionKey"`
				MarketFlight                         string      `json:"marketFlight"`
				SeatSelectionType                    interface{} `json:"seatSelectionType"`
				FirstBagPolicyType                   interface{} `json:"firstBagPolicyType"`
				ItineraryPartType                    string      `json:"itineraryPartType"`
				LookUpTime                           int         `json:"lookUpTime"`
				IsGroundedReservation                bool        `json:"isGroundedReservation"`
				SeamlessCheckin                      bool        `json:"seamlessCheckin"`
				OperatingHandler                     interface{} `json:"operatingHandler"`
				OvsReservation                       bool        `json:"ovsReservation"`
			} `json:"_collection"`
		} `json:"legsOrigin"`
		CreationDate        string `json:"creationDate"`
		TicketDocumentDate  string `json:"ticketDocumentDate"`
		FlightItineraryType string `json:"flightItineraryType"`
		GlobalMarketFlight  string `json:"globalMarketFlight"`
	} `json:"_collection"`
	Warnings struct {
		Meta struct {
			Class string `json:"class"`
		} `json:"_meta"`
		Collection []interface{} `json:"_collection"`
	} `json:"_warnings"`
	CreationDate int64 `json:"creationDate"`
}
