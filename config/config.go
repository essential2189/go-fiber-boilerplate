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
	AppHome  string
	Profile  string
	Port     string
	Test     bool
	LogLevel string
	Apis     string
)

func init() {
	AppHome = os.Getenv("APP_HOME")
	if len(AppHome) == 0 {
		AppHome = "."
	}

	Profile = os.Getenv("wavve_env")
	if len(Profile) == 0 {
		Profile = "local"
	}

	Port = os.Getenv("wavve_port")
	if len(Port) == 0 {
		Port = "8080"
	}

	Test = os.Getenv("TEST") == "true"

	LogLevel = os.Getenv("LOG_LEVEL")

	if Profile == "prd" {
		Apis = "http://apis.local.wavve.com/"
	} else if Profile == "qa" {
		Apis = "http://qa-apis.wavve.com/"
	} else {
		//Apis = "http://dev-apis.wavve.com/"
		Apis = "http://apis.local.wavve.com/"
	}
}

type Vars struct {
	AppHome  string
	Profile  string
	Port     string
	Test     bool
	LogLevel string
	Apis     string
}

func NewVars() *Vars {
	return &Vars{
		AppHome:  AppHome,
		Profile:  Profile,
		Port:     Port,
		Test:     Test,
		LogLevel: LogLevel,
		Apis:     Apis,
	}
}
