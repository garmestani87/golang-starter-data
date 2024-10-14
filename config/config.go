package config

import (
	"errors"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	Application
	Server
	Cache
	Data
}

type Application struct {
	Name string
}

type Server struct {
	Port string
}

type Cache struct {
	Redis
}

type Redis struct {
	Uri                string
	Password           string
	Db                 int
	ReadTimeout        int
	WriteTimeout       int
	PoolSize           int
	PoolTimeout        int
	IdleTimeout        int
	IdleCheckFrequency int
}

type Data struct {
	Postgres
}

type Postgres struct {
	Host            string
	Port            int
	User            string
	Password        string
	DbName          string
	SslMode         string
	MaxIdleConns    int
	MaxOpenConns    int
	ConnMaxLifetime int
}

func GetConfig() *Config {

	configPath := getConfigPath(os.Getenv("APP_ENV"))
	v, err := loadConfig(configPath, "yml")
	if err != nil {
		log.Fatalf("error in loading file : %v", err)
	}

	cfg, err := parseConfig(v)

	if err != nil {
		log.Fatalf("Error in parse config %v", err)
	}

	return cfg

}

func parseConfig(v *viper.Viper) (cfg *Config, err error) {
	err = v.Unmarshal(&cfg)
	if err != nil {
		log.Printf("unable to parse config : %v", err)
		return nil, err
	}
	return cfg, nil
}

func loadConfig(configName string, configType string) (*viper.Viper, error) {
	v := viper.New()
	v.SetConfigType(configType)
	v.SetConfigName(configName)
	v.AddConfigPath(".")
	v.AutomaticEnv()

	err := v.ReadInConfig()
	if err != nil {
		log.Printf("unable to read config file : %v \n", err)
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return nil, errors.New("config file not found")
		}
		return nil, err
	}
	return v, nil

}

func getConfigPath(env string) (path string) {
	switch env {
	case "dev":
		path = "config/application-dev"
	case "test":
		path = "config/application-test"
	case "prod":
		path = "config/application-prod"
	default:
		path = ""
	}
	return
}
