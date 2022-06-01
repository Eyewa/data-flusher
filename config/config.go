package config

import (
	"os"
	"strings"

	"github.com/eyewa/eyewa-go-lib/db"
	"github.com/eyewa/eyewa-go-lib/log"
	"github.com/ory/viper"
	"go.uber.org/zap"
)

var (
	Config  Configuration
	envVars = []string{
		"ENV",
		"LOG_LEVEL",
		"ALGOLIA_APPLICATION_ID",
		"ALGOLIA_API_KEY",
		"DB_DRIVER",
		"DB_HOST",
		"DB_DATABASE",
		"DB_PORT",
		"DB_USER",
		"DB_PASSWORD",
	}
)

type Configuration struct {
	Env      string
	LogLevel string    `mapstructure:"log_level"`
	DB       db.Config `mapstructure:",squash"`
	Algolia  struct {
		ApplicationID string `mapstructure:"algolia_application_id"`
		APIKey        string `mapstructure:"algolia_api_key"`
	} `mapstructure:",squash"`
}

// Init attempts to read env vars from the environ as first priority.
// If none is set, falls back to reading vars from a .env file (if exists).
func Init() error {
	// attempt to read vars from env
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	// only look for/use .env file if no env is set
	viper.SetConfigName(".env." + os.Getenv("ENV"))
	viper.AddConfigPath("../")
	viper.AddConfigPath("../../") // for pact/provider testing package
	viper.AddConfigPath(".")
	viper.SetConfigType("env")

	err := viper.ReadInConfig()
	if err != nil {
		return err
	}

	// Bind viper keys to env vars
	for _, v := range envVars {
		if err := viper.BindEnv(v); err != nil {
			return err
		}
	}

	if err := viper.Unmarshal(&Config); err != nil {
		return err
	}

	// ensure vars are added to env so golib can pick them up
	// for cases where envs are only found in .env
	for _, v := range envVars {
		os.Setenv(v, viper.GetString(v))
	}

	log.SetLogLevel()
	log.Info("Environment detected.", zap.String("env", Config.Env))

	return nil
}
