package pnr

func convertFlights(res GetPNRResponse, pnr *PNR) {
	for _, segment := range res.Detail.FlightSegments {
		pnr.Flights = append(pnr.Flights, Flight{
			OriginAirportCode:      segment.FlightSegment.DepartureAirport.IATACode,
			DestinationAirportCode: segment.FlightSegment.ArrivalAirport.IATACode,
			OperatingAirlineCode:   segment.FlightSegment.OperatingAirlineCode,
			FlightNumber:           segment.FlightSegment.FlightNumber,
			CurrentActionCode:      segment.FlightSegment.FlightSegmentType,
			ClassOfService:         segment.BookingClass.Cabin.Name,
			ScheduledDeparture:     segment.FlightSegment.DepartureDateTime,
			ScheduledArrival:       segment.FlightSegment.ArrivalDateTime,
			Cabin:                  segment.BookingClass.Code,
			// TODO: Probably unsafe to assert there is an entry here.
			MarketingAirlineCode: segment.FlightSegment.MarketedFlightSegment[0].MarketingAirlineCode,
		})
	}
}

func convertPassengers(res GetPNRResponse, pnr *PNR) {
	for _, pax := range res.Detail.Travelers {
		pnr.Passengers = append(pnr.Passengers, Passenger{
			Name:   pax.Person.GivenName + " " + pax.Person.Surname,
			Status: pax.LoyaltyProgramProfile.LoyaltyProgramMemberTierLevel,
		})
	}
}

func convertRemarks(res GetPNRResponse, pnr *PNR) {
	for _, rmk := range res.Detail.Remarks {
		pnr.Remarks = append(pnr.Remarks, Remark{
			Description:     rmk.Description,
			DisplaySequence: rmk.DisplaySequence,
		})
	}
}

func convertTickets(res GetPNRResponse, pnr *PNR) {
	for _, pax := range res.Detail.Travelers {
		for _, tkt := range pax.Tickets {
			formattedTicket := Ticket{
				DocumentID:         tkt.DocumentID,
				IssueDate:          tkt.IssueDate,
				TicketValidityDate: tkt.TicketValidityDate,
				Coupons:            []Coupon{},
			}

			for _, coupon := range tkt.FlightCoupons {
				formattedTicket.Coupons = append(formattedTicket.Coupons, Coupon{
					Status:               coupon.Status.Code,
					DepartureAirport:     coupon.FlightSegment.DepartureAirport.IATACode,
					ArrivalAirport:       coupon.FlightSegment.ArrivalAirport.IATACode,
					FlightNumber:         coupon.FlightSegment.FlightNumber,
					OperatingAirlineCode: coupon.FlightSegment.OperatingAirlineCode,
				})
			}

			pnr.Tickets = append(pnr.Tickets, formattedTicket)
		}
	}
}
