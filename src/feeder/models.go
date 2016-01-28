package feeder

type TimeStruct struct {
	DateUTC   string `json:"dateUtc,omitempty"`
	DateLocal string `json:"dateLocal,omitempty"`
}

type ScheduleStruct struct {
	FlightType     string        `json:"flightType"`
	ServiceClasses string        `json:"serviceClasses"`
	Restrictions   string        `json:"restrictions"`
	Uplines        []interface{} `json:"uplines"`
	Downlines      []interface{} `json:"downlines"`
}

type FlightStatusUpdateStruct struct {
	UpdatedAt         TimeStruct          `json:"updatedAt"`
	Source            string              `json:"source"`
	UpdatedTextFields []map[string]string `json:"updatedTextFields"`
	UpdatedDateFields []map[string]string `json:"updatedDateFields"`
}

type DelayStruct struct {
	DepartureGateDelayMinutes   int64 `json:"departureGateDelayMinutes"`
	DepartureRunwayDelayMinutes int64 `json:"departureRunwayDelayMinutes"`
	ArrivalGateDelayMinutes     int64 `json:"arrivalGateDelayMinutes"`
	ArrivalRunwayDelayMinutes   int64 `json: "arrivalRunwayDelayMinutes"`
}

type FlightDurationsStruct struct {
	ScheduledBlockMinutes   int64 `json:"scheduledBlockMinutes"`
	ScheduledAirMinutes     int64 `json:"scheduledAirMinutes"`
	AirMinutes              int64 `json:"airMinutes"`
	ScheduledTaxiOutMinutes int64 `json:"scheduledTaxiOutMinutes"`
	TaxiOutMinutes          int64 `json:"taxiOutMinutes"`
	ScheduledTaxiInMinutes  int64 `json:"scheduledTaxiInMinutes"`
}

type AirportResourcesStruct struct {
	DepartureTerminal string `json:"departureTerminal"`
	DepartureGate     string `json:"departureGate"`
	ArrivalGate       string `json:"arrivalGate"`
}

type FlightStatusesStruct struct {
	FlightID               int64                      `json:"flightId"`
	CarrierFsCode          string                     `json:"carrierFsCode"`
	OperatingCarrierFsCode string                     `json:"operatingCarrierFsCode"`
	PrimaryCarrierFsCode   string                     `json:"primaryCarrierFsCode"`
	FlightNumber           string                     `json:"flightNumber"`
	DepartureAirportFsCode string                     `json:"departureAirportFsCode"`
	ArrivalAirportFsCode   string                     `json:"arrivalAirportFsCode"`
	DepartureDate          TimeStruct                 `json:"departureDate"`
	ArrivalDate            TimeStruct                 `json:"arrivalDate"`
	Status                 string                     `json:"status"`
	Schedule               ScheduleStruct             `json:"schedule"`
	OperationalTimes       map[string]TimeStruct      `json:"operationalTimes"`
	Codeshares             []map[string]string        `json:"codeshares"`
	Delays                 DelayStruct                `json:"delays"`
	FlightDurations        FlightDurationsStruct      `json:"flightDurations"`
	AirportResources       AirportResourcesStruct     `json:"airportResources"`
	FlightEquipment        map[string]string          `json:"flightEquipment"`
	FlightStatusUpdates    []FlightStatusUpdateStruct `json:"flightStatusUpdates"`
	IrregularOperations    []interface{}              `json:"irregularOperations"`
}

type RequestStruct struct {
	URL string `json:"url"`
}

type FlightResp struct {
	Request        RequestStruct          `json:"request"`
	FlightStatuses []FlightStatusesStruct `json:"flightStatuses"`
}

func (r *FlightResp) Convert() *FlightResult {

}

type FlightResult struct {
	Flights []Flight `json:"flights"`
}

// internal structs for storing flights. not quite what gets returned to the client
type Flight struct {
	Id           string  `json:"id"`
	ContentHash  string  `json:"content_hash"`
	CarrierCode  string  `json:"carrier_code"`
	CarrierName  *string `json:"carrier_name,omitempty"` // translatable string, not stored in stringd
	FlightNumber string  `json:"flight_number"`
	Legs         []Leg   `json:"legs"`
}

type Airport struct {
	Code      string  `json:"code"`
	City      *string `json:"city,omitempty"` // translatable string, not stored in stringd
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	UTCOffset int64   `json:"utc_offset"`
}

type Leg struct {
	Status                 *string `json:"status,omitempty"` // translatable string, not stored in stringd
	DepartureScheduledTime int64   `json:"departure_scheduled_time"`
	DepartureEstimatedTime int64   `json:"departure_estimated_time"`
	DepartureTerminal      string  `json:"departure_terminal"`
	DepartureGate          string  `json:"departure_gate"`
	DepartureAirport       Airport `json:"departure_airport"`
	ArrivalScheduledTime   int64   `json:"arrival_scheduled_time"`
	ArrivalEstimatedTime   int64   `json:"arrival_estimated_time"`
	ArrivalTerminal        string  `json:"arrival_terminal"`
	ArrivalGate            string  `json:"arrival_gate"`
	ArrivalAirport         Airport `json:"arrival_airport"`
}
