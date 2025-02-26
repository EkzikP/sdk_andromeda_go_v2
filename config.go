package andromeda

import "time"

const (
	endpointGetSites             = "/Sites"
	endpointGetCustomers         = "/Customers"
	endpointCheckPanic           = "/CheckPanic"
	endpointMyAlarm              = "/MyAlarm"
	endpointGetUserObjectMyAlarm = "/MyAlarm/UserObjects"
	endpointGetParts             = "/Parts"
	endpointGetZones             = "/Zones"

	defaultTimeout = 5 * time.Second
)

// Config содержит обязательные параметры для всех запросов к API
type Config struct {
	ApiKey string `json:"-"`
	Host   string `json:"-"`
}
