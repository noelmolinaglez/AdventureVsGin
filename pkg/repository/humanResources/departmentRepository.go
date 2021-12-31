package humanResources

import (
	"adventureVsModule/config"
	"adventureVsModule/pkg/db/common"
	"adventureVsModule/pkg/dto"
	"adventureVsModule/pkg/model"
	constants "adventureVsModule/pkg/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

const (
	departmentRepository = "DepartmentRepository"
	lDepartments         = "ListDepartments"
	addDepartment        = "addDepartments"
)

func ListDepartments(c *gin.Context, request dto.Request) {
	log.WithFields(log.Fields{constants.FileName: departmentRepository, constants.FunctionName: lDepartments}).Info(constants.StartFunction)
	db := config.DB
	var departments []model.Department

	common.ReadValues(c, request, db, departments)
	log.WithFields(log.Fields{constants.FileName: departmentRepository, constants.FunctionName: lDepartments}).Info(constants.EndFunction)
}

func CreateDepartment(c *gin.Context, department model.Department) {
	log.WithFields(log.Fields{constants.FileName: departmentRepository, constants.FunctionName: addDepartment}).Info(constants.StartFunction)
	db := config.DB

	department.ModifiedDate = time.Now().Format("2006-01-02 15:04:05")
	if err := db.Create(&department).Error; err != nil {
		log.WithFields(log.
			Fields{constants.Error: err.Error()}).
			Info(constants.EndException)

		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})
	} else {
		c.JSON(http.StatusCreated, gin.H{"data": department})
	}
	log.WithFields(log.Fields{constants.FileName: departmentRepository, constants.FunctionName: addDepartment}).Info(constants.EndFunction)
}

func UpdateDepartment(c *gin.Context, department model.Department) {
	log.WithFields(log.Fields{constants.FileName: departmentRepository, constants.FunctionName: addDepartment}).Info(constants.StartFunction)
	db := config.DB

	department.ModifiedDate = time.Now().Format("2006-01-02 15:04:05")
	if err := db.Updates(&department).Error; err != nil {
		log.WithFields(log.
			Fields{constants.Error: err.Error()}).
			Info(constants.EndException)

		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})
	} else {
		c.JSON(http.StatusCreated, gin.H{"data": department})
	}
	log.WithFields(log.Fields{constants.FileName: departmentRepository, constants.FunctionName: addDepartment}).Info(constants.EndFunction)
}
