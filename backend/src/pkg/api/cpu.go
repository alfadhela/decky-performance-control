package api

import (
	"backend/src/internal/constant"
	"backend/src/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/yousuf64/shift"
)

func getTDPInformation(w http.ResponseWriter, r *http.Request, route shift.Route) error {
	fmt.Println("Getting device current tdp limit information")
	w.Header().Add("Content-Type", "application/json")
	//
	device := utils.DeviceInfo()
	if device == nil || device.MaxTDP == 0 {
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(map[string]any{
			"error": "device not supported",
		})
		_, err := w.Write(body)
		return err
	}
	// Check if ryzenadj
	if !utils.IsRyzenAdjFound() {
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(map[string]any{
			"error": "ryzenadj not found",
		})
		_, err := w.Write(body)
		return err
	}
	// Get using ryzenadj
	table := utils.Execute("/usr/bin/ryzenadj", "-i")
	if len(table) == 0 {
		body, _ := json.Marshal(TDPInformation{
			TDPLimit: uint(device.MaxTDP),
		})
		_, err := w.Write(body)
		return err
	}
	tdp, _ := strconv.ParseFloat(utils.FindInRyzenAdjTable(table, constant.StapmLimit), 64)
	body, _ := json.Marshal(TDPInformation{
		TDPLimit: uint(tdp),
	})
	_, err := w.Write(body)
	return err
}

type TDPInformation struct {
	TDPLimit uint `json:"tdp_limit,omitempty"`
}
