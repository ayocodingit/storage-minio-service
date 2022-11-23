package config

import "github.com/spf13/viper"

type MinioConfig struct {
	AccessKey string `validate:"required"`
	SecretKey string `validate:"required"`
	Endpoint  string `validate:"required"`
	Bucket    string `validate:"required"`
	Url       string `validate:"required,url"`
	Region    string `validate:"required"`
	Ssl       bool   `validate:"required,boolean"`
}

func LoadMinioConfig() MinioConfig {
	return MinioConfig{
		AccessKey: viper.GetString("MINIO_ACCESS_KEY"),
		SecretKey: viper.GetString("MINIO_SECRET_KEY"),
		Endpoint:  viper.GetString("MINIO_ENDPOINT"),
		Bucket:    viper.GetString("MINIO_BUCKET"),
		Url:       viper.GetString("MINIO_URL"),
		Region:    viper.GetString("MINIO_REGION"),
		Ssl:       viper.GetBool("MINIO_SSL"),
	}
}
