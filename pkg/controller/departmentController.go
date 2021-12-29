package controller

import (
	"adventureVsModule/pkg/dto"
	"adventureVsModule/pkg/repository/humanResources"
	constants "adventureVsModule/pkg/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	departmentController = "DepartmentController"
	lDepartments         = "ListDepartments"
)

func ListDepartments(c *gin.Context) {
	log.WithFields(log.Fields{constants.FileName: departmentController, constants.FunctionName: lDepartments}).Info(constants.StartFunction)
	var request dto.Request
	if err := c.ShouldBind(&request); err != nil {

		log.WithFields(log.Fields{constants.Error: err.Error()}).Info(constants.EndException)
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})

	} else {

		humanResources.ListDepartments(c, request)
	}
	log.WithFields(log.Fields{constants.FileName: departmentController, constants.FunctionName: lDepartments}).Info(constants.EndFunction)

}
