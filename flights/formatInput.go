package flights

type Input struct {
	Metadata      Metadata      `json:"metadata"`
	Passengers    []Passenger   `json:"passengers"`
	RequestHeader RequestHeader `json:"requestHeader"`
	Slices        []Slice       `json:"slices"`
	TripOptions   TripOptions   `json:"tripOptions"`
	LoyaltyInfo   *interface{}  `json:"loyaltyInfo"`
	Version       string        `json:"version"`
	QueryParams   QueryParams   `json:"queryParams"`
}

type Metadata struct {
	SelectedProducts []interface{} `json:"selectedProducts"`
	TripType         string        `json:"tripType"`
	Udo              Udo           `json:"udo"`
}

type Udo struct {
	SearchMethod string `json:"search_method"`
}

type Passenger struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type RequestHeader struct {
	ClientId string `json:"clientId"`
}

type Slice struct {
	AllCarriers               bool   `json:"allCarriers"`
	Cabin                     string `json:"cabin"`
	DepartureDate             string `json:"departureDate"`
	Destination               string `json:"destination"`
	DestinationNearbyAirports bool   `json:"destinationNearbyAirports"`
	MaxStops                  *int   `json:"maxStops"`
	Origin                    string `json:"origin"`
	OriginNearbyAirports      bool   `json:"originNearbyAirports"`
}

type TripOptions struct {
	CorporateBooking bool         `json:"corporateBooking"`
	FareType         string       `json:"fareType"`
	Locale           string       `json:"locale"`
	PointOfSale      *interface{} `json:"pointOfSale"`
	SearchType       string       `json:"searchType"`
}

type QueryParams struct {
	SliceIndex  int    `json:"sliceIndex"`
	SessionId   string `json:"sessionId"`
	SolutionSet string `json:"solutionSet"`
	SolutionId  string `json:"solutionId"`
}
