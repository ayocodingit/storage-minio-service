package config

import (
	"log"

	"github.com/jabardigitalservice/utilities-go/validator"
	"github.com/spf13/viper"
)

type Config struct {
	Port   string      `validate:"required"`
	Secret string      `validate:"required"`
	Dst    string      `validate:"required"`
	Minio  MinioConfig `validate:"required"`
}

func LoadConfig(validator validator.Validator) Config {
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		log.Fatal(err)
	}

	config := Config{
		Port:   viper.GetString("APP_PORT"),
		Secret: viper.GetString("APP_SECRET"),
		Dst:    "tmp",
		Minio:  LoadMinioConfig(),
	}

	if err := validator.Validation(config); err != nil {
		log.Fatal(err)
	}

	return config
}
