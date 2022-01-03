package config

import "github.com/spf13/viper"

type appConfig struct {
	AppEnv       string `mapstructure:"APP_ENV"`
	AppServePort string `mapstructure:"APP_PORT"`

	MySqlHost     string `mapstructure:"MYSQL_HOST"`
	MySqlDatabase string `mapstructure:"MYSQL_DATABASE"`
	MySqlUser     string `mapstructure:"MYSQL_USER"`
	MySqlPassword string `mapstructure:"MYSQL_PASSWORD"`
	MySqlPort     string `mapstructure:"MYSQL_PORT"`

	RedisHost string `mapstructure:"REDIS_HOST"`
	RedisPort string `mapstructure:"REDIS_PORT"`
}

func NewAppConfig() *appConfig {
	return &appConfig{
		AppEnv:        viper.GetString("APP_ENV"),
		AppServePort:  viper.GetString("APP_PORT"),
		MySqlHost:     viper.GetString("MYSQL_HOST"),
		MySqlDatabase: viper.GetString("MYSQL_DATABASE"),
		MySqlUser:     viper.GetString("MYSQL_USER"),
		MySqlPassword: viper.GetString("MYSQL_PASSWORD"),
		MySqlPort:     viper.GetString("MYSQL_PORT"),
		RedisHost:     viper.GetString("REDIS_HOST"),
		RedisPort:     viper.GetString("REDIS_PORT"),
	}
}
