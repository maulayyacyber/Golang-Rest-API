package main

import (
	"backend/controllers/postcontroller"
	"backend/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//import model
	models.ConnectDatabase()

	//routing
	router.GET("/api/posts", postcontroller.Index)
	router.POST("/api/posts", postcontroller.Store)
	router.GET("/api/posts/:id", postcontroller.Show)
	router.PUT("/api/posts/:id", postcontroller.Update)
	router.DELETE("/api/posts/:id", postcontroller.Destroy)
	router.Run(":3000")
}
