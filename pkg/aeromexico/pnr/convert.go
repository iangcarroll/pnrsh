package pnr

import "fmt"

func convertRemarks(res ManagePnrResponse, pnr *PNR) {
	for _, collection := range res.Collection {
		for _, remark := range collection.Remarks {
			pnr.Remarks = append(pnr.Remarks, Remark{
				FreeFormText: remark,
				RemarkType:   "RRMK",
			})
		}
	}
}

func convertFlights(res ManagePnrResponse, pnr *PNR) {
	for _, collection := range res.Collection {
		for _, flight := range collection.Legs.Collection {
			for _, segment := range flight.Segments.Collection {
				segment := segment.Segment
				pnr.Flights = append(pnr.Flights, Flight{
					OriginAirportCode:      segment.DepartureAirport,
					DestinationAirportCode: segment.ArrivalAirport,
					OperatingAirlineCode:   segment.OperatingCarrier,
					MarketingAirlineCode:   segment.MarketingCarrier,
					FlightNumber:           segment.MarketingFlightCode,
					CurrentActionCode:      segment.SegmentStatus,
					ClassOfService:         segment.BookingClass,
					Cabin:                  segment.Cabin,
					Status:                 flight.ManageStatus,
					ScheduledDeparture:     segment.DepartureDateTime,
					ScheduledArrival:       segment.ArrivalDateTime,
					FareBasis:              segment.FareBasisCode,
					FareName:               segment.FareAndCabinName,
				})
			}
		}
	}
}

func convertPassengers(res ManagePnrResponse, pnr *PNR) {
	for _, collection := range res.Collection {
		for _, pax := range collection.Cart.TravelerInfo.Collection {
			p := Passenger{
				Name:                   pax.DisplayName,
				OverbookingEligible:    pax.IsOverBookingEligible,
				SkyPriority:            pax.IsSkyPriority,
				FeeForSeatSelection:    true,
				FeeForAmPlusUpgrade:    true,
				FeeForPreferredUpgrade: true,
			}

			if len(pax.Loyalty) > 0 {
				p.Status = pax.Loyalty[0].Number + " " + pax.Loyalty[0].TierTag
			}

			for _, benefit := range pax.Benefit.Collection {
				p.BenefitCodes += benefit.Code + " "
				p.FeeForSeatSelection = p.FeeForSeatSelection && benefit.SeatBenefits.FeeRequired.SeatSelection
				p.FeeForPreferredUpgrade = p.FeeForPreferredUpgrade && benefit.SeatBenefits.FeeRequired.PrefferedUpgrade
				p.FeeForAmPlusUpgrade = p.FeeForAmPlusUpgrade && benefit.SeatBenefits.FeeRequired.AmPlusUgrade
			}

			pnr.Passengers = append(pnr.Passengers, p)
		}
	}
}
func convertTickets(res ManagePnrResponse, pnr *PNR) {
	for _, collection := range res.Collection {
		for _, pax := range collection.Cart.TravelerInfo.Collection {
			for _, ticket := range pax.TicketNumbers {
				pnr.Tickets = append(pnr.Tickets, Ticket{
					PassengerName:         pax.DisplayName,
					Number:                ticket.Number,
					CouponNumber:          ticket.Coupon,
					Status:                ticket.Status,
					PreviousStatus:        ticket.PreviousStatus,
					RelatedDocumentNumber: ticket.RelatedDocumentNumber,
					OriginDestination:     ticket.StartLocation + "/" + ticket.EndLocation,
					TotalCost:             fmt.Sprintf("%.2f", pax.AmountTicket.Total),
				})
			}
		}
	}
}

func convertEarnings(_ ManagePnrResponse, pnr *PNR) {
	pnr.SMCalcLink = generateSmcalcLink(pnr)
}
