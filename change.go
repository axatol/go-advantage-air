package advantageair

import "github.com/axatol/go-advantage-air/internal/util"

// Change is the interface for building a change to the system.
type Change interface {
	SetAirconState(aircon string, state AirconState) Change                        // SetAirconState turns an aircon on or off.
	SetAirconMode(aircon string, mode AirconMode) Change                           // SetAirconMode sets the mode of an aircon, e.g. heating/cooling.
	SetAirconFan(aircon string, speed AirconFanSpeed) Change                       // SetAirconFan sets the fan speed of an aircon.
	SetAirconZoneTemperature(aircon string, zone string, temperature int64) Change // SetAirconZoneTemperature sets the temperature of a zone.
	SetAirconZoneState(aircon string, zone string, state AirconZoneState) Change   // SetAirconZoneState opens or closes airflow to a zone.
}

func NewChange() Change { return &change{} }

type change map[string]any

func (c *change) SetAirconState(aircon string, state AirconState) Change {
	util.SetRecursively(*c, state, "aircons", aircon, "info", "state")
	return c
}

func (c *change) SetAirconMode(aircon string, mode AirconMode) Change {
	util.SetRecursively(*c, mode, "aircons", aircon, "info", "mode")
	return c
}

func (c *change) SetAirconFan(aircon string, speed AirconFanSpeed) Change {
	util.SetRecursively(*c, speed, "aircons", aircon, "info", "fan")
	return c
}

func (c *change) SetAirconZoneTemperature(aircon string, zone string, temperature int64) Change {
	util.SetRecursively(*c, temperature, "aircons", aircon, "zones", zone, "setTemp")
	return c
}

func (c *change) SetAirconZoneState(aircon string, zone string, state AirconZoneState) Change {
	util.SetRecursively(*c, state, "aircons", aircon, "zones", zone, "state")
	return c
}
