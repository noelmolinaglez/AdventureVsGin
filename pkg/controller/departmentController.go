package controller

import (
	"adventureVsModule/pkg/Repository/humanResources"
	constants "adventureVsModule/pkg/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

const (
	departmentController = "DepartmentController"
	lDepartments         = "ListDepartments"
)

func ListDepartments(c *gin.Context) {
	log.WithFields(log.Fields{constants.FileName: departmentController, constants.FunctionName: lDepartments}).Info(constants.StartFunction)

	log.WithFields(log.Fields{constants.FileName: departmentController, constants.FunctionName: lDepartments}).Info(constants.EndFunction)
	humanResources.ListDepartments(c)

}
