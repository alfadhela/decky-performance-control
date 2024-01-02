package utils

import (
	"os/exec"
)

// Execute command
func Execute(name string, arg ...string) string {
	command := exec.Command(name, arg...)
	output, err := command.Output()
	if err != nil {
		return ""
	}
	return string(output)
}
