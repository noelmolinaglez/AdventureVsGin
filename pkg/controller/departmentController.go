package controller

import (
	"adventureVsModule/config"
	"adventureVsModule/pkg/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ListDepartments(c *gin.Context) {

	db := config.DB
	var departments []model.Department

	db.Table("HumanResources.Department").Find(&departments)

	c.JSON(http.StatusOK, gin.H{"data": departments})

}
