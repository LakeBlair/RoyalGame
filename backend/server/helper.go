package server

import (
	"github.com/spf13/viper"
)

func Get(relativePath string) *viper.Viper {
	config := viper.New()
	config.SetConfigFile(relativePath)
	config.AutomaticEnv()
	err := config.ReadInConfig()
	if err != nil {
		panic(err)
	}
	return config
}

func ContainsString(l []string, s string) bool {
	for _, a := range l {
		if a == s {
			return true
		}
	}
	return false
}
