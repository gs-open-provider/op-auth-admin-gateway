package config

import (
	"fmt"

	"github.com/spf13/viper"
)

/*
InitializeViper Function initializes viper to read config.yml file and environment variables in the application.
*/
func InitializeViper() {
	viper.SetConfigType("yml")

	// Set the file name of the configurations file
	viper.SetConfigName("private_config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Set the file name of the configurations file
	viper.SetConfigName("public_config")

	// Set the path to look for the configurations file
	viper.AddConfigPath(".")
	if err := viper.MergeInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

}
