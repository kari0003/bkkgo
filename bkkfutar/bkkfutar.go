package bkkfutar

type ScheduleForStopParams struct {
	// API kulcs Example: apaiary-test.
	Key string `json:"key"`
	// API verzió Example: 3.
	Version int `json:"version"`
	// Az API-t használó alkalmazás verziója Example: apiary-1.0.
	AppVersion string `json:"appVersion"`
	// A válaszban szereplő referenciák típusa ,-vel összefűzve. Default: true.
	IncludeReferences bool `json:"includeReferences"`
	// Megálló azonosítója. Example: KKK_F00001.
	StopID string `json:"stopId"`
	// Kizárólag induló járatok megadása a válaszban. Example: false.
	OnlyDepartures bool `json:"onlyDepartures"`
	// Dátum. Example: 20140920.
	Date string `json:"date"`
}

type ScheduleForStopResponse struct {
	Entry ScheduleForStopData `json:"entry"`
}

type ScheduleForStopData struct {
	StopID      string     `json:"stopId"`
	ServiceDate string     `json:"serviceDate"`
	AlertIds    []string   `json:"alertIds"`
	Schedules   []Schedule `json:"schedules"`
}

type Schedule struct {
	RouteID    string      `json:"routeId"`
	Directions []Direction `json:"directions"`
}

type Direction struct {
	DirectionID string                    `json:"directionId"`
	Groups      map[string]DirectionGroup `json:"groups"`
	StopTimes   []StopTime                `json:"stopTimes"`
}

type DirectionGroup struct {
	GroupID     string `json:"groupId"`
	Headsign    string `json:"headsign"`
	Description string `json:"description"`
}

type StopTime struct {
	ArrivalTime            int      `json:"arrivalTime"`
	DepartureTime          int      `json:"departureTime"`
	PredictedArrivalTime   int      `json:"predictedArrivalTime"`
	PredictedDepartureTime int      `json:"predictedDepartureTime"`
	TripID                 string   `json:"tripId"`
	ServiceDate            string   `json:"serviceDate"`
	WheelchairAccessible   bool     `json:"wheelchairAccessible"`
	GroupIDs               []string `json:"groupIds"`
}

type BkkFutar interface {
	ScheduleForStop(params ScheduleForStopParams) ScheduleForStopResponse
}
