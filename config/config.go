package config

import (
	"fmt"
	"os"
)

type Config struct {
	MySqlHost     string `mapstructure:"mysql_host"`
	MySqlPort     string `mapstructure:"mysql_port"`
	MySqlUser     string `mapstructure:"mysql_user"`
	MySqlPassword string `mapstructure:"mysql_password"`
	MySqlDBName   string `mapstructure:"mysql_db_name"`
}

func LoadConfigFromEnv() *Config {
	config := Config{
		MySqlHost:     os.Getenv("MYSQL_HOST"),
		MySqlPort:     os.Getenv("MYSQL_PORT"),
		MySqlUser:     os.Getenv("MYSQL_USER"),
		MySqlPassword: os.Getenv("MYSQL_PASSWORD"),
		MySqlDBName:   os.Getenv("MYSQL_DBNAME"),
	}

	return &config
}

func (config *Config) CreateDBURL() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.MySqlUser,
		config.MySqlPassword,
		config.MySqlHost,
		config.MySqlPort,
		config.MySqlDBName,
	)
}
