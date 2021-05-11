package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Table driven test
func TestLoadConfigFromEnv(t *testing.T) {
	cases := map[string]struct {
		env            map[string]string
		expectedConfig Config
	}{
		"usual case": {
			env: map[string]string{"DB_USER": "TEST-TEST-1"},
			expectedConfig: Config{
				Application: Application{
					Port: ":8080",
				},
				Database: Database{
					User:     "TEST-TEST-1",
					Password: "test_pass",
					Host:     "postgres",
					DBName:   "test_db",
				},
			},
		},
		"more custom variables": {
			env: map[string]string{"DB_USER": "TEST-TEST-2", "PORT": ":TEST-TEST-3", "DB_HOST": "TEST-TEST-4"},
			expectedConfig: Config{
				Application: Application{
					Port: ":TEST-TEST-3",
				},
				Database: Database{
					User:     "TEST-TEST-2",
					Password: "test_pass",
					Host:     "TEST-TEST-4",
					DBName:   "test_db",
				},
			},
		},
	}
	for name, c := range cases {
		t.Run(name, func(t *testing.T) {
			for key, value := range c.env {
				err := os.Setenv(key, value)
				require.NoError(t, err)
			}
			actual := LoadConfigFromEnv(&nopLogger{})
			assert.Equal(t, c.expectedConfig, actual)
		})
	}
}

type nopLogger struct{}

func (n *nopLogger) Infof(template string, args ...interface{}) {}
