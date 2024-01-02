package utils

import "os"

func IsRyzenAdjFound() bool {
	_, err := os.Stat("/usr/bin/ryzenadj")
	return err == nil
}
