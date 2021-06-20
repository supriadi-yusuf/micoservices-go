package config

import "os"

const (
	LogLevel = "info"
)

func IsProduction() bool {
	return os.Getenv("GO_ENVIRONMENT") == "production"
}
