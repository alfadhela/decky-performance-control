package utils

import (
	"backend/src/internal/constant"
	"strings"
)

func FindInRyzenAdjTable(output string, name constant.RyzenAdjName) string {
	value := output[strings.Index(output, string(name)) : strings.Index(output, string(name))+53]
	value = strings.ReplaceAll(value, " ", "")
	value = value[strings.Index(value, "|")+1 : strings.LastIndex(value, "|")]
	return value
}
