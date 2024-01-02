package api

import (
	"backend/src/internal/utils"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/yousuf64/shift"
)

func device(w http.ResponseWriter, r *http.Request, route shift.Route) error {
	fmt.Println("Getting device information")
	device := utils.DeviceInfo()
	if device == nil {
		w.WriteHeader(http.StatusInternalServerError)
		return nil
	}
	w.Header().Add("Content-Type", "application/json")
	body, _ := json.Marshal(device)
	_, err := w.Write(body)
	return err
}
