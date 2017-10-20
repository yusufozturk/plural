package config

import (
	"github.com/spf13/viper"
)

func init() {
	// Configuration file settings using key-value
	viper.SetConfigName("plural")
	viper.AddConfigPath("/opt/plural/conf")
	viper.AddConfigPath("/etc/plural/conf")
	viper.AddConfigPath("conf")
	viper.AddConfigPath(".")
	viper.ReadInConfig()

	// Default settings if no config file is supplied
	viper.SetDefault("elastic_host", "localhost")
	viper.SetDefault("elastic_port", "9200")
	viper.SetDefault("environment", "dev")
	viper.SetDefault("interval", "300")
	viper.SetDefault("username", "")
	viper.SetDefault("password", "")
	viper.SetDefault("overwrite", "")
	viper.SetDefault("secure", "false")
}

func ConfigStr(key string) string {
	return viper.GetString(key)
}

func ConfigInt(key string) int {
	return viper.GetInt(key)
}

func ConfigBool(key string) bool {
	return viper.GetBool(key)
}
