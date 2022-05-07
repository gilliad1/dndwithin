package utils

import (
	"os"
	"strconv"
)

func GetenvStringOrDefault(varName, varDef string) string {
	ret := os.Getenv(varName)
	if ret == "" {
		return varDef
	}
	return ret
}

func GetenvBoolOrDefault(varName string, varDef bool) bool {
	ret := os.Getenv(varName)
	if ret == "" {
		return varDef
	}
	b, _ := strconv.ParseBool(ret)
	return b
}
