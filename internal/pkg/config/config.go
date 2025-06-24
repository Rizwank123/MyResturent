package config

import (
	"fmt"

	"github.com/rizwank123/myResturent/internal/domain"
	"github.com/spf13/viper"
)

const (
	SourceKey = "CONFIG_SOURCE"
	SourceEnv = "ENVOIRMENT"
)

type Options struct {
	ConfigSource string
	ConfigFile   string
}

type ResturantConfig struct {
	AppName string `mapstructure:"APP_NAME"`
	AppEnv  string `mapstructure:"APP_ENV"`
	AppPort int    `mapstructure:"APP_PORT"`

	DatabaseHost     string `mapstructure:"DB_HOST"`
	DatabasePort     string `mapstructure:"DB_PORT"`
	DatabaseUsername string `mapstructure:"DB_USERNAME"`
	DatabasePassword string `mapstructure:"DB_PASSWORD"`
	DatabaseName     string `mapstructure:"DB_DATABASE_NAME"`

	SwaggerHostUrl    string `mapstructure:"SWAGGER_HOST_URL"`
	SwaggerHostScheme string `mapstructure:"SWAGGER_HOST_SCHEME"`
	SwaggerUsername   string `mapstructure:"SWAGGER_USERNAME"`
	SwaggerPassword   string `mapstructure:"SWAGGER_PASSWORD"`

	AuthSecret       string `mapstructure:"AUTH_SECRET"`
	AuthExpiryPeriod int    `mapstructure:"AUTH_EXPIRY_PERIOD"`
}

func NewResturantConfig(options Options) (ResturantConfig, error) {
	if options.ConfigSource == SourceEnv {
		return newFromEnvironment(options)
	}
	return ResturantConfig{}, domain.SystemError{Message: fmt.Sprintf("invalid config source: %s", options.ConfigSource)}

}

func newFromEnvironment(opts Options) (cfg ResturantConfig, err error) {
	viper.SetConfigFile(opts.ConfigFile)
	viper.SetConfigType("env")

	err = viper.ReadInConfig()
	if err != nil {
		return cfg, fmt.Errorf("failed to read config file: %v", err)
	}

	cfg = ResturantConfig{}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		return cfg, fmt.Errorf("failed to load configuration: %v", err)
	}

	return cfg, nil
}
