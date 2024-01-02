package utils

import (
	"backend/src/pkg/model"
	"encoding/json"
	"fmt"
	"os"
)

func getSettingsDir() string {
	home := os.Getenv("HOME")
	//
	if _, err := os.Stat(fmt.Sprintf("%s/.config", home)); err != nil {
		os.Mkdir(fmt.Sprintf("%s/.config", home), 0700)
	}
	return fmt.Sprintf("%s/.config/decky-performance-control", home)
}

func createSettingsFolder() {
	_, err := os.Stat(getSettingsDir())
	if err == nil {
		return
	}
	os.Mkdir(getSettingsDir(), 0700)
}

func ReadSettings(appId uint) *model.Settings {
	createSettingsFolder()
	appFile, err := os.ReadFile(fmt.Sprintf("%s/%d.json", getSettingsDir(), appId))
	if err != nil {
		fmt.Printf("Read file: %e\n", err)
	}
	var settings model.Settings
	if err == nil {
		json.Unmarshal(appFile, &settings)
		return &settings
	}
	settings.AppID = appId
	settings.Boost = false
	settings.TDP = false
	if DeviceInfo() != nil {
		settings.TDPLimit = DeviceInfo().MaxTDP
	}
	settings.EGPU = false
	SaveSettings(&settings)
	return &settings
}

func SaveSettings(settings *model.Settings) {
	createSettingsFolder()
	//
	d, _ := json.Marshal(settings)
	if err := os.WriteFile(fmt.Sprintf("%s/%d.json", getSettingsDir(), settings.AppID), d, 0644); err != nil {
		fmt.Printf("Read file: %e\n", err)
	}
}
