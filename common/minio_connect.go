package common

import (
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/spf13/viper"
)

var OOS *minio.Client

func InitOOS() *minio.Client {
	// 初始化minio的连接
	endpoint := viper.GetString("oos_datasource.endpoint")
	accessKeyID := viper.GetString("oos_datasource.accessKeyID")
	secretAccessKey := viper.GetString("oos_datasource.secretAccessKey")
	useSSL := viper.GetBool("oos_datasource.useSSl")

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		panic("fail to connect OOS, err: " + err.Error())
	}
	OOS = minioClient
	return minioClient

}

func GetOOS() *minio.Client {
	return OOS
}
