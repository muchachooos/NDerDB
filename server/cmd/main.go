package main

import (
	"NDerDB/server/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.New()

	router.POST("/save_string", handler.SaveStringHandler)
	router.GET("/get_string", handler.GetStringHandler)

	err := router.Run(":8080")
	if err != nil {
		panic(err)
	}
}
