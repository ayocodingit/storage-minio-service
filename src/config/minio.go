package config

import "github.com/spf13/viper"

type MinioConfig struct {
	AccessKey string
	SecretKey string
	Endpoint  string
	Bucket    string
	Url       string
	Region    string
}

func LoadMinioConfig() *MinioConfig {
	return &MinioConfig{
		AccessKey: viper.GetString("MINIO_ACCESS_KEY"),
		SecretKey: viper.GetString("MINIO_SECRET_KEY"),
		Endpoint:  viper.GetString("MINIO_ENDPOINT"),
		Bucket:    viper.GetString("MINIO_BUCKET"),
		Url:       viper.GetString("MINIO_URL"),
		Region:    viper.GetString("MINIO_REGION"),
	}
}
