package api

import (
	"backend/src/internal/constant"
	"backend/src/internal/utils"
	"encoding/json"
	"net/http"
	"os"
	"strconv"

	"github.com/yousuf64/shift"
)

func eGPU(w http.ResponseWriter, r *http.Request, route shift.Route) error {
	w.Header().Add("Content-Type", "application/json")
	if _, err := os.Stat("/dev/dri/card2"); err != nil {
		body, _ := json.Marshal(map[string]any{
			"value": false,
		})
		_, err := w.Write(body)
		return err
	}
	vendorID := utils.Execute("cat", "/sys/class/drm/card2/device/vendor")
	uVendorID, _ := strconv.Atoi(vendorID)
	body, _ := json.Marshal(map[string]any{
		"value":  true,
		"vendor": constant.PCIVendors[uint(uVendorID)],
	})
	_, err := w.Write(body)
	return err
}

func enable_eGPU(w http.ResponseWriter, r *http.Request, route shift.Route) error {
	// TODO: Add DRI_PRIME=1
	// to {home}/.steam/steam/userdata/{steam_user_id}/config/localconfig.vdf
	// keep file permissions
	return nil
}

func disable_eGPU(w http.ResponseWriter, r *http.Request, route shift.Route) error {
	return nil
}
