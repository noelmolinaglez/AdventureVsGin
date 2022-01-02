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

	r.GET("/generic", controller.List)
	r.GET("/departments", controller.ListDepartments)
	r.POST("/departments/add", controller.CreateDepartment)
	r.PUT("/departments", controller.UpdateDepartment)
	r.DELETE("/departments", controller.DeleteDepartment)

	r.Run()
}
