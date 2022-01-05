package repository

import (
	"adventureVsModule/config"
	"adventureVsModule/pkg/db/common"
	"adventureVsModule/pkg/dto"
	constants "adventureVsModule/pkg/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func List(c *gin.Context, request dto.Request, result interface{}, action string, query string) {
	log.WithFields(log.Fields{constants.FileName: action, constants.FunctionName: query}).Info(constants.StartFunction)
	db := config.DB

	common.ReadValues(c, request, db, result)

	log.WithFields(log.Fields{constants.FileName: action, constants.FunctionName: query}).Info(constants.EndFunction)
}

func Create(c *gin.Context, model interface{}, data interface{}, action string, query string) {
	log.WithFields(log.Fields{constants.FileName: action, constants.FunctionName: query}).Info(constants.StartFunction)
	db := config.DB

	common.Create(c, model, data, db)

	log.WithFields(log.Fields{constants.FileName: action, constants.FunctionName: query}).Info(constants.EndFunction)
}

func Update(c *gin.Context, model interface{}, data interface{}, action string, query string) {
	log.WithFields(log.Fields{constants.FileName: action, constants.FunctionName: query}).Info(constants.StartFunction)
	db := config.DB

	common.Update(c, model, data, db)

	log.WithFields(log.Fields{constants.FileName: action, constants.FunctionName: query}).Info(constants.EndFunction)
}

func Delete(c *gin.Context, model interface{}, data interface{}, action string, query string) {
	log.WithFields(log.Fields{constants.FileName: action, constants.FunctionName: query}).Info(constants.StartFunction)
	db := config.DB

	common.Delete(c, model, data, db)

	log.WithFields(log.Fields{constants.FileName: action, constants.FunctionName: query}).Info(constants.EndFunction)
}
