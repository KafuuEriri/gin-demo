package main

import (
	"gin-demo/api"
	"gin-demo/middleware"

	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", api.Register)
	r.POST("/api/auth/login", api.Login)
	r.GET("/api/auth/info", middleware.AuthMiddleware(), api.Info)

	formRoutes := r.Group("/from")
	formApi := api.NewFormApi()
	formRoutes.POST("", formApi.Create)
	formRoutes.PUT(":id", formApi.Update)
	formRoutes.GET(":id", formApi.Show)
	formRoutes.DELETE(":id", formApi.Delete)

	return r
}
