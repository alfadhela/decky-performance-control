package api

import (
	"github.com/yousuf64/shift"
)

func Register(g *shift.Group) {
	// Device information
	g.GET("/device", device)
	// Settings
	g.GET("/settings/:app_id", getSettings)
	g.PUT("/settings/:app_id", saveSettings)
	// CPU
	g.GET("/tdp-information", getTDPInformation)
	// eGPU
	g.GET("/egpu", eGPU)
	g.GET("/egpu/enable", enable_eGPU)
	g.GET("/egpu/disable", disable_eGPU)
	// RGB
	g.GET("/rgb/:mode", getRGB)
	g.PUT("/rgb/:mode", setRGB)
}
