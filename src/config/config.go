package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port   string
	Secret string
	Dst    string
	Minio  MinioConfig
}

func LoadConfig() Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	return Config{
		Port:   viper.GetString("APP_PORT"),
		Secret: viper.GetString("APP_SECRET"),
		Dst:    "tmp",
		Minio:  LoadMinioConfig(),
	}
}
