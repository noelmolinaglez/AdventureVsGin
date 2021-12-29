package main

import (
	"net/http"

	"adventureVsModule/config"
	"adventureVsModule/pkg/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	r.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"data": "hello world"})
	})

	r.GET("/departments", controller.ListDepartments)

	r.Run()
}
