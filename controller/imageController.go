package controller

import (
	"context"
	"fmt"
	"gin_vue_project/common"
	"gin_vue_project/response"
	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	uuid "github.com/satori/go.uuid"
	"github.com/spf13/viper"
	"net/http"
	"strings"
)

func Uploadimage(c *gin.Context) {
	// 读取文件
	file, err := c.FormFile("file")
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "上传图片出错")
		return
	}
	src, err := file.Open()
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "上传图片出错")
		return
	}

	//创建图片uuid
	u2 := uuid.NewV4()

	//检查文件格式
	file_split := strings.Split(file.Filename, ".")
	file_format := file_split[len(file_split)-1]
	fmt.Printf("文件类型: %s\n", file_format)
	if !(file_format == "png" || file_format == "jpg" || file_format == "bmp") {
		response.Response(c, http.StatusBadRequest, 400, nil, "图片格式不正确")
		return
	}

	//检查文件大小
	if (file.Size / 1024) > 500 {
		response.Response(c, http.StatusBadRequest, 400, nil, "图片大小超过限制")
		return
	}

	bucketName := "first"
	objectName := u2.String() + "." + file_format

	ctx := context.Background()
	OOS := common.GetOOS()
	endpoint := viper.GetString("oos_datasource.endpoint")

	n, err := OOS.PutObject(ctx, bucketName, objectName, src, file.Size, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		response.Response(c, http.StatusInternalServerError, 500, nil, "上传图片出错")
		return
	}
	//返回url
	url := "http://" + endpoint + "/" + bucketName + "/" + objectName
	response.Success(c, gin.H{"url": url}, "上传图片成功")
	fmt.Println("image_upload_message", n)
	return
}
