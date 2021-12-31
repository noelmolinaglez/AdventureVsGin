package controller

import (
	"adventureVsModule/pkg/dto"
	"adventureVsModule/pkg/model"
	"adventureVsModule/pkg/repository/humanResources"
	constants "adventureVsModule/pkg/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	departmentController = "DepartmentController"
	lDepartments         = "ListDepartments"
	addDepartment        = "CreateDepartment"
	updateDepartment     = "UpdateDepartment"
)

func ListDepartments(c *gin.Context) {
	log.WithFields(log.Fields{constants.FileName: departmentController, constants.FunctionName: lDepartments}).Info(constants.StartFunction)
	var request dto.Request
	if err := c.ShouldBindJSON(&request); err != nil {

		log.WithFields(log.Fields{constants.Error: err.Error()}).Info(constants.EndException)
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})

	} else {

		humanResources.ListDepartments(c, request)
	}
	log.WithFields(log.Fields{constants.FileName: departmentController, constants.FunctionName: lDepartments}).Info(constants.EndFunction)

}

func CreateDepartment(c *gin.Context) {
	log.WithFields(log.Fields{constants.FileName: departmentController, constants.FunctionName: addDepartment}).Info(constants.StartFunction)
	var department model.Department
	if err := c.ShouldBindJSON(&department); err != nil {

		log.WithFields(log.Fields{constants.Error: err.Error()}).Info(constants.EndException)
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})

	} else {

		humanResources.CreateDepartment(c, department)
	}
	log.WithFields(log.Fields{constants.FileName: departmentController, constants.FunctionName: addDepartment}).Info(constants.EndFunction)

}

func UdpateDepartment(c *gin.Context) {
	log.WithFields(log.Fields{constants.FileName: departmentController, constants.FunctionName: updateDepartment}).Info(constants.StartFunction)
	var department model.Department
	if err := c.ShouldBindJSON(&department); err != nil {

		log.WithFields(log.Fields{constants.Error: err.Error()}).Info(constants.EndException)
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})

	} else {

		humanResources.UpdateDepartment(c, department)
	}
	log.WithFields(log.Fields{constants.FileName: departmentController, constants.FunctionName: updateDepartment}).Info(constants.EndFunction)

}
