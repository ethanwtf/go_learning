package routes

import (
	"gin_vue_project/controller"
	"gin_vue_project/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), controller.Info)
	r.POST("/api/img_upload", middleware.AuthMiddleware(), controller.Uploadimage)
	//
	//categoryRoutes := r.Group("/categories")
	//categoryController := controller.NewCategoryController()
	//categoryRoutes.POST("", categoryController.Create)
	//categoryRoutes.PUT("/:id", categoryController.Update) //替换
	//categoryRoutes.GET("/:id", categoryController.Show)
	//categoryRoutes.DELETE("/:id", categoryController.Delete)
	//
	//postRoutes := r.Group("/posts")
	//postRoutes.Use(middleware.AuthMiddleware())
	//postController := controller.NewPostController()
	//postRoutes.POST("", postController.Create)
	//postRoutes.PUT("/:id", postController.Update) //替换
	//postRoutes.GET("/:id", postController.Show)
	//postRoutes.DELETE("/:id", postController.Delete)
	//postRoutes.POST("/page/list", postController.PageList)

	return r
}
