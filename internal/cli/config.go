package cli

import (
	"errors"
	"fmt"
	"os"

	"github.com/adrg/xdg"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const ConfigDirName = "advent-of-code"

type Config struct {
	// SessionCookie is used only in the event that a local input file is not present
	SessionCookie string `yaml:"sessionCookie"`
	Verbose       bool   `yaml:"verbose"`

	configDir string // not configurable
	cacheDir  string
}

func initConfig(cmd *cobra.Command) (config Config, err error) {
	config.configDir, err = xdg.ConfigFile(ConfigDirName)
	if err != nil {
		return config, fmt.Errorf("unable to determine config dir: %w", err)
	}

	config.cacheDir, err = xdg.CacheFile(ConfigDirName)
	if err != nil {
		return config, fmt.Errorf("unable to determine cache dir: %w", err)
	}

	if err := os.MkdirAll(config.configDir, 0750); err != nil {
		return config, fmt.Errorf("unable to create config dir: %w", err)
	}

	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(config.configDir)

	v.MustBindEnv("sessioncookie", "AOC_SESSION_COOKIE")

	v.AutomaticEnv()
	v.SetEnvPrefix("AOC")

	if err := v.BindPFlags(cmd.PersistentFlags()); err != nil {
		return config, fmt.Errorf("unable to bind to pflags: %w", err)
	}

	if err := v.ReadInConfig(); err != nil {
		if _, ok := errors.AsType[viper.ConfigFileNotFoundError](err); !ok {
			return config, fmt.Errorf("unable to read config: %w", err)
		}
	}

	if err := v.Unmarshal(&config); err != nil {
		return config, fmt.Errorf("unable to unmarshal yaml: %w", err)
	}

	return config, nil
}
