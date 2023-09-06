package pnr

import "strconv"

func convertResponse(res GetPNRResponse) (pnr PNR) {
	convertTickets(res, &pnr)

	// Look for DL flights.
	for _, connection := range res.Itinerary.Connections {
		for _, segment := range connection.Segments {
			if segment.MarketingFlight.Carrier.Code == "DL" || segment.MarketingFlight.OperatingFlight.Carrier.Code == "DL" {
				pnr.HasDLFlights = true
				break
			}
		}
	}

	pnr.SMCalcLink = generateSmcalcLink(&pnr)
	pnr.QMCalcLink = generateQMCalcLink(&pnr)

	return pnr
}

func convertTickets(res GetPNRResponse, pnr *PNR) {

	for _, pax := range res.Passengers {

		var ffpMembership FrequentTravelerMembership
		for _, membership := range pax.Memberships {
			if membership.FrequentTravelerMembership.Number != "" {
				ffpMembership = membership.FrequentTravelerMembership
				break
			}
		}

		for _, ticket := range pax.Tickets {
			numCoupons := 0

			var flights []Flight
			for _, booklet := range ticket.TicketBooklets {
				numCoupons += len(booklet.TicketCoupons) + len(booklet.InvalidTicketCoupons)

				for _, coupon := range booklet.TicketCoupons {
					for _, connection := range res.Itinerary.Connections {
						for _, flight := range connection.Segments {
							if coupon.SegmentID != flight.ID {
								continue
							}

							classOfService := ""
							skyPriority := false
							for _, detail := range res.Passengers[0].PassengerFlightDetails {
								if detail.SegmentID != flight.ID {
									continue
								}

								classOfService = detail.Cabin.Name
								skyPriority = detail.SkyPriority
							}

							flights = append(flights, Flight{
								OriginAirportCode:      flight.DepartFrom.IATACode,
								DestinationAirportCode: flight.ArriveOn.IATACode,
								MarketingAirlineCode:   flight.MarketingFlight.Carrier.Code,
								MarketingAirlineName:   flight.MarketingFlight.Carrier.Name,
								OperatingAirlineCode:   flight.MarketingFlight.OperatingFlight.Carrier.Code,
								OperatingAirlineName:   flight.MarketingFlight.OperatingFlight.Carrier.Name,
								FlightNumber:           flight.MarketingFlight.Number,
								ClassOfService:         classOfService,
								ScheduledDeparture:     flight.MarketingFlight.OperatingFlight.LocalScheduledDeparture,
								ScheduledArrival:       flight.MarketingFlight.OperatingFlight.LocalScheduledArrival,
								Status:                 flight.MarketingFlight.Status.Code,
								FareClass:              flight.MarketingFlight.SellingClass.Code,
								AircraftType:           flight.MarketingFlight.OperatingFlight.Equipment.Name + " (" + flight.MarketingFlight.OperatingFlight.Equipment.Code + ")",
								SkyPriority:            skyPriority,
							})

							break
						}
					}
				}
			}

			var carrierImposedSurcharge float64 = 0
			carrierImposedSurchargeCode := ""

			for _, charge := range ticket.AirTransportationCharges.Charges {
				if charge.CostType.Code == "AIRLINE_COSTS" {
					if carrierImposedSurchargeCode == "" || carrierImposedSurchargeCode == charge.Amount.Currency {
						carrierImposedSurcharge += charge.Amount.Amount
						carrierImposedSurchargeCode = charge.Amount.Currency
					} else {
						carrierImposedSurcharge = 0
						carrierImposedSurchargeCode = "Mixed Currencies"
						break
					}
				}
			}

			var taxes []TaxRow
			var totalTax float64 = 0
			totalTaxCurrencyCode := ""

			for _, tax := range ticket.Taxes {
				taxes = append(taxes, TaxRow{
					TaxType:           tax.Code + " - " + tax.Name,
					Amount:            strconv.FormatFloat(tax.Amount.Amount, 'f', 2, 64),
					Currency:          tax.Amount.Currency,
					CarrierImposedFee: false,
				})
				if totalTaxCurrencyCode == "" || totalTaxCurrencyCode == tax.Amount.Currency {
					totalTax += tax.Amount.Amount
					totalTaxCurrencyCode = tax.Amount.Currency
				} else {
					totalTax = 0
					totalTaxCurrencyCode = "Mixed Currencies"
					break
				}
			}

			pnr.Tickets = append(pnr.Tickets, Ticket{
				Number:                 ticket.Number,
				IssueDate:              ticket.TicketingDate,
				PassengerName:          pax.FirstName + " " + pax.LastName,
				NumCoupons:             numCoupons,
				ValidatedAgainstCoupon: false,
				Passenger: Passenger{
					Name:                 pax.FirstName + " " + pax.LastName,
					FrequentFlyerNumber:  ffpMembership.Number,
					FrequentFlyerProgram: ffpMembership.LoyaltyProgram.Name,
				},
				Flights: flights,
				Fare: Fare{
					BaseFare:                    strconv.FormatFloat(ticket.BaseFare.Amount, 'f', 2, 64),
					BaseCurrencyCode:            ticket.BaseFare.Currency,
					CarrierImposedSurcharge:     strconv.FormatFloat(float64(carrierImposedSurcharge), 'f', 2, 64),
					CarrierImposedSurchargeCode: carrierImposedSurchargeCode,
					TotalTax:                    strconv.FormatFloat(float64(totalTax), 'f', 2, 64),
					TotalTaxCurrencyCode:        totalTaxCurrencyCode,
					TotalFare:                   strconv.FormatFloat(ticket.TotalFare.Amount, 'f', 2, 64),
					TotalCurrencyCode:           ticket.TotalFare.Currency,
				},
			})
		}
	}
}
