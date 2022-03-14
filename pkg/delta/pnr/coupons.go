package pnr

import "log"

func couponsMatchFlights(res RetrievePnrResponse, ticketNum string) bool {
	numTotalFlights := 0
	for _, itin := range res.TripsResponse.Journey.Pnr.Itineraries.DomainObjectList.DomainObject {
		numTotalFlights += len(itin.Flights.DomainObjectList.DomainObject)
	}

	log.Printf("%d total flights in itin...", numTotalFlights)

	for _, pax := range res.TripsResponse.Journey.Pnr.Passengers.DomainObjectList.DomainObject {
		for _, ticket := range pax.Tickets.DomainObjectList.DomainObject {
			if ticket.Number != ticketNum {
				continue
			}

			numValidatedCoupons := 0

			for _, coupon := range ticket.TicketCoupons.DomainObjectList.DomainObject {
				for _, itin := range res.TripsResponse.Journey.Pnr.Itineraries.DomainObjectList.DomainObject {
					for _, flight := range itin.Flights.DomainObjectList.DomainObject {
						if coupon.Flight.Destination.Code == flight.Destination.Code && coupon.Flight.Origin.Code == flight.Origin.Code && coupon.Flight.DepartureDateTime == flight.DepartureDateTime {
							numValidatedCoupons += 1
						}
					}
				}
			}

			log.Printf("%d total validated coupons for this pax...", numValidatedCoupons)
			return numValidatedCoupons >= numTotalFlights
		}
	}

	return false
}
