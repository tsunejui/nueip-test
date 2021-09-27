package conf

import (
	"os"
)

type Config struct {
	ListenPort string
	DBHost     string
	DBPort     string
	DBName     string
	DBUsername string
	DBPassword string
}

const (
	envListenPort string = "LISTEN_PORT"
	envDBName     string = "DB_NAME"
	envDBHost     string = "DB_HOST"
	envDBPort     string = "DB_PORT"
	envDBUsername string = "DB_USERNAME"
	envDBPassword string = "DB_PASSWORD"
)

var config = &Config{}

func init() {
	config.ListenPort = os.Getenv(envListenPort)
	config.DBName = os.Getenv(envDBName)
	config.DBHost = os.Getenv(envDBHost)
	config.DBPort = os.Getenv(envDBPort)
	config.DBUsername = os.Getenv(envDBUsername)
	config.DBPassword = os.Getenv(envDBPassword)
}

func GetConfig() *Config {
	return config
}
