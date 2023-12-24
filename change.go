package advantageair

type Change interface {
	SetAirconState(aircon string, state AirconState) Change
	SetAirconMode(aircon string, mode AirconMode) Change
	SetAirconFan(aircon string, speed AirconFanSpeed) Change
	SetAirconZoneTemperature(aircon string, zone string, temperature int64) Change
	SetAirconZoneState(aircon string, zone string, state AirconZoneState) Change
}

func NewChange() Change { return &change{} }

type change map[string]any

func (c *change) SetAirconState(aircon string, state AirconState) Change {
	setRecursively(*c, state, "aircons", aircon, "info", "state")
	return c
}

func (c *change) SetAirconMode(aircon string, mode AirconMode) Change {
	setRecursively(*c, mode, "aircons", aircon, "info", "mode")
	return c
}

func (c *change) SetAirconFan(aircon string, speed AirconFanSpeed) Change {
	setRecursively(*c, speed, "aircons", aircon, "info", "fan")
	return c
}

func (c *change) SetAirconZoneTemperature(aircon string, zone string, temperature int64) Change {
	setRecursively(*c, temperature, "aircons", aircon, "zones", zone, "setTemp")
	return c
}

func (c *change) SetAirconZoneState(aircon string, zone string, state AirconZoneState) Change {
	setRecursively(*c, state, "aircons", aircon, "zones", zone, "state")
	return c
}
