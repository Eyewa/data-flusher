package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	vars := map[string]string{
		"ENV":                    "debug",
		"LOG_LEVEL":              "debug",
		"ALGOLIA_APPLICATION_ID": "debug",
		"ALGOLIA_API_KEY":        "debug",
		"DB_DRIVER":              "debug",
		"DB_HOST":                "debug",
		"DB_DATABASE":            "debug",
		"DB_PORT":                "debug",
		"DB_USER":                "debug",
		"DB_PASSWORD":            "debug",
	}
	for e, v := range vars {
		os.Setenv(e, v)
	}

	func(t *testing.T) {
		assert.Nil(t, Init())
	}(new(testing.T))

	exitVal := m.Run()

	for _, v := range envVars {
		os.Unsetenv(v)
	}

	os.Exit(exitVal)
}

func TestConfig(t *testing.T) {
	assert.NotZero(t, Config)
	assert.Equal(t, "debug", Config.Env)
	assert.Equal(t, "debug", Config.LogLevel)
	assert.Equal(t, "debug", Config.Algolia.ApplicationID)
	assert.Equal(t, "debug", Config.Algolia.APIKey)
	assert.Equal(t, "debug", Config.DB.Driver)
	assert.Equal(t, "debug", Config.DB.Database.Host)
	assert.Equal(t, "debug", Config.DB.Database.Name)
	assert.Equal(t, "debug", Config.DB.Database.Port)
	assert.Equal(t, "debug", Config.DB.Database.User)
	assert.Equal(t, "debug", Config.DB.Database.Password)
}
