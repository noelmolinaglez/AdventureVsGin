package controller

import (
	"adventureVsModule/pkg/dto"
	"adventureVsModule/pkg/model"
	"adventureVsModule/pkg/repository"
	constants "adventureVsModule/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

const (
	genericController = "genericController"
	genericList       = "genericList"
)

func Crud(c *gin.Context) {
	log.WithFields(log.Fields{constants.FileName: genericController, constants.FunctionName: genericList}).Info(constants.StartFunction)
	var request dto.Request

	if err := c.ShouldBindJSON(&request); err != nil {

		log.WithFields(log.Fields{constants.Error: err.Error()}).Info(constants.EndException)
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})

	} else {
		actionString := fmt.Sprintf("%sAction", request.Type)
		queryString := fmt.Sprintf("%sQuery", request.Action)

		if request.Action == "List" {
			ListQuery(c, request, actionString, queryString)
		} else {

			models := FillModels()

			actions := map[string]func(c *gin.Context, model interface{}, data interface{}, action string, query string){
				"Create": repository.Create,
				"Update": repository.Update,
				"Delete": repository.Delete,
			}

			myInstance, ok := models[request.Type]
			if !ok {

			}

			//myInstance.ModifiedDate = time.Now().Format("2006-01-02 15:04:05")
			actions[request.Action](c, myInstance, request.Data, actionString, queryString)
		}
	}
	log.WithFields(log.Fields{constants.FileName: genericController, constants.FunctionName: genericList}).Info(constants.EndFunction)

}

func FillModels() map[string]interface{} {
	var department model.Department
	var employee model.Employee
	var employeeDepartment model.EmployeeDepartmentHistory
	var shift model.Shift
	var payment model.EmployeePayHistory

	models := map[string]interface{}{
		"Department": department,
		"Employee":   employee,
		"History":    employeeDepartment,
		"Shift":      shift,
		"Payment":    payment,
	}
	return models
}

func ListQuery(c *gin.Context, request dto.Request, actionString string, queryString string) {
	var departments []model.Department
	var employees []model.Employee
	var employeeDepartments []model.EmployeeDepartmentHistory
	var shifts []model.Shift
	var payments []model.EmployeePayHistory

	results := map[string]interface{}{
		"Department": departments,
		"Employee":   employees,
		"History":    employeeDepartments,
		"Shift":      shifts,
		"Payment":    payments,
	}

	actions := map[string]func(c *gin.Context, request dto.Request, result interface{}, action string, query string){
		"List": repository.List,
	}

	actions[request.Action](c, request, results[request.Type], actionString, queryString)
}
