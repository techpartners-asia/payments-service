package storage

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

func InitStorage() *minio.Client {
	minioClient, err := minio.New(viper.GetString("STORAGE_HOST"), &minio.Options{
		Creds:  credentials.NewStaticV4(viper.GetString("STORAGE_USER"), viper.GetString("STORAGE_PASSWORD"), ""),
		Secure: viper.GetBool("STORAGE_USESSL"),
	})
	if err != nil {
		panic(err.Error())
	}
	return minioClient
}
