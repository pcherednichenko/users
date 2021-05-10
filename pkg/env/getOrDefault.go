package env

import (
	"os"
)

type Logger interface {
	Infof(template string, args ...interface{})
}

//
func GetOrDefault(l Logger, env string, def string) string {
	v := os.Getenv(env)
	if v != "" {
		return v
	}
	l.Infof("using default value for %s", env)
	return def
}