package goaa

import (
	"goaa/airports"
	"goaa/flights"
	"net/url"
)

func Airports(searchText string, proxyURL *url.URL) ([]AirportData, error) {
	datas, err := airports.Get(searchText, proxyURL)
	if err != nil {
		return nil, err
	}
	var allDatas []AirportData
	for _, data := range datas {
		allDatas = append(allDatas, AirportData(data))
	}
	return allDatas, err
}

func Flights(locale, originAirport, destinationAirport, departDate, returnDate string, passengers int, proxyURL *url.URL) (FlightsData, error) {
	data, err := flights.Get(locale, originAirport, destinationAirport, departDate, returnDate, passengers, proxyURL)
	if err != nil {
		return FlightsData{}, err
	}
	return FlightsData(data), nil
}
