package model

type Settings struct {
	AppID    uint `json:"app_id"`
	Boost    bool `json:"boost"`
	TDP      bool `json:"tdp"`
	TDPLimit uint `json:"tdp_limit"`
	EGPU     bool `json:"egpu"`
}
