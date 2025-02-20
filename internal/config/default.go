package config

import (
	"os"
)

func GetFromEnv(varname string) string {
	return os.Getenv(varname)
}
