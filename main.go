package main

import (
	"gin_vue_project/common"
	"gin_vue_project/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()

	r := gin.Default()
	r = routes.CollectRoute(r)

	panic(r.Run())

}
