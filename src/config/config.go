package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Port           string
	Secret         string
	IsPublicAccess bool
	Dst            string
	Minio          MinioConfig
}

func LoadConfig() Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	return Config{
		Port:           viper.GetString("APP_PORT"),
		Secret:         viper.GetString("APP_SECRET"),
		IsPublicAccess: viper.GetBool("IS_PUBLIC_ACCESS"),
		Dst:            "public",
		Minio:          LoadMinioConfig(),
	}
}
