package config

import (
	"os"
	"time"

	"github.com/spf13/viper"
)

const (
	defaultHTTPPort                = "8080"
	defaultHTTPRWTimeout           = 10 * time.Second
	defaultHTTPMaxHeaderMegabytes  = 1
	defaultDatabaseRefreshInterval = 30 * time.Second
)

type (
	Config struct {
		HTTP HTTPConfig
	}

	HTTPConfig struct {
		Host               string
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegaBytes int           `mapstructure:"maxHeaderMegaBytes"`
	}
)

func InitConfig(configPath string) (*Config, error) {
	setDefaults()

	if err := parseConfigFile(configPath); err != nil {
		return nil, err
	}

	var cfg Config
	if err := unmarshal(&cfg); err != nil {
		return nil, err
	}

	setFromEnv(&cfg)

	return &cfg, nil
}

func unmarshal(cfg *Config) error {
	return viper.UnmarshalKey("http", &cfg.HTTP)
}

func setFromEnv(cfg *Config) {
	cfg.HTTP.Host = os.Getenv("HTTP_HOST")
}

func parseConfigFile(folder string) error {
	viper.AddConfigPath(folder)
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return viper.MergeInConfig()
}

func setDefaults() {
	viper.SetDefault("http.port", defaultHTTPPort)
	viper.SetDefault("http.maxHeaderMegaBytes", defaultHTTPMaxHeaderMegabytes)
	viper.SetDefault("http.readTimeout", defaultHTTPRWTimeout)
	viper.SetDefault("http.writeTimeout", defaultHTTPRWTimeout)
}
