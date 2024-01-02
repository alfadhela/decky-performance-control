package model

type Device struct {
	Name   string `json:"name,omitempty"`
	Vendor string `json:"vendor,omitempty"`
	TDP    bool   `json:"tdp,omitempty"`
	MaxTDP uint   `json:"max_tdp,omitempty"`
	MinTDP uint   `json:"min_tdp,omitempty"`
	RGB    bool   `json:"rgb,omitempty"`
	Boost  bool   `json:"boost,omitempty"`
}

// SustainedPowerLimit (stapm-limit)
func (d *Device) SustainedPowerLimit(tdp uint) uint {
	if tdp > d.MaxTDP {
		tdp = d.MaxTDP
	}
	return tdp * 1000
}

// ActualPowerLimit (fast-limit)
func (d *Device) ActualPowerLimit(tdp uint) uint {
	if tdp > d.MaxTDP {
		tdp = d.MaxTDP
	}
	return uint(float64(tdp)*0.75) * 1000
}

// AveragePowerLimit (slow-limit)
func (d *Device) AveragePowerLimit(tdp uint) uint {
	if tdp > d.MaxTDP {
		tdp = d.MaxTDP
	}
	return uint(float64(tdp)*1.25) * 1000
}

// FixedPowerLimit
func (d *Device) FixedPowerLimit(tdp uint) uint {
	if tdp > d.MaxTDP {
		tdp = d.MaxTDP
	}
	if tdp < d.MinTDP {
		tdp = d.MinTDP
	}
	return tdp * 1000
}
