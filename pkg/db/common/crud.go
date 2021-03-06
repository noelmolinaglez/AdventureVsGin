package common

import (
	"adventureVsModule/pkg/dto"
	"adventureVsModule/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
)

func ReadValues(c *gin.Context, request dto.Request, db *gorm.DB, results interface{}) {
	orderBy := fmt.Sprintf("%s %s", request.Field, request.Dir)

	query := db.Order(orderBy)
	query.Offset(request.Start).Limit(request.Limit)

	if err := query.Find(&results).Error; err != nil {
		log.WithFields(log.
			Fields{utils.Error: err.Error()}).
			Info(utils.EndException)

		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": results})
	}
}

func Create(c *gin.Context, model interface{}, data interface{}, db *gorm.DB) {

	if err := db.Model(&model).Create(data).Error; err != nil {
		log.WithFields(log.
			Fields{utils.Error: err.Error()}).
			Info(utils.EndException)

		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})
	} else {
		c.JSON(http.StatusCreated, gin.H{"data": data})
	}
}

func Update(c *gin.Context, model interface{}, data interface{}, db *gorm.DB) {
	if err := db.Model(&model).Updates(data).Error; err != nil {
		log.WithFields(log.
			Fields{utils.Error: err.Error()}).
			Info(utils.EndException)

		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})
	} else {
		c.JSON(http.StatusCreated, gin.H{"data": data})
	}
}

func Delete(c *gin.Context, model interface{}, data interface{}, db *gorm.DB) {
	if err := db.Model(&model).Delete(data).Error; err != nil {
		log.WithFields(log.
			Fields{utils.Error: err.Error()}).
			Info(utils.EndException)

		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})
	}
}
