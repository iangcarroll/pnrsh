package pnr

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

func generateSmcalcRoute(pnr *PNR) (out string) {
	var fallbackClass string

	for _, ticket := range pnr.Tickets {

		// We only know the current PNR fare class, which is not useful when it's
		// a longer two-character code like SU -- which just implies a C+ upgrade.
		// As a result, we use the fare basis code as the fallback when we hit
		// a long fare class in the PNR.
		if len(ticket.Fare.FareBasisCode) > 0 {
			fallbackClass = string(ticket.Fare.FareBasisCode[0])
		}

		for idx, flight := range ticket.Flights {
			class := flight.FareClass

			if len(class) != 1 && len(fallbackClass) == 1 {
				class = fallbackClass
			} else if len(class) != 1 {
				class = "V"
			}

			if idx > 0 {
				out += fmt.Sprintf("-%s.%s-%s", flight.MarketingAirlineCode, class, flight.DestinationAirportCode)
			} else {
				out += fmt.Sprintf("%s-%s.%s-%s", flight.OriginAirportCode, flight.MarketingAirlineCode, class, flight.DestinationAirportCode)
			}
		}

		if mileageEarningFare := computeMileageEarningFare(ticket); mileageEarningFare > 0 {
			out += " $" + strconv.FormatFloat(mileageEarningFare, 'f', 2, 64)
		}

		out += "\n"
	}

	return strings.TrimSpace(out)
}

func generateSmcalcLink(pnr *PNR) string {
	route := generateSmcalcRoute(pnr)
	return fmt.Sprintf("https://skymilescalculator.com/dist.php?route=%s&start_mqm=0&start_rdm=0&default_fare=T&default_carrier=DL&elite=%s", url.QueryEscape(route), "peon")
}

func generateQMCalcRoute(pnr *PNR) (out string) {
	for _, ticket := range pnr.Tickets {
		for idx, flight := range ticket.Flights {
			if idx == 0 {
				out += flight.OriginAirportCode
			}

			var fare string
			if len(flight.FareClass) == 0 {
				fare = flight.MarketingAirlineCode
			} else {
				fare = fmt.Sprintf("%s.%s", flight.MarketingAirlineCode, flight.FareClass)
			}

			if len(flight.OperatingAirlineCode) > 0 {
				fare += fmt.Sprintf("/%s", flight.OperatingAirlineCode)
			}

			out += fmt.Sprintf("-%s-%s", fare, flight.DestinationAirportCode)
		}

		if mileageEarningFare := computeMileageEarningFare(ticket); mileageEarningFare > 0 {
			out += " $" + strconv.FormatFloat(mileageEarningFare, 'f', 2, 64)
		}

		out += "\n"
	}

	return strings.TrimSpace(out)
}

func generateQMCalcLink(pnr *PNR) string {
	route := generateQMCalcRoute(pnr)
	queryString := "?q=" + url.QueryEscape(route)

	return fmt.Sprintf("https://www.qualifyingmiles.com/%s", queryString)
}

func computeMileageEarningFare(ticket Ticket) float64 {
	if ticket.Fare.BaseCurrencyCode != "USD" {
		return 0
	}

	fare, err := strconv.ParseFloat(ticket.Fare.BaseFare, 64)
	if err != nil {
		return 0
	}

	if ticket.Fare.CarrierImposedSurcharge != "" {
		if ticket.Fare.CarrierImposedSurchargeCode != "USD" {
			return fare
		}

		surcharges, err := strconv.ParseFloat(ticket.Fare.CarrierImposedSurcharge, 64)
		if err != nil {
			return 0
		}

		fare += surcharges
	}

	return fare
}
