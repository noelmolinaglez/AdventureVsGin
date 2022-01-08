package main

import (
	"net/http"

	"adventureVsModule/config"
	"adventureVsModule/pkg/controller"
	"github.com/gin-gonic/gin"
	"github.com/noelmolinaglez/simplePdf/pdf"
)

func main() {
	r := gin.Default()

	config.ConnectDatabase()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})

	r.POST("/generic", controller.Crud)

	pdf.SamplePdf()

	r.Run()
}
