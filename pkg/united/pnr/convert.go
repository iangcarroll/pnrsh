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
			Cabin:                  segment.BookingClass.Cabin.Name,
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
