package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2/log"
	"github.com/spf13/viper"
)

func NewConfig(vars *Vars) (*Config, error) {
	var config *Config
	path := filepath.Join(fmt.Sprintf("config/configs/config.%s.yml", vars.Profile))

	viper.SetConfigType("yaml")
	viper.SetConfigFile(path)

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Cannot read %s", path)
	}

	if err := viper.Unmarshal(&config); err != nil {
		return nil, err
	}

	return config, nil
}

var (
	Profile  string
	Port     string
	Test     bool
	LogLevel string
	Apis     string
)

func init() {
	Profile = os.Getenv("env")
	if len(Profile) == 0 {
		Profile = "dev"
	}
}

type Vars struct {
	Profile string
}

func NewVars() *Vars {
	return &Vars{
		Profile: Profile,
	}
}
