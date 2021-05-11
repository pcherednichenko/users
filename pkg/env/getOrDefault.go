package env

import (
	"os"
)

// Logger that we need for config loading
type Logger interface {
	Infof(template string, args ...interface{})
}

// GetOrDefault configuration and log it
func GetOrDefault(l Logger, env string, def string) string {
	v := os.Getenv(env)
	if v != "" {
		return v
	}
	l.Infof("using default value for %s", env)
	return def
}
