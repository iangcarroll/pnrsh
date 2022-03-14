package pnr

import "encoding/xml"

type PNR struct {
	Remarks    []Remark
	Flights    []Flight
	Passengers []Passenger
	Flags      []Flag
	Tickets    []Ticket
	Fare       Fare
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
	UpgradeStatus          string
	ScheduledDeparture     string
	ScheduledArrival       string
	FlightNumber           string
}

type Passenger struct {
	Name       string
	CustomerID string
	CheckedIn  bool
	SSRs       []SSR
	Status     string
}

type SSR struct {
	AirlineCode string
	Type        string
	Remark      string
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
	NumCoupons             uint64
	ValidatedAgainstCoupon bool
}

type Fare struct {
	BaseCurrencyCode string
	BaseFare         string
	TotalTax         string
	TotalFare        string
	TaxRows          []struct {
		TaxType           string
		Amount            string
		Currency          string
		CarrierImposedFee bool
	}
	TotalCurrencyCode string
	FareBasisCode     string
	EstimatedMQD      string
	SMCalcLink        string
}

// Raw API response from Delta.
type RetrievePnrResponse struct {
	XMLName                   xml.Name `xml:"retrievePnrResponse"`
	Text                      string   `xml:",chardata"`
	Ns2                       string   `xml:"ns2,attr"`
	Status                    string   `xml:"status"`
	EconomyComfortEligibility struct {
		Text  string `xml:",chardata"`
		EcPnr struct {
			Text   string `xml:",chardata"`
			RecLoc string `xml:"recLoc,attr"`
			EcItin []struct {
				Text     string `xml:",chardata"`
				Eligible string `xml:"eligible,attr"`
			} `xml:"ecItin"`
		} `xml:"ecPnr"`
	} `xml:"economyComfortEligibility"`
	Links struct {
		Text    string `xml:",chardata"`
		Flights struct {
			Text       string `xml:",chardata"`
			FlightLink struct {
				Text               string `xml:",chardata"`
				ConfirmationNumber string `xml:"confirmationNumber"`
				DeepLinks          struct {
					Text         string `xml:",chardata"`
					DeepLinkData []struct {
						Text     string `xml:",chardata"`
						Key      string `xml:"key"`
						PostData struct {
							Text             string `xml:",chardata"`
							Xsi              string `xml:"xsi,attr"`
							Type             string `xml:"type,attr"`
							FirstName        string `xml:"firstName"`
							LastName         string `xml:"lastName"`
							ConfirmationCode string `xml:"confirmationCode"`
						} `xml:"postData"`
						URL string `xml:"url"`
					} `xml:"deepLinkData"`
				} `xml:"deepLinks"`
			} `xml:"flightLink"`
		} `xml:"flights"`
		Ignore string `xml:"ignore"`
	} `xml:"links"`
	LoyaltySummaries struct {
		Text           string `xml:",chardata"`
		LoyaltySummary struct {
			Text               string `xml:",chardata"`
			ConfirmationNumber string `xml:"confirmationNumber"`
			LoyaltyPassenger   []struct {
				Text              string `xml:",chardata"`
				FirstNIN          string `xml:"firstNIN"`
				LastNIN           string `xml:"lastNIN"`
				LoyaltyLineNumber string `xml:"loyaltyLineNumber"`
				LoyaltyNumber     string `xml:"loyaltyNumber"`
				PassengerIndex    string `xml:"passengerIndex"`
			} `xml:"loyaltyPassenger"`
		} `xml:"loyaltySummary"`
	} `xml:"loyaltySummaries"`
	ModifyFlightsPayloads struct {
		Text              string `xml:",chardata"`
		ModifyFlightsPnrs struct {
			Text               string `xml:",chardata"`
			ConfirmationNumber string `xml:"confirmationNumber,attr"`
			Payload            string `xml:"payload,attr"`
		} `xml:"modifyFlightsPnrs"`
	} `xml:"modifyFlightsPayloads"`
	PnrData struct {
		Text     string `xml:",chardata"`
		PnrDatas struct {
			Text               string `xml:",chardata"`
			ConfirmationNumber string `xml:"confirmationNumber,attr"`
			Encrypted          string `xml:"encrypted,attr"`
			Token              string `xml:"token,attr"`
		} `xml:"pnrDatas"`
	} `xml:"pnrData"`
	TripsResponse struct {
		Text    string `xml:",chardata"`
		Journey struct {
			Text   string `xml:",chardata"`
			CartId string `xml:"cartId"`
			Pnr    struct {
				Text               string `xml:",chardata"`
				ConfirmationNumber string `xml:"confirmationNumber"`
				PnrType            string `xml:"pnrType"`
				Remarks            struct {
					Text             string `xml:",chardata"`
					DomainObjectList struct {
						Text         string `xml:",chardata"`
						DomainObject []struct {
							Text         string `xml:",chardata"`
							Xsi          string `xml:"xsi,attr"`
							Type         string `xml:"type,attr"`
							FreeFormText string `xml:"freeFormText"`
							RemarkType   string `xml:"remarkType"`
							RemarkLevel  string `xml:"remarkLevel"`
						} `xml:"domainObject"`
					} `xml:"domainObjectList"`
					Index string `xml:"index"`
				} `xml:"remarks"`
				Passengers struct {
					Text             string `xml:",chardata"`
					DomainObjectList struct {
						Text         string `xml:",chardata"`
						DomainObject []struct {
							Text                string `xml:",chardata"`
							Xsi                 string `xml:"xsi,attr"`
							Type                string `xml:"type,attr"`
							SecurityQAStoredInd string `xml:"securityQAStoredInd"`
							SqaLocked           string `xml:"sqaLocked"`
							AddrLocked          string `xml:"addrLocked"`
							PnrName             struct {
								Text      string `xml:",chardata"`
								FirstName string `xml:"firstName"`
								LastName  string `xml:"lastName"`
							} `xml:"pnrName"`
							Name struct {
								Text      string `xml:",chardata"`
								FirstName string `xml:"firstName"`
								LastName  string `xml:"lastName"`
							} `xml:"name"`
							LoyaltyAccounts struct {
								Text             string `xml:",chardata"`
								DomainObjectList struct {
									Text         string `xml:",chardata"`
									DomainObject struct {
										Text                       string `xml:",chardata"`
										Type                       string `xml:"type,attr"`
										AirlineCode                string `xml:"airlineCode"`
										ProgramName                string `xml:"programName"`
										ID                         string `xml:"id"`
										AliasAcctFoundInd          string `xml:"aliasAcctFoundInd"`
										MembershipStatusCd         string `xml:"membershipStatusCd"`
										MembershipStatusDesc       string `xml:"membershipStatusDesc"`
										PinFoundInd                string `xml:"pinFoundInd"`
										PayWithMilesEligibilityInd string `xml:"payWithMilesEligibilityInd"`
									} `xml:"domainObject"`
								} `xml:"domainObjectList"`
								Index string `xml:"index"`
							} `xml:"loyaltyAccounts"`
							SkyClubPurchaseEligible string `xml:"skyClubPurchaseEligible"`
							AcctLocked              string `xml:"acctLocked"`
							PasswordCreated         string `xml:"passwordCreated"`
							LanguagePromptIndicator string `xml:"languagePromptIndicator"`
							CustomerId              string `xml:"customerId"`
							FirstNIN                string `xml:"firstNIN"`
							LastNIN                 string `xml:"lastNIN"`
							GroupLeader             string `xml:"groupLeader"`
							EmerContactInfoDeclined string `xml:"emerContactInfoDeclined"`
							CheckedIn               string `xml:"checkedIn"`
							TravelInfoProvided      string `xml:"travelInfoProvided"`
							Ptc                     string `xml:"ptc"`
							FlightSeats             struct {
								Text             string `xml:",chardata"`
								DomainObjectList struct {
									Text         string `xml:",chardata"`
									DomainObject []struct {
										Text                          string `xml:",chardata"`
										Type                          string `xml:"type,attr"`
										SeqNum                        string `xml:"seqNum"`
										Status                        string `xml:"status"`
										Available                     string `xml:"available"`
										Blocked                       string `xml:"blocked"`
										ForHandicapped                string `xml:"forHandicapped"`
										ForAdultWithInfant            string `xml:"forAdultWithInfant"`
										LimitedRecline                string `xml:"limitedRecline"`
										NoMovieView                   string `xml:"noMovieView"`
										NoStowage                     string `xml:"noStowage"`
										NoWindow                      string `xml:"noWindow"`
										NoSeat                        string `xml:"noSeat"`
										Preferred                     string `xml:"preferred"`
										PreferredAisle                string `xml:"preferredAisle"`
										EconomyComfortFlag            string `xml:"economyComfortFlag"`
										Middle                        string `xml:"middle"`
										Infant                        string `xml:"infant"`
										NonRevenueStandby             string `xml:"nonRevenueStandby"`
										RevenueStandby                string `xml:"revenueStandby"`
										Through                       string `xml:"through"`
										UnaccompaniedMinor            string `xml:"unaccompaniedMinor"`
										Window                        string `xml:"window"`
										Reserved                      string `xml:"reserved"`
										Restricted                    string `xml:"restricted"`
										CurrentPassengerSeat          string `xml:"currentPassengerSeat"`
										CurrentPassengerCompanionSeat string `xml:"currentPassengerCompanionSeat"`
										Cabin                         string `xml:"cabin"`
										SeatNumber                    string `xml:"seatNumber"`
										Exit                          string `xml:"exit"`
										Choice                        string `xml:"choice"`
										IsPaidSeat                    string `xml:"isPaidSeat"`
										Ineligible                    string `xml:"ineligible"`
										Occupied                      string `xml:"occupied"`
										LegId                         string `xml:"legId"`
										SegmentId                     string `xml:"segmentId"`
										CheckedIn                     string `xml:"checkedIn"`
									} `xml:"domainObject"`
								} `xml:"domainObjectList"`
								Index string `xml:"index"`
							} `xml:"flightSeats"`
							Selectee                  string `xml:"selectee"`
							PremiumPassengerIndicator string `xml:"premiumPassengerIndicator"`
							ExitRowEligible           string `xml:"exitRowEligible"`
							CheckedInOnALS            string `xml:"checkedInOnALS"`
							DoNotBoard                string `xml:"doNotBoard"`
							Tickets                   struct {
								Text             string `xml:",chardata"`
								DomainObjectList struct {
									Text         string `xml:",chardata"`
									DomainObject []struct {
										Text           string `xml:",chardata"`
										AttrType       string `xml:"type,attr"`
										Number         string `xml:"number"`
										Type           string `xml:"type"`
										IssueDate      string `xml:"issueDate"`
										ExpirationDate string `xml:"expirationDate"`
										PlaceOfIssue   string `xml:"placeOfIssue"`
										IssueCity      string `xml:"issueCity"`
										IssueLocation  string `xml:"issueLocation"`
										IssuingAgentId string `xml:"issuingAgentId"`
										Status         string `xml:"status"`
										TicketCoupons  struct {
											Text             string `xml:",chardata"`
											DomainObjectList struct {
												Text         string `xml:",chardata"`
												DomainObject []struct {
													Text         string `xml:",chardata"`
													Type         string `xml:"type,attr"`
													TicketNumber string `xml:"ticketNumber"`
													Status       string `xml:"status"`
													Flight       struct {
														Text   string `xml:",chardata"`
														Origin struct {
															Text      string `xml:",chardata"`
															Code      string `xml:"code"`
															Name      string `xml:"name"`
															CityServe string `xml:"cityServe"`
															Address   struct {
																Text     string `xml:",chardata"`
																Name     string `xml:"name"`
																Code     string `xml:"code"`
																Timezone string `xml:"timezone"`
																Country  struct {
																	Text   string `xml:",chardata"`
																	Code   string `xml:"code"`
																	Region struct {
																		Text string `xml:",chardata"`
																		Name string `xml:"name"`
																	} `xml:"region"`
																} `xml:"country"`
																Category string `xml:"category"`
															} `xml:"address"`
															ObserveDST string `xml:"observeDST"`
															Domestic   string `xml:"domestic"`
														} `xml:"origin"`
														Destination struct {
															Text      string `xml:",chardata"`
															Code      string `xml:"code"`
															Name      string `xml:"name"`
															CityServe string `xml:"cityServe"`
															Address   struct {
																Text     string `xml:",chardata"`
																Name     string `xml:"name"`
																Code     string `xml:"code"`
																Timezone string `xml:"timezone"`
																Country  struct {
																	Text   string `xml:",chardata"`
																	Code   string `xml:"code"`
																	Region struct {
																		Text string `xml:",chardata"`
																		Name string `xml:"name"`
																	} `xml:"region"`
																} `xml:"country"`
																Category string `xml:"category"`
															} `xml:"address"`
															ObserveDST string `xml:"observeDST"`
															Domestic   string `xml:"domestic"`
														} `xml:"destination"`
														DayChange               string `xml:"dayChange"`
														RedEyeFlag              string `xml:"redEyeFlag"`
														StandbyFlight           string `xml:"standbyFlight"`
														SeatMapEligible         string `xml:"seatMapEligible"`
														EquipmentChange         string `xml:"equipmentChange"`
														DlMarketedCodeshare     string `xml:"dlMarketedCodeshare"`
														DlConnectionCarrier     string `xml:"dlConnectionCarrier"`
														GroundHandled           string `xml:"groundHandled"`
														CleanedFlag             string `xml:"cleanedFlag"`
														MisconnectFlag          string `xml:"misconnectFlag"`
														PotentialMisconnectFlag string `xml:"potentialMisconnectFlag"`
														BasicEconomy            string `xml:"basicEconomy"`
														SeatAssigned            string `xml:"seatAssigned"`
														DisplayBrandGradient    string `xml:"displayBrandGradient"`
														DepartureDateTime       string `xml:"departureDateTime"`
														EticketEligible         string `xml:"eticketEligible"`
													} `xml:"flight"`
												} `xml:"domainObject"`
											} `xml:"domainObjectList"`
											Index string `xml:"index"`
										} `xml:"ticketCoupons"`
										Upsell        string `xml:"upsell"`
										Award         string `xml:"award"`
										MilesPlusCash string `xml:"milesPlusCash"`
									} `xml:"domainObject"`
								} `xml:"domainObjectList"`
								Index string `xml:"index"`
							} `xml:"tickets"`
							Ssrs struct {
								Text             string `xml:",chardata"`
								DomainObjectList struct {
									Text         string `xml:",chardata"`
									DomainObject []struct {
										Text            string `xml:",chardata"`
										Type            string `xml:"type,attr"`
										LineNumber      string `xml:"lineNumber"`
										Code            string `xml:"code"`
										GivenNameNumber string `xml:"givenNameNumber"`
										LastNameNumber  string `xml:"lastNameNumber"`
										AirlineCode     string `xml:"airlineCode"`
										SegmentNumber   string `xml:"segmentNumber"`
										StatusCode      string `xml:"statusCode"`
										Remarks         struct {
											Text   string `xml:",chardata"`
											Remark string `xml:"remark"`
										} `xml:"remarks"`
									} `xml:"domainObject"`
								} `xml:"domainObjectList"`
								Index string `xml:"index"`
							} `xml:"ssrs"`
							EnrollmentEligible string `xml:"enrollmentEligible"`
						} `xml:"domainObject"`
					} `xml:"domainObjectList"`
					Index string `xml:"index"`
				} `xml:"passengers"`
				MaxPartySizeExceeded     string `xml:"maxPartySizeExceeded"`
				LateCheckIn              string `xml:"lateCheckIn"`
				CheckInTooEarly          string `xml:"checkInTooEarly"`
				ProhibitedSSRFound       string `xml:"prohibitedSSRFound"`
				ReEntry                  string `xml:"reEntry"`
				EBoardingPassRequested   string `xml:"eBoardingPassRequested"`
				EBoardingPassEligible    string `xml:"eBoardingPassEligible"`
				SeatRequestCardRequested string `xml:"seatRequestCardRequested"`
				VholTicket               string `xml:"vholTicket"`
				PnrFlags                 struct {
					Text             string `xml:",chardata"`
					DomainObjectList struct {
						Text         string `xml:",chardata"`
						DomainObject []struct {
							Text     string `xml:",chardata"`
							Xsi      string `xml:"xsi,attr"`
							AttrType string `xml:"type,attr"`
							Name     string `xml:"name"`
							Value    string `xml:"value"`
							Type     string `xml:"type"`
						} `xml:"domainObject"`
					} `xml:"domainObjectList"`
					Index string `xml:"index"`
				} `xml:"pnrFlags"`
				Itineraries struct {
					Text             string `xml:",chardata"`
					DomainObjectList struct {
						Text         string `xml:",chardata"`
						DomainObject []struct {
							Text   string `xml:",chardata"`
							Xsi    string `xml:"xsi,attr"`
							Type   string `xml:"type,attr"`
							Origin struct {
								Text      string `xml:",chardata"`
								Code      string `xml:"code"`
								Name      string `xml:"name"`
								CityServe string `xml:"cityServe"`
								Address   struct {
									Text     string `xml:",chardata"`
									Name     string `xml:"name"`
									Code     string `xml:"code"`
									Timezone string `xml:"timezone"`
									Country  struct {
										Text   string `xml:",chardata"`
										Code   string `xml:"code"`
										Region struct {
											Text string `xml:",chardata"`
											Name string `xml:"name"`
										} `xml:"region"`
									} `xml:"country"`
									Category string `xml:"category"`
								} `xml:"address"`
								ObserveDST string `xml:"observeDST"`
								Domestic   string `xml:"domestic"`
							} `xml:"origin"`
							Destination struct {
								Text      string `xml:",chardata"`
								Code      string `xml:"code"`
								Name      string `xml:"name"`
								CityServe string `xml:"cityServe"`
								Address   struct {
									Text     string `xml:",chardata"`
									Name     string `xml:"name"`
									Code     string `xml:"code"`
									Timezone string `xml:"timezone"`
									Country  struct {
										Text   string `xml:",chardata"`
										Code   string `xml:"code"`
										Region struct {
											Text string `xml:",chardata"`
											Name string `xml:"name"`
										} `xml:"region"`
									} `xml:"country"`
									Category string `xml:"category"`
								} `xml:"address"`
								ObserveDST string `xml:"observeDST"`
								Domestic   string `xml:"domestic"`
							} `xml:"destination"`
							DepartureDateTime string `xml:"departureDateTime"`
							ArrivalDateTime   string `xml:"arrivalDateTime"`
							FlightTime        string `xml:"flightTime"`
							Flights           struct {
								Text             string `xml:",chardata"`
								DomainObjectList struct {
									Text         string `xml:",chardata"`
									DomainObject []struct {
										Text      string `xml:",chardata"`
										Type      string `xml:"type,attr"`
										LegId     string `xml:"legId"`
										SegmentId string `xml:"segmentId"`
										Origin    struct {
											Text      string `xml:",chardata"`
											Code      string `xml:"code"`
											Name      string `xml:"name"`
											CityServe string `xml:"cityServe"`
											Address   struct {
												Text     string `xml:",chardata"`
												Name     string `xml:"name"`
												Code     string `xml:"code"`
												Timezone string `xml:"timezone"`
												Country  struct {
													Text   string `xml:",chardata"`
													Code   string `xml:"code"`
													Region struct {
														Text string `xml:",chardata"`
														Name string `xml:"name"`
													} `xml:"region"`
												} `xml:"country"`
												Category string `xml:"category"`
											} `xml:"address"`
											ObserveDST string `xml:"observeDST"`
											Domestic   string `xml:"domestic"`
										} `xml:"origin"`
										Destination struct {
											Text      string `xml:",chardata"`
											Code      string `xml:"code"`
											Name      string `xml:"name"`
											CityServe string `xml:"cityServe"`
											Address   struct {
												Text     string `xml:",chardata"`
												Name     string `xml:"name"`
												Code     string `xml:"code"`
												Timezone string `xml:"timezone"`
												Country  struct {
													Text   string `xml:",chardata"`
													Code   string `xml:"code"`
													Region struct {
														Text string `xml:",chardata"`
														Name string `xml:"name"`
													} `xml:"region"`
												} `xml:"country"`
												Category string `xml:"category"`
											} `xml:"address"`
											ObserveDST string `xml:"observeDST"`
											Domestic   string `xml:"domestic"`
										} `xml:"destination"`
										ScheduledDepartureDateTime string `xml:"scheduledDepartureDateTime"`
										ScheduledArrivalDateTime   string `xml:"scheduledArrivalDateTime"`
										DayChange                  string `xml:"dayChange"`
										RedEyeFlag                 string `xml:"redEyeFlag"`
										LayOverTime                string `xml:"layOverTime"`
										FlightTime                 string `xml:"flightTime"`
										Distance                   string `xml:"distance"`
										Status                     string `xml:"status"`
										MarketingAirlineCode       string `xml:"marketingAirlineCode"`
										OperatingAirlineCode       string `xml:"operatingAirlineCode"`
										UpgradeStatus              string `xml:"upgradeStatus"`
										UpgradeStatusWCabin        string `xml:"upgradeStatusWCabin"`
										UpgradeCabinDescription    string `xml:"upgradeCabinDescription"`
										UpgradeWCabinDescription   string `xml:"upgradeWCabinDescription"`
										BrandAssociatedCabinId     string `xml:"brandAssociatedCabinId"`
										ClassesOfService           struct {
											Text             string `xml:",chardata"`
											DomainObjectList struct {
												Text         string `xml:",chardata"`
												DomainObject struct {
													Text        string `xml:",chardata"`
													Type        string `xml:"type,attr"`
													Code        string `xml:"code"`
													Description string `xml:"description"`
													AirlineCode string `xml:"airlineCode"`
													BrandId     string `xml:"brandId"`
												} `xml:"domainObject"`
											} `xml:"domainObjectList"`
											Index string `xml:"index"`
										} `xml:"classesOfService"`
										Meals struct {
											Text             string `xml:",chardata"`
											DomainObjectList struct {
												Text         string `xml:",chardata"`
												DomainObject []struct {
													Text        string `xml:",chardata"`
													Type        string `xml:"type,attr"`
													SeqNum      string `xml:"seqNum"`
													Description string `xml:"description"`
													ItemCode    string `xml:"itemCode"`
												} `xml:"domainObject"`
											} `xml:"domainObjectList"`
											Index string `xml:"index"`
										} `xml:"meals"`
										Amenities struct {
											Text             string `xml:",chardata"`
											DomainObjectList struct {
												Text         string `xml:",chardata"`
												DomainObject []struct {
													Text          string `xml:",chardata"`
													Type          string `xml:"type,attr"`
													SeqNum        string `xml:"seqNum"`
													Description   string `xml:"description"`
													ImageURL      string `xml:"imageURL"`
													IconCode      string `xml:"iconCode"`
													ItemCode      string `xml:"itemCode"`
													NewItemCode   string `xml:"newItemCode"`
													Tagline       string `xml:"tagline"`
													Disclaimer    string `xml:"disclaimer"`
													LearnMoreLink string `xml:"learnMoreLink"`
												} `xml:"domainObject"`
											} `xml:"domainObjectList"`
											Index string `xml:"index"`
										} `xml:"amenities"`
										StandbyFlight   string `xml:"standbyFlight"`
										SeatMapEligible string `xml:"seatMapEligible"`
										Equipment       struct {
											Text          string `xml:",chardata"`
											DeltaCode     string `xml:"deltaCode"`
											IndustryCode  string `xml:"industryCode"`
											Description   string `xml:"description"`
											NumPassengers string `xml:"numPassengers"`
										} `xml:"equipment"`
										EquipmentChange         string `xml:"equipmentChange"`
										CurrentActionCode       string `xml:"currentActionCode"`
										PreviousActionCode      string `xml:"previousActionCode"`
										DlMarketedCodeshare     string `xml:"dlMarketedCodeshare"`
										DlConnectionCarrier     string `xml:"dlConnectionCarrier"`
										GroundHandled           string `xml:"groundHandled"`
										IropType                string `xml:"iropType"`
										CleanedFlag             string `xml:"cleanedFlag"`
										MisconnectFlag          string `xml:"misconnectFlag"`
										PotentialMisconnectFlag string `xml:"potentialMisconnectFlag"`
										BasicEconomy            string `xml:"basicEconomy"`
										SeatAssigned            string `xml:"seatAssigned"`
										BrandGradientColorStart string `xml:"brandGradientColorStart"`
										BrandGradientColorEnd   string `xml:"brandGradientColorEnd"`
										BrandGradientAngle      string `xml:"brandGradientAngle"`
										DisplayBrandGradient    string `xml:"displayBrandGradient"`
										FlightNo                string `xml:"flightNo"`
										Airline                 struct {
											Text string `xml:",chardata"`
											Name string `xml:"name"`
											Code string `xml:"code"`
										} `xml:"airline"`
										DepartureDateTime string `xml:"departureDateTime"`
										ArrivalDateTime   string `xml:"arrivalDateTime"`
										EticketEligible   string `xml:"eticketEligible"`
										Legs              struct {
											Text             string `xml:",chardata"`
											DomainObjectList string `xml:"domainObjectList"`
											Index            string `xml:"index"`
										} `xml:"legs"`
										PreSelectMealService struct {
											Text             string `xml:",chardata"`
											EligibleFlag     string `xml:"eligibleFlag"`
											HoursToDeparture string `xml:"hoursToDeparture"`
										} `xml:"preSelectMealService"`
										AccumulatedFlightTime string `xml:"accumulatedFlightTime"`
										SchedChangeType       string `xml:"schedChangeType"`
									} `xml:"domainObject"`
								} `xml:"domainObjectList"`
								Index string `xml:"index"`
							} `xml:"flights"`
							SchedChgFlag        string `xml:"schedChgFlag"`
							CheckinEligible     string `xml:"checkinEligible"`
							WithinCheckinWindow string `xml:"withinCheckinWindow"`
							CleanedFlag         string `xml:"cleanedFlag"`
							MisconnectFlag      string `xml:"misconnectFlag"`
							StandByFlights      struct {
								Text             string `xml:",chardata"`
								DomainObjectList string `xml:"domainObjectList"`
								Index            string `xml:"index"`
							} `xml:"standByFlights"`
						} `xml:"domainObject"`
					} `xml:"domainObjectList"`
					Index string `xml:"index"`
				} `xml:"itineraries"`
				BookingDate string `xml:"bookingDate"`
				TotalFare   struct {
					Text             string `xml:",chardata"`
					BaseCurrencyCode string `xml:"baseCurrencyCode"`
					BaseFare         string `xml:"baseFare"`
					TotalTax         string `xml:"totalTax"`
					TotalFare        string `xml:"totalFare"`
					NumberOfDecimals string `xml:"numberOfDecimals"`
					TaxBreakDownList struct {
						Text         string `xml:",chardata"`
						FareFaxTable []struct {
							Text              string `xml:",chardata"`
							TaxType           string `xml:"taxType"`
							Amount            string `xml:"amount"`
							Currency          string `xml:"currency"`
							CarrierImposedFee string `xml:"carrierImposedFee"`
						} `xml:"fareFaxTable"`
					} `xml:"taxBreakDownList"`
					TotalCurrencyCode string `xml:"totalCurrencyCode"`
					FareBasisCode     string `xml:"fareBasisCode"`
				} `xml:"totalFare"`
				UpsellOffers struct {
					Text        string `xml:",chardata"`
					UpsellOffer []struct {
						Text              string `xml:",chardata"`
						OfferItemNumber   string `xml:"offerItemNumber"`
						UpsellType        string `xml:"upsellType"`
						UpsellDescription string `xml:"upsellDescription"`
						OfferScope        string `xml:"offerScope"`
						FarePerPax        struct {
							Text                string `xml:",chardata"`
							BaseCurrencyCode    string `xml:"baseCurrencyCode"`
							BaseFare            string `xml:"baseFare"`
							TotalFare           string `xml:"totalFare"`
							NumberOfDecimals    string `xml:"numberOfDecimals"`
							BaseFareUSD         string `xml:"baseFareUSD"`
							SpecialFareRuleList struct {
								Text            string `xml:",chardata"`
								SpecialFareRule struct {
									Text              string `xml:",chardata"`
									BookingClass      string `xml:"bookingClass"`
									SegmentNumber     string `xml:"segmentNumber"`
									DepartureDate     string `xml:"departureDate"`
									ArrivalCityCode   string `xml:"arrivalCityCode"`
									DepartureCityCode string `xml:"departureCityCode"`
									BrandId           string `xml:"brandId"`
								} `xml:"specialFareRule"`
							} `xml:"specialFareRuleList"`
							TaxBreakDownList struct {
								Text         string `xml:",chardata"`
								FareFaxTable struct {
									Text              string `xml:",chardata"`
									TaxType           string `xml:"taxType"`
									Amount            string `xml:"amount"`
									Currency          string `xml:"currency"`
									CarrierImposedFee string `xml:"carrierImposedFee"`
								} `xml:"fareFaxTable"`
							} `xml:"taxBreakDownList"`
							TotalCurrencyCode  string `xml:"totalCurrencyCode"`
							LoyaltyUpsellPrice struct {
								Text                string `xml:",chardata"`
								LoyaltyProgramType  string `xml:"loyaltyProgramType"`
								LoyaltyProgramName  string `xml:"loyaltyProgramName"`
								LoyaltyProgramCode  string `xml:"loyaltyProgramCode"`
								LoyaltyProgramPoint string `xml:"loyaltyProgramPoint"`
							} `xml:"loyaltyUpsellPrice"`
							AssociationIds struct {
								Text  string `xml:",chardata"`
								Entry []struct {
									Text  string `xml:",chardata"`
									Key   string `xml:"key"`
									Value string `xml:"value"`
								} `xml:"entry"`
							} `xml:"associationIds"`
						} `xml:"farePerPax"`
						FareClass  string `xml:"fareClass"`
						BrandId    string `xml:"brandId"`
						Eligible   string `xml:"eligible"`
						Passengers struct {
							Text             string `xml:",chardata"`
							DomainObjectList struct {
								Text         string `xml:",chardata"`
								DomainObject struct {
									Text                string `xml:",chardata"`
									Xsi                 string `xml:"xsi,attr"`
									Type                string `xml:"type,attr"`
									SecurityQAStoredInd string `xml:"securityQAStoredInd"`
									SqaLocked           string `xml:"sqaLocked"`
									AddrLocked          string `xml:"addrLocked"`
									Name                struct {
										Text      string `xml:",chardata"`
										FirstName string `xml:"firstName"`
										LastName  string `xml:"lastName"`
									} `xml:"name"`
									SkyClubPurchaseEligible   string `xml:"skyClubPurchaseEligible"`
									AcctLocked                string `xml:"acctLocked"`
									PasswordCreated           string `xml:"passwordCreated"`
									LanguagePromptIndicator   string `xml:"languagePromptIndicator"`
									FirstNIN                  string `xml:"firstNIN"`
									LastNIN                   string `xml:"lastNIN"`
									GroupLeader               string `xml:"groupLeader"`
									EmerContactInfoDeclined   string `xml:"emerContactInfoDeclined"`
									CheckedIn                 string `xml:"checkedIn"`
									TravelInfoProvided        string `xml:"travelInfoProvided"`
									Selectee                  string `xml:"selectee"`
									PremiumPassengerIndicator string `xml:"premiumPassengerIndicator"`
									ExitRowEligible           string `xml:"exitRowEligible"`
									CheckedInOnALS            string `xml:"checkedInOnALS"`
									DoNotBoard                string `xml:"doNotBoard"`
									EnrollmentEligible        string `xml:"enrollmentEligible"`
								} `xml:"domainObject"`
							} `xml:"domainObjectList"`
							Index string `xml:"index"`
						} `xml:"passengers"`
						GroupId string `xml:"groupId"`
					} `xml:"upsellOffer"`
				} `xml:"upsellOffers"`
				UpsellFaresPerPax struct {
					Text       string `xml:",chardata"`
					UpsellFare struct {
						Text               string `xml:",chardata"`
						ClassOfService     string `xml:"classOfService"`
						ClassOfServiceCode string `xml:"classOfServiceCode"`
						BaseCurrencyCode   string `xml:"baseCurrencyCode"`
						TotalFare          string `xml:"totalFare"`
						TotalFareUSD       string `xml:"totalFareUSD"`
						NumberOfDecimals   string `xml:"numberOfDecimals"`
						PartialUpsell      string `xml:"partialUpsell"`
						LoyaltyUpsellPrice struct {
							Text                string `xml:",chardata"`
							LoyaltyProgramType  string `xml:"loyaltyProgramType"`
							LoyaltyProgramName  string `xml:"loyaltyProgramName"`
							LoyaltyProgramCode  string `xml:"loyaltyProgramCode"`
							LoyaltyProgramPoint string `xml:"loyaltyProgramPoint"`
						} `xml:"loyaltyUpsellPrice"`
						UpsoldRoutes struct {
							Text        string `xml:",chardata"`
							UpsoldRoute struct {
								Text        string `xml:",chardata"`
								Origin      string `xml:"origin"`
								Destination string `xml:"destination"`
								SegmentId   string `xml:"segmentId"`
							} `xml:"upsoldRoute"`
						} `xml:"upsoldRoutes"`
						UpsellIndex      string `xml:"upsellIndex"`
						OfferItemNumbers struct {
							Text            string   `xml:",chardata"`
							OfferItemNumber []string `xml:"offerItemNumber"`
						} `xml:"offerItemNumbers"`
						HasOfferForAllSegments string `xml:"hasOfferForAllSegments"`
						HasSameOfferPerTrip    string `xml:"hasSameOfferPerTrip"`
						PromoImage             string `xml:"promoImage"`
						IntroText              string `xml:"introText"`
						FareClassDescription   string `xml:"fareClassDescription"`
						PromoDescription       string `xml:"promoDescription"`
						BrandIds               struct {
							Text    string `xml:",chardata"`
							BrandId string `xml:"brandId"`
						} `xml:"brandIds"`
						UpsellPaymentOption  string `xml:"upsellPaymentOption"`
						UpsellPaymentOption2 string `xml:"upsellPaymentOption2"`
						AssociationIds       struct {
							Text  string `xml:",chardata"`
							Entry []struct {
								Text  string `xml:",chardata"`
								Key   string `xml:"key"`
								Value string `xml:"value"`
							} `xml:"entry"`
						} `xml:"associationIds"`
						GradientStartColor string `xml:"gradientStartColor"`
						GradientEndColor   string `xml:"gradientEndColor"`
						GradientAngle      string `xml:"gradientAngle"`
						Amenities          struct {
							Text             string `xml:",chardata"`
							DomainObjectList struct {
								Text         string `xml:",chardata"`
								DomainObject []struct {
									Text        string `xml:",chardata"`
									Xsi         string `xml:"xsi,attr"`
									Type        string `xml:"type,attr"`
									ID          string `xml:"id"`
									SeqNum      string `xml:"seqNum"`
									Description string `xml:"description"`
									ImageURL    string `xml:"imageURL"`
								} `xml:"domainObject"`
							} `xml:"domainObjectList"`
							Index string `xml:"index"`
						} `xml:"amenities"`
					} `xml:"upsellFare"`
				} `xml:"upsellFaresPerPax"`
				PhoneNumber       string `xml:"phoneNumber"`
				BusinessIndicator string `xml:"businessIndicator"`
			} `xml:"pnr"`
			Vacation string `xml:"vacation"`
		} `xml:"Journey"`
	} `xml:"tripsResponse"`
}
