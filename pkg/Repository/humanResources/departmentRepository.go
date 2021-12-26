package humanResources

import (
	"adventureVsModule/config"
	"adventureVsModule/pkg/model"
	constants "adventureVsModule/pkg/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	departmentRepository = "DepartmentRepository"
	lDepartments         = "ListDepartments"
)

func ListDepartments(c *gin.Context) {
	log.WithFields(log.Fields{constants.FileName: departmentRepository, constants.FunctionName: lDepartments}).Info(constants.StartFunction)
	db := config.DB
	var departments []model.Department

	db.Table("HumanResources.Department").Find(&departments)

	log.WithFields(log.Fields{constants.FileName: departmentRepository, constants.FunctionName: lDepartments}).Info(constants.EndFunction)
	c.JSON(http.StatusOK, gin.H{"data": departments})
}
