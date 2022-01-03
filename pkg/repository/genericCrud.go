package repository

import (
	"adventureVsModule/config"
	"adventureVsModule/pkg/db/common"
	"adventureVsModule/pkg/dto"
	"adventureVsModule/pkg/model"
	constants "adventureVsModule/pkg/utils"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"time"
)

func List(c *gin.Context, request dto.Request, result interface{}, action string, query string) {
	log.WithFields(log.Fields{constants.FileName: action, constants.FunctionName: query}).Info(constants.StartFunction)
	db := config.DB

	common.ReadValues(c, request, db, result)

	log.WithFields(log.Fields{constants.FileName: action, constants.FunctionName: query}).Info(constants.EndFunction)
}

func Create(c *gin.Context, model interface{}, action string, query string) {
	log.WithFields(log.Fields{constants.FileName: action, constants.FunctionName: query}).Info(constants.StartFunction)
	db := config.DB

	model.ModifiedDate = time.Now().Format("2006-01-02 15:04:05")
	common.Create(c, &model, db)

	log.WithFields(log.Fields{constants.FileName: action, constants.FunctionName: query}).Info(constants.EndFunction)
}
