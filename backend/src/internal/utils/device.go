package utils

import (
	"backend/src/internal/constant"
	"backend/src/pkg/model"
	"strings"
	"sync"
)

var device *model.Device
var once sync.Once

func DeviceInfo() *model.Device {
	once.Do(func() {
		// Get Vendor
		vendor := Execute("cat", "/sys/class/dmi/id/board_vendor")
		if len(vendor) == 0 {
			return
		}
		// Get Product
		product := Execute("cat", "/sys/class/dmi/id/product_name")
		if len(product) == 0 {
			return
		}
		// TODO: Get CPU info and features
		_device := model.Device{
			Name:   strings.TrimSuffix(product, "\n"),
			Vendor: strings.TrimSuffix(vendor, "\n"),
		}
		if d, ok := constant.Devices[_device.Name]; ok {
			_device.MaxTDP = d.MaxTDP
			_device.MinTDP = d.MinTDP
			_device.TDP = true
			_device.RGB = d.RGB
		}
		device = &_device
	})
	return device
}
