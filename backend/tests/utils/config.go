package utils

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
	"social-network/pkg/config"
	"testing"
)

func NewTestConfiguration(_ *testing.T) *config.AppConfig {
	//
	//err := godotenv.Load(fmt.Sprintf("%s/.env"), cwd)
	//if err != nil {
	//	t.Fatal("Error loading .env file")
	//}

	viper.SetConfigType("env")
	viper.AddConfigPath(".env")
	viper.SetConfigFile(".env")

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		_, _ = fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}

	return config.NewAppConfig()
}
