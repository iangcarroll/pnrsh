package pnr

func convertFlights(res GetPNRResponse, pnr *PNR) {
	for _, segment := range res.Pnr.Segments {
		pnr.Flights = append(pnr.Flights, Flight{
			FareBasis:              segment.FareBasisCode,
			OriginAirportCode:      segment.Departure.Code,
			DestinationAirportCode: segment.Arrival.Code,
			MarketingAirlineCode:   segment.MarketingCarrier.Code,
			FlightNumber:           segment.MarketingCarrier.FlightNumber,
			OperatingAirlineCode:   segment.OperationoperatingCarrier.Code,
			CurrentActionCode:      segment.ActionCode,
			ClassOfService:         segment.ClassOfService,
			ScheduledDeparture:     segment.ScheduledDepartureDateTime,
			ScheduledArrival:       segment.ScheduledArrivalDateTime,
			PreviousActionCode:     segment.UpgradeVisibility.PreviousSegmentActionCode,
			Cabin:                  segment.CabinType,
		})
	}
}

func convertPassengers(res GetPNRResponse, pnr *PNR) {
	for _, pax := range res.Pnr.Passengers {
		pnr.Passengers = append(pnr.Passengers, Passenger{
			Name:                pax.PassengerName.First + " " + pax.PassengerName.Last,
			Status:              pax.MileagePlus.CurrentEliteLevelDescription,
			IsCEO:               pax.MileagePlus.IsCEO,
			IsInfiniteElite:     pax.MileagePlus.IsInfiniteElite,
			IsLifetimeCompanion: pax.MileagePlus.IsLifetimeCompanion,
			IsUnitedClubMember:  pax.MileagePlus.IsUnitedClubMember,
			IsPresidentialPlus:  pax.MileagePlus.IsPresidentialPlus,
			SharesPosition:      pax.SharesPosition,
		})
	}
}
