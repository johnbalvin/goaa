package flights

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

func Get(locale, originAirport, destinationAirport, departDate, returnDate string, passengers int, proxyURL *url.URL) (Output, error) {
	input := Input{
		Metadata: Metadata{
			TripType: "RoundTrip",
			Udo: Udo{
				SearchMethod: "Lowest",
			},
		},
		Passengers: []Passenger{
			{Type: "adult", Count: passengers},
		},
		RequestHeader: RequestHeader{
			ClientId: "AAcom",
		},
		Slices: []Slice{
			{AllCarriers: true, DepartureDate: departDate, Destination: destinationAirport, Origin: originAirport},
			{AllCarriers: true, DepartureDate: returnDate, Destination: originAirport, Origin: destinationAirport},
		},
		TripOptions: TripOptions{FareType: "Lowest", Locale: locale, SearchType: "Revenue"},
	}
	var buffer bytes.Buffer
	if err := json.NewEncoder(&buffer).Encode(input); err != nil {
		return Output{}, err
	}
	req, err := http.NewRequest("POST", ep2, &buffer)
	if err != nil {
		return Output{}, err
	}
	req.Header.Add("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Add("Accept-Language", "en")
	req.Header.Add("Cache-Control", "no-cache")
	req.Header.Add("Pragma", "no-cache")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Sec-Ch-Ua", `"Not_A Brand";v="8", "Chromium";v="120", "Google Chrome";v="120"`)
	req.Header.Add("Sec-Ch-Ua-Mobile", "?0")
	req.Header.Add("Sec-Ch-Ua-Platform", `"Windows"`)
	req.Header.Add("Sec-Fetch-Dest", "document")
	req.Header.Add("Sec-Fetch-Mode", "navigate")
	req.Header.Add("Sec-Fetch-Site", "none")
	req.Header.Add("Sec-Fetch-User", "?1")
	req.Header.Add("Upgrade-Insecure-Requests", "1")
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36")
	transport := &http.Transport{
		MaxIdleConnsPerHost: 30,
		DisableKeepAlives:   true,
	}
	if proxyURL != nil {
		transport.Proxy = http.ProxyURL(proxyURL)
	}
	client := &http.Client{
		Timeout: time.Minute,
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Transport: transport,
	}
	resp, err := client.Do(req)
	if err != nil {
		return Output{}, err
	}
	if resp.StatusCode != 200 {
		errData := fmt.Errorf("status: %d headers: %+v", resp.StatusCode, resp.Header)
		return Output{}, errData
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return Output{}, err
	}
	var data Output
	if err := json.Unmarshal(body, &data); err != nil {
		return Output{}, err
	}
	return data, nil
}
