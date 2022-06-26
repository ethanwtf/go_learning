package main

import (
	"gin_vue_project/common"
	"gin_vue_project/routes"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	InitConfig()
	common.InitDB()

	r := gin.Default()
	r = routes.CollectRoute(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	} else {
		panic(r.Run())
	}

}
