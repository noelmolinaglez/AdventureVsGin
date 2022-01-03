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

func List(c *gin.Context) {
	log.WithFields(log.Fields{constants.FileName: genericController, constants.FunctionName: genericList}).Info(constants.StartFunction)
	var request dto.Request

	if err := c.ShouldBindJSON(&request); err != nil {

		log.WithFields(log.Fields{constants.Error: err.Error()}).Info(constants.EndException)
		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})

	} else {
		actionString := fmt.Sprintf("%sAction", request.Type)
		queryString := fmt.Sprintf("%sQuery", request.Action)

		var departments []model.Department

		results := map[string]interface{}{
			"Department": departments,
		}

		repository.List(c, request, results[request.Type], actionString, queryString)
	}

	log.WithFields(log.Fields{constants.FileName: genericController, constants.FunctionName: genericList}).Info(constants.EndFunction)
}
