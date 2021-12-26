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

	if err := db.Table("HumanResources.Department").Find(&departments).Error; err != nil {
		log.WithFields(log.
			Fields{constants.Error: err.Error()}).
			Info(constants.EndException)

		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": departments})
	}
	log.WithFields(log.Fields{constants.FileName: departmentRepository, constants.FunctionName: lDepartments}).Info(constants.EndFunction)
}
