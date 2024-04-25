package flights

type Output struct {
	Error                string           `json:"error"`
	FareBenefits         []string         `json:"fareBenefits"`
	FlexibleFareProducts []string         `json:"flexibleFareProducts"`
	LowestFareProducts   []string         `json:"lowestFareProducts"`
	Products             []string         `json:"products"`
	Utag                 Utag             `json:"utag"`
	Slices               []Slice2         `json:"slices"`
	ResponseMetadata     ResponseMetadata `json:"responseMetadata"`
}
type ResponseMetadata struct {
	SearchAirports         []Airport     `json:"searchAirports"`
	DepartureDate          string        `json:"departureDate"`
	Destination            Airport       `json:"destination"`
	Direction              string        `json:"direction"`
	Duration               int           `json:"duration"`
	International          bool          `json:"international"`
	Origin                 Airport       `json:"origin"`
	ResultId               string        `json:"resultId"`
	SearchType             string        `json:"searchType"`
	SessionId              string        `json:"sessionId"`
	SolutionSet            string        `json:"solutionSet"`
	Notifications          []interface{} `json:"notifications"` // Using interface{} as placeholder for unknown structure
	ShowFlagship           bool          `json:"showFlagship"`
	CorporateFare          bool          `json:"corporateFare"`
	SliceCount             int           `json:"sliceCount"`
	PricedSliceIndex       int           `json:"pricedSliceIndex"`
	Analytics              Analytics     `json:"analytics"`
	TripType               string        `json:"tripType"`
	StatusMessage          string        `json:"statusMessage"`
	SelectedProductChanged bool          `json:"selectedProductChanged"`
	RoundTrip              bool          `json:"roundTrip"`
	MultiCity              bool          `json:"multiCity"`
	MixedCabin             bool          `json:"mixedCabin"`
	MixedFare              bool          `json:"mixedFare"`
	LongHaulMarket         bool          `json:"longHaulMarket"`
	BusinessPlusOffered    bool          `json:"businessPlusOffered"`
	PricedSlice            int           `json:"pricedSlice"`
}

type Airport struct {
	City        string `json:"city"`
	CityName    string `json:"cityName"`
	Code        string `json:"code"`
	CountryCode string `json:"countryCode"`
	Name        string `json:"name"`
	StateCode   string `json:"stateCode"`
	Domestic    bool   `json:"domestic"`
}

type Analytics struct {
	BasicEconomy     FareDetails `json:"BASIC_ECONOMY"`
	Coach            FareDetails `json:"COACH"`
	CoachFlexible    FareDetails `json:"COACH_FLEXIBLE"`
	Business         FareDetails `json:"BUSINESS"`
	BusinessFlexible FareDetails `json:"BUSINESS_FLEXIBLE"`
}

type FareDetails struct {
	CheapestAmount float64 `json:"cheapestAmount"`
	Duration       int     `json:"duration"`
	Connections    int     `json:"connections"`
}
type Utag struct {
	AdultPassengers            string   `json:"adult_passengers"`
	AppBuildVersion            string   `json:"app_build_version"`
	AppName                    string   `json:"app_name"`
	AppRegion                  string   `json:"app_region"`
	Channel                    string   `json:"channel"`
	CurrentSlice               string   `json:"current_slice"`
	DestinationCountry         string   `json:"destination_country"`
	DestinationState           string   `json:"destination_state"`
	ItaQueryResultID           string   `json:"ita_query_result_id"`
	LoginStatus                string   `json:"login_status"`
	LowestSellingFarePrice     string   `json:"lowest_selling_fare_price"`
	LowestSellingFareType      string   `json:"lowest_selling_fare_type"`
	MatrixFareTypes            []string `json:"matrix_fare_types"`
	OriginCountry              string   `json:"origin_country"`
	OriginState                string   `json:"origin_state"`
	PageName                   string   `json:"page_name"`
	ProductsOffered            []string `json:"products_offered"`
	ProductsOfferedLowestPrice []string `json:"products_offered_lowestprice"`
	RouteType                  string   `json:"route_type"`
	SearchAdvancedDays         string   `json:"search_advanced_days"`
	SearchCabinType            string   `json:"search_cabin_type"`
	SearchCarrierOptions       string   `json:"search_carrier_options"`
	SearchDepartureDate        string   `json:"search_departure_date"`
	SearchDestinationCity      string   `json:"search_destination_city"`
	SearchMethod               string   `json:"search_method"`
	SearchNumberOfSlices       string   `json:"search_number_of_slices"`
	SearchOriginCity           string   `json:"search_origin_city"`
	SearchProduct              string   `json:"search_product"`
	SearchReturnDate           string   `json:"search_return_date"`
	SearchTripDuration         string   `json:"search_trip_duration"`
	SellingFarePriceAvailable  string   `json:"selling_fare_price_available"`
	SiteCountry                string   `json:"site_country"`
	SiteCurrency               string   `json:"site_currency"`
	SiteIndicator              string   `json:"site_indicator"`
	SiteLanguage               string   `json:"site_language"`
	SliceDate                  string   `json:"slice_date"`
	SliceOnd                   string   `json:"slice_ond"`
	TealiumEnvironment         string   `json:"tealium_environment"`
	TealiumProfile             string   `json:"tealium_profile"`
	TicketType                 string   `json:"ticket_type"`
	TimeStamp                  string   `json:"time_stamp"`
	TripType                   string   `json:"trip_type"`
	TrueOnd                    string   `json:"true_ond"`
	TrueclientIP               string   `json:"trueclient_ip"`
}

type Slice2 struct {
	DurationInMinutes               int                       `json:"durationInMinutes"`
	Segments                        []Segment                 `json:"segments"`
	Alerts                          []interface{}             `json:"alerts"`
	ID                              int                       `json:"id"`
	FilterData                      interface{}               `json:"filterData"`
	HasCorporateFare                bool                      `json:"hasCorporateFare"`
	CheapestPrice                   PriceDetails              `json:"cheapestPrice"`
	PricingDetail                   []PriceDetails            `json:"pricingDetail"`
	ProductGroups                   map[string][]ProductGroup `json:"productGroups"`
	ProductPricing                  []ProductPricing          `json:"productPricing"`
	SystemwideUpgrades              interface{}               `json:"systemwideUpgrades"`
	EconomyWebSpecialIncluded       bool                      `json:"economyWebSpecialIncluded"`
	WebSpecialIncluded              bool                      `json:"webSpecialIncluded"`
	DepartureDateTime               string                    `json:"departureDateTime"`
	CarrierNames                    []string                  `json:"carrierNames"`
	ProductDetails                  []ProductDetail           `json:"productDetails"`
	Origin                          Location                  `json:"origin"`
	ArrivalDateTime                 string                    `json:"arrivalDateTime"`
	Destination                     Location                  `json:"destination"`
	Stops                           int                       `json:"stops"`
	AllSegmentsOperatedByAA         bool                      `json:"allSegmentsOperatedByAA"`
	ConnectingCities                []interface{}             `json:"connectingCities"`
	OperationalDisclosureForDisplay bool                      `json:"operationalDisclosureForDisplay"`
	AllSegmentsOA                   bool                      `json:"allSegmentsOA"`
	ArrivesNextDay                  int                       `json:"arrivesNextDay"`
}
type ProductGroup struct {
	Main    []ProductDetails `json:"MAIN"`
	Premium []ProductDetails `json:"PREMIUM"`
}
type ProductDetails struct {
	PerPassengerAwardPoints    int                 `json:"perPassengerAwardPoints"`
	TripType                   string              `json:"tripType"`
	AllPassengerTaxesAndFees   MonetaryValue       `json:"allPassengerTaxesAndFees"`
	ExtendedFareCode           string              `json:"extendedFareCode"`
	PerPassengerTaxesAndFees   MonetaryValue       `json:"perPassengerTaxesAndFees"`
	ProductBenefits            string              `json:"productBenefits"`
	ProductType                string              `json:"productType"`
	ProductTypeExt             interface{}         `json:"productTypeExt"`
	ProductGroup               string              `json:"productGroup"`
	SeatsRemaining             int                 `json:"seatsRemaining"`
	SolutionID                 string              `json:"solutionID"`
	RefundableProducts         []RefundableProduct `json:"refundableProducts"`
	BenefitKey                 string              `json:"benefitKey"`
	DynamicFare                bool                `json:"dynamicFare"`
	Flexible                   bool                `json:"flexible"`
	CorporateFare              bool                `json:"corporateFare"`
	LieFlat                    bool                `json:"lieFlat"`
	LowestPriceForProductGroup bool                `json:"lowestPriceForProductGroup"`
	AllPassengerDisplayTotal   MonetaryValue       `json:"allPassengerDisplayTotal"`
	PerPassengerDisplayTotal   MonetaryValue       `json:"perPassengerDisplayTotal"`
	ProductAvailable           bool                `json:"productAvailable"`
	BasicEconomyPlus           bool                `json:"basicEconomyPlus"`
	Flagship                   bool                `json:"flagship"`
	BusinessPlus               bool                `json:"businessPlus"`
	WebSpecial                 bool                `json:"webSpecial"`
	MustBookAtAirport          bool                `json:"mustBookAtAirport"`
	FlagshipRiskyConnection    bool                `json:"flagshipRiskyConnection"`
}
type Segment struct {
	Alerts              []interface{} `json:"alerts"`
	Fares               []interface{} `json:"fares"`
	Flight              Flight        `json:"flight"`
	Legs                []Leg         `json:"legs"`
	MarriedSegmentIndex interface{}   `json:"marriedSegmentIndex"`
	DepartureDateTime   string        `json:"departureDateTime"`
	Origin              Location      `json:"origin"`
	ArrivalDateTime     string        `json:"arrivalDateTime"`
	Destination         Location      `json:"destination"`
	ChangeOfGauge       bool          `json:"changeOfGauge"`
	ThroughFlight       bool          `json:"throughFlight"`
}
type PriceDetails struct {
	PerPassengerAwardPoints    int                 `json:"perPassengerAwardPoints"`
	TripType                   string              `json:"tripType"`
	AllPassengerTaxesAndFees   MonetaryValue       `json:"allPassengerTaxesAndFees"`
	ExtendedFareCode           string              `json:"extendedFareCode"`
	PerPassengerTaxesAndFees   MonetaryValue       `json:"perPassengerTaxesAndFees"`
	ProductBenefits            string              `json:"productBenefits"`
	ProductType                string              `json:"productType"`
	ProductTypeExt             interface{}         `json:"productTypeExt"`
	ProductGroup               string              `json:"productGroup"`
	SeatsRemaining             int                 `json:"seatsRemaining"`
	SolutionID                 string              `json:"solutionID"`
	RefundableProducts         []RefundableProduct `json:"refundableProducts"`
	BenefitKey                 string              `json:"benefitKey"`
	DynamicFare                bool                `json:"dynamicFare"`
	Flexible                   bool                `json:"flexible"`
	CorporateFare              bool                `json:"corporateFare"`
	LieFlat                    bool                `json:"lieFlat"`
	LowestPriceForProductGroup bool                `json:"lowestPriceForProductGroup"`
	AllPassengerDisplayTotal   MonetaryValue       `json:"allPassengerDisplayTotal"`
	PerPassengerDisplayTotal   MonetaryValue       `json:"perPassengerDisplayTotal"`
	ProductAvailable           bool                `json:"productAvailable"`
	BasicEconomyPlus           bool                `json:"basicEconomyPlus"`
	Flagship                   bool                `json:"flagship"`
	BusinessPlus               bool                `json:"businessPlus"`
	WebSpecial                 bool                `json:"webSpecial"`
	MustBookAtAirport          bool                `json:"mustBookAtAirport"`
	FlagshipRiskyConnection    bool                `json:"flagshipRiskyConnection"`
}

type MonetaryValue struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

type ProductPricing struct {
	RegularPrice    PriceDetails `json:"regularPrice"`
	WebSpecialPrice interface{}  `json:"webSpecialPrice"`
	CheapestPrice   PriceDetails `json:"cheapestPrice"`
}
type Origin struct {
	City        string `json:"city"`
	CityName    string `json:"cityName"`
	Code        string `json:"code"`
	CountryCode string `json:"countryCode"`
	Name        string `json:"name"`
	StateCode   string `json:"stateCode"`
	Domestic    bool   `json:"domestic"`
}
type Flight struct {
	CarrierCode  string `json:"carrierCode"`
	CarrierName  string `json:"carrierName"`
	FlightNumber string `json:"flightNumber"`
}

type Leg struct {
	Aircraft                Aircraft          `json:"aircraft"`
	ArrivalDateTime         string            `json:"arrivalDateTime"`
	BrazilOnTimePerformance interface{}       `json:"brazilOnTimePerformance"`
	ConnectionTimeInMinutes int               `json:"connectionTimeInMinutes"`
	DepartureDateTime       string            `json:"departureDateTime"`
	Destination             Location          `json:"destination"`
	DurationInMinutes       int               `json:"durationInMinutes"`
	Flight                  interface{}       `json:"flight"`
	OnTimePerformance       OnTimePerformance `json:"onTimePerformance"`
	OperationalDisclosure   string            `json:"operationalDisclosure"`
	Origin                  Location          `json:"origin"`
	ProductDetails          []ProductDetail   `json:"productDetails"`
	Alerts                  []interface{}     `json:"alerts"`
	Amenities               []string          `json:"amenities"`
	ArrivesNextDay          int               `json:"arrivesNextDay"`
	Domestic                bool              `json:"domestic"`
	Brazilian               bool              `json:"brazilian"`
	AircraftCode            string            `json:"aircraftCode"`
}

type Aircraft struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	ShortName string `json:"shortName"`
}

type OnTimePerformance struct {
	PerformanceData interface{} `json:"performanceData"`
	WarningRequired bool        `json:"warningRequired"`
}

type Location struct {
	City        string `json:"city"`
	CityName    string `json:"cityName"`
	Code        string `json:"code"`
	CountryCode string `json:"countryCode"`
	Name        string `json:"name"`
	StateCode   string `json:"stateCode"`
	Domestic    bool   `json:"domestic"`
}
type ProductDetail struct {
	BookingCode  string        `json:"bookingCode"`
	CabinType    string        `json:"cabinType"`
	Meals        []string      `json:"meals"`
	ProductType  string        `json:"productType"`
	WebSpecial   bool          `json:"webSpecial"`
	Flagship     bool          `json:"flagship"`
	BusinessPlus bool          `json:"businessPlus"`
	Upgradeable  bool          `json:"upgradeable"`
	Alerts       []interface{} `json:"alerts"`
}
type Destination struct {
	City        string `json:"city"`
	CityName    string `json:"cityName"`
	Code        string `json:"code"`
	CountryCode string `json:"countryCode"`
	Name        string `json:"name"`
	StateCode   string `json:"stateCode"`
	Domestic    bool   `json:"domestic"`
}
type Product struct {
	PerPassengerAwardPoints    int                 `json:"perPassengerAwardPoints"`
	TripType                   string              `json:"tripType"`
	AllPassengerTaxesAndFees   MonetaryAmount      `json:"allPassengerTaxesAndFees"`
	ExtendedFareCode           string              `json:"extendedFareCode"`
	PerPassengerTaxesAndFees   MonetaryAmount      `json:"perPassengerTaxesAndFees"`
	ProductBenefits            string              `json:"productBenefits"`
	ProductType                string              `json:"productType"`
	ProductTypeExt             *string             `json:"productTypeExt"`
	ProductGroup               string              `json:"productGroup"`
	SeatsRemaining             int                 `json:"seatsRemaining"`
	SolutionID                 string              `json:"solutionID"`
	RefundableProducts         []RefundableProduct `json:"refundableProducts"`
	BenefitKey                 string              `json:"benefitKey"`
	DynamicFare                bool                `json:"dynamicFare"`
	AllPassengerDisplayTotal   MonetaryAmount      `json:"allPassengerDisplayTotal"`
	PerPassengerDisplayTotal   MonetaryAmount      `json:"perPassengerDisplayTotal"`
	LieFlat                    bool                `json:"lieFlat"`
	MustBookAtAirport          bool                `json:"mustBookAtAirport"`
	FlagshipRiskyConnection    bool                `json:"flagshipRiskyConnection"`
	Flexible                   bool                `json:"flexible"`
	CorporateFare              bool                `json:"corporateFare"`
	WebSpecial                 bool                `json:"webSpecial"`
	BasicEconomyPlus           bool                `json:"basicEconomyPlus"`
	Flagship                   bool                `json:"flagship"`
	BusinessPlus               bool                `json:"businessPlus"`
	ProductAvailable           bool                `json:"productAvailable"`
	LowestPriceForProductGroup bool                `json:"lowestPriceForProductGroup"`
}

type RefundableProduct struct {
	CorporateFare bool           `json:"corporateFare"`
	Indicator     string         `json:"indicator"`
	JsonKey       string         `json:"jsonKey"`
	ProductType   string         `json:"productType"`
	RefundAmount  MonetaryAmount `json:"refundAmount"`
	SolutionID    string         `json:"solutionID"`
}
type MonetaryAmount struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}
