package controller

import (
	"adventureVsModule/config"
	"adventureVsModule/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func List(c *gin.Context) {
	var departments []model.Department
	config.DB.Find(&departments)

	c.JSON(http.StatusOK, gin.H{"data": departments})

}
