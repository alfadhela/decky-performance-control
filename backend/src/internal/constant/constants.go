package constant

import "backend/src/pkg/model"

var Devices = map[string]model.Device{
	"AIR": {
		MaxTDP: 18,
		MinTDP: 5,
		RGB:    true,
	},
	"AIR Plus": {
		MaxTDP: 28,
		MinTDP: 5,
		RGB:    false,
	},
	"AIR Pro": {
		MaxTDP: 20,
		MinTDP: 5,
		RGB:    true,
	},
	"AYANEO 2": {
		MaxTDP: 30,
		MinTDP: 5,
		RGB:    true,
	},
	"AYANEO 2S": {
		MaxTDP: 30,
		MinTDP: 5,
		RGB:    false,
	},
	"GEEK": {
		MaxTDP: 30,
		MinTDP: 5,
		RGB:    true,
	},
	"GEEK 1S": {
		MaxTDP: 30,
		MinTDP: 5,
		RGB:    false,
	},
}

type RyzenAdjName string

const (
	StapmLimit        RyzenAdjName = "STAPM LIMIT"
	StapmValue        RyzenAdjName = "STAPM VALUE"
	FastLimit         RyzenAdjName = "PPT LIMIT FAST"
	FastValue         RyzenAdjName = "PPT VALUE FAST"
	SlowLimit         RyzenAdjName = "PPT LIMIT SLOW"
	Slowvalue         RyzenAdjName = "PPT VALUE FAST"
	StapmTimeConst    RyzenAdjName = "StapmTimeConst"
	SlowPPTTimeConst  RyzenAdjName = "SlowPPTTimeConst"
	ApuSlowLimit      RyzenAdjName = "PPT LIMIT APU"
	ApuSlowValue      RyzenAdjName = "PPT VALUE APU"
	VRMCurrentLimit   RyzenAdjName = "TDC LIMIT VDD"
	VRMCurrentValue   RyzenAdjName = "TDC VALUE VDD"
	VRMSoCLimit       RyzenAdjName = "TDC LIMIT SOC"
	VRMSoCValue       RyzenAdjName = "TDC VALUE SOC"
	VRMMaxLimit       RyzenAdjName = "EDC LIMIT VDD"
	VRMMaxValue       RyzenAdjName = "EDC VALUE VDD"
	VRMSoCMaxLimit    RyzenAdjName = "EDC LIMIT SOC"
	VRMSoCMaxValue    RyzenAdjName = "EDC VALUE SOC"
	TempLimit         RyzenAdjName = "THM LIMIT CORE"
	TempValue         RyzenAdjName = "THM VALUE CORE"
	ApuSkinTempLimit  RyzenAdjName = "STT LIMIT APU"
	ApuSkinTempValue  RyzenAdjName = "STT VALUE APU"
	DGpuSkinTempLimit RyzenAdjName = "STT LIMIT dGPU"
	DGpuSkinTempValue RyzenAdjName = "STT VALUE dGPU"
	PowerSaving       RyzenAdjName = "CCLK Boost SETPOINT"
	MaxPerformance    RyzenAdjName = "CCLK BUSY VALUE"
)

var PCIVendors = map[uint]string{
	0x8086: "Intel",
	0x1002: "AMD",
	0x10de: "Nvidia",
}
