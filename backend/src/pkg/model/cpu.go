package model

type TDPInformation struct {
	TDP       bool `json:"tdp,omitempty"`
	TDPLimit  uint `json:"tdp_limit,omitempty"`
	Temp      uint `json:"temp,omitempty"`
	TempLimit uint `json:"temp_limit,omitempty"`
}
