package pnr

import (
	"fmt"
	"net/url"
	"strconv"
)

func asFloat(n string) float64 {
	s, _ := strconv.ParseFloat(n, 64)
	return s
}

func asString(n float64) string {
	return fmt.Sprintf("%.2f", n)
}

func estimateMQD(pnr *PNR) string {
	// We can't calculate MQD very easily for non-USD.
	if pnr.Fare.BaseCurrencyCode != "USD" || pnr.Fare.TotalCurrencyCode != "USD" {
		return "unknown currency"
	}

	total := float64(0)
	total += asFloat(pnr.Fare.BaseFare)

	for _, row := range pnr.Fare.TaxRows {
		if row.CarrierImposedFee && row.Currency == "USD" {
			total += asFloat(row.Amount)
		} else if row.Currency != "USD" {
			return "unknown currency"
		}
	}

	return asString(total)
}

// Calculate the highest status in the reservation and pass to SMCalc.
func generateSmcalcStatus(pnr *PNR) (highestStatus string) {
	for _, pax := range pnr.Passengers {
		switch pax.Status {
		case "DM":
			highestStatus = "diamond"
		case "PM":
			if highestStatus != "diamond" {
				highestStatus = "platinum"
			}
		case "GM":
			if highestStatus != "diamond" && highestStatus != "platinum" {
				highestStatus = "gold"
			}
		case "SM":
			if highestStatus == "" {
				highestStatus = "silver"
			}
		}
	}

	if highestStatus == "" {
		highestStatus = "peon"
	}

	return highestStatus
}

func generateSmcalcRoute(pnr *PNR) (out string) {
	var fallbackClass string

	// We only know the current PNR fare class, which is not useful when it's
	// a longer two-character code like SU -- which just implies a C+ upgrade.
	// As a result, we use the fare basis code as the fallback when we hit
	// a long fare class in the PNR.
	if len(pnr.Fare.FareBasisCode) > 0 {
		fallbackClass = string(pnr.Fare.FareBasisCode[0])
	}

	for idx, flight := range pnr.Flights {
		class := flight.ClassOfService

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

	if pnr.Fare.TotalFare != "" && pnr.Fare.TotalCurrencyCode == "USD" {
		out += " $" + pnr.Fare.EstimatedMQD
	}

	return out
}

func generateSmcalcLink(pnr *PNR) string {
	route := generateSmcalcRoute(pnr)
	return fmt.Sprintf("https://skymilescalculator.com/dist.php?route=%s&start_mqm=0&start_rdm=0&default_fare=T&default_carrier=DL&elite=%s", url.QueryEscape(route), generateSmcalcStatus(pnr))
}

func generateQMCalcRoute(pnr *PNR) (out string) {
	for idx, flight := range pnr.Flights {
		if idx == 0 {
			out += flight.OriginAirportCode
		}

		var fare string
		if len(flight.ClassOfService) == 0 {
			fare = flight.MarketingAirlineCode
		} else {
			fare = fmt.Sprintf("%s.%s", flight.MarketingAirlineCode, flight.ClassOfService)
		}

		if len(flight.OperatingAirlineCode) > 0 {
			fare += fmt.Sprintf("/%s", flight.OperatingAirlineCode)
		}

		out += fmt.Sprintf("-%s-%s", fare, flight.DestinationAirportCode)
	}

	if pnr.Fare.TotalFare != "" && pnr.Fare.TotalCurrencyCode == "USD" {
		out += " $" + pnr.Fare.EstimatedMQD
	}

	return out
}

func generateQMCalcLink(pnr *PNR) string {
	route := generateQMCalcRoute(pnr)
	queryString := "?q=" + url.QueryEscape(route)

	status := highestStatus(pnr)
	if status != "" {
		queryString += "&s=" + url.QueryEscape(status)
	}

	return fmt.Sprintf("https://www.qualifyingmiles.com/%s", queryString)
}

func highestStatus(pnr *PNR) string {
	tiers := []string{"360", "DM", "PM", "GM", "FO", "SM", "GM", ""}
	for _, tier := range tiers {
		for _, pax := range pnr.Passengers {
			if pax.Status == tier {
				return tier
			}
		}
	}

	return ""
}
