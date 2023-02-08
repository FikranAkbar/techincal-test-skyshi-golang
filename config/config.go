package config

import (
	"fmt"
	"github.com/spf13/viper"
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
		MySqlHost:     viper.GetString("MYSQL_HOST"),
		MySqlPort:     viper.GetString("MYSQL_PORT"),
		MySqlUser:     viper.GetString("MYSQL_USER"),
		MySqlPassword: viper.GetString("MYSQL_PASSWORD"),
		MySqlDBName:   viper.GetString("MYSQL_DBNAME"),
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
