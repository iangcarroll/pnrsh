package pnr

import (
	"fmt"
	"net/url"
	"strings"
)

func generateSmcalcStatus(pnr *PNR) (highestStatus string) {
	for _, pax := range pnr.Passengers {
		benefits := strings.ToLower(pax.BenefitCodes)

		if strings.Contains(benefits, "diamond") {
			highestStatus = "diamond"
		} else if strings.Contains(benefits, "platinum") && highestStatus != "diamond" {
			highestStatus = "platinum"
		} else if strings.Contains(benefits, "gold") && highestStatus != "diamond" && highestStatus != "platinum" {
			highestStatus = "gold"
		} else if strings.Contains(benefits, "silver") && highestStatus == "" {
			highestStatus = "silver"
		}
	}

	if highestStatus == "" {
		return "peon"
	}

	return highestStatus
}

func generateSmcalcRoute(pnr *PNR) (out string) {
	var fallbackClass string
	if len(pnr.Flights) > 0 && len(pnr.Flights[0].FareBasis) > 0 {
		fallbackClass = string(pnr.Flights[0].FareBasis[0])
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

	if len(pnr.Tickets) > 0 {
		out += " $" + pnr.Tickets[0].TotalCost
	}

	return out
}

func generateSmcalcLink(pnr *PNR) string {
	route := generateSmcalcRoute(pnr)
	return fmt.Sprintf("https://fly.qux.us/smcalc/dist.php?route=%s&start_mqm=0&start_rdm=0&default_fare=V&default_carrier=AM&elite=%s", url.QueryEscape(route), generateSmcalcStatus(pnr))
}
