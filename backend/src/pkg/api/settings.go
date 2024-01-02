package api

import (
	"backend/src/internal/utils"
	"backend/src/pkg/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/yousuf64/shift"
)

func getSettings(w http.ResponseWriter, r *http.Request, route shift.Route) error {
	w.Header().Add("Content-Type", "application/json")
	fmt.Printf("Get settings for app-id:%s\n", route.Params.Get("app_id"))
	appId, _ := strconv.Atoi(route.Params.Get("app_id"))
	settings := utils.ReadSettings(uint(appId))
	// applySettings(settings)
	body, _ := json.Marshal(settings)
	_, err := w.Write(body)
	return err
}

func saveSettings(w http.ResponseWriter, r *http.Request, route shift.Route) error {
	w.Header().Add("Content-Type", "application/json")
	fmt.Printf("Setting settings for app-id:%s\n", route.Params.Get("app_id"))
	appId, _ := strconv.Atoi(route.Params.Get("app_id"))
	var settings model.Settings
	json.NewDecoder(r.Body).Decode(&settings)
	settings.AppID = uint(appId)
	utils.SaveSettings(&settings)
	applySettings(&settings)
	body, _ := json.Marshal(settings)
	_, err := w.Write(body)
	return err
}

func applySettings(settings *model.Settings) {
	// Check before applying settings
	device := utils.DeviceInfo()
	if device == nil || !device.TDP || settings.TDPLimit == 0 {
		return
	}
	// Apply Settings
	tdpLimit := settings.TDPLimit
	if !settings.TDP {
		tdpLimit = device.MaxTDP
	}
	utils.Execute("/usr/bin/ryzenadj",
		fmt.Sprintf("--stapm-limit=%d", device.SustainedPowerLimit(tdpLimit)),
		fmt.Sprintf("--fast-limit=%d", device.FixedPowerLimit(tdpLimit)),
		fmt.Sprintf("--slow-limit=%d", device.FixedPowerLimit(tdpLimit)))
}
