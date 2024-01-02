package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/yousuf64/shift"
)

func getRGB(w http.ResponseWriter, r *http.Request, route shift.Route) error {
	w.Header().Add("Content-Type", "application/json")
	mode := route.Params.Get("mode")
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	resp, err := client.Get(fmt.Sprintf("http://127.0.0.1:21371/get/%s", mode))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(map[string]any{
			"error": err.Error(),
		})
		_, err := w.Write(body)
		return err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(map[string]any{
			"error": err.Error(),
		})
		_, err := w.Write(body)
		return err
	}
	colors := strings.Split(string(body), ":")
	red, _ := strconv.ParseUint(colors[0], 10, 8)
	green, _ := strconv.ParseUint(colors[1], 10, 8)
	blue, _ := strconv.ParseUint(colors[2], 10, 8)
	body, _ = json.Marshal(map[string]any{
		"hex": fmt.Sprintf("#%02x%02x%02x", red, green, blue),
	})
	_, err = w.Write(body)
	return err
}

func setRGB(w http.ResponseWriter, r *http.Request, route shift.Route) error {
	w.Header().Add("Content-Type", "application/json")
	mode := route.Params.Get("mode")
	var color colorRequest
	if err := json.NewDecoder(r.Body).Decode(&color); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(map[string]any{
			"error": err.Error(),
		})
		_, err := w.Write(body)
		return err
	}
	color.HEX = strings.Replace(color.HEX, "#", "", 1)
	values, err := strconv.ParseUint(color.HEX, 16, 32)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(map[string]any{
			"error": "invalid hex color",
		})
		_, err := w.Write(body)
		return err
	}
	rgb := []uint8{
		uint8(values >> 16),         //Red
		uint8((values >> 8) & 0xFF), //Green
		uint8(values & 0xFF),        //Blue
	}
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(map[string]any{
			"error": err.Error(),
		})
		_, err := w.Write(body)
		return err
	}
	client := http.Client{Timeout: time.Duration(1) * time.Second}
	_, err = client.Get(fmt.Sprintf("http://127.0.0.1:21371/set/%s/%d/%d/%d", mode, rgb[0], rgb[1], rgb[2]))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		body, _ := json.Marshal(map[string]any{
			"error": err.Error(),
		})
		_, err := w.Write(body)
		return err
	}
	return nil
}

type colorRequest struct {
	HEX string `json:"hex,omitempty"`
}
