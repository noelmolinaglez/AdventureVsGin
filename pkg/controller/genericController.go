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

		humanResources.ListDepartments(c, request)
	}
	log.WithFields(log.Fields{constants.FileName: genericController, constants.FunctionName: genericList}).Info(constants.EndFunction)

}
