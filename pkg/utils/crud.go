package utils

import (
	"adventureVsModule/pkg/dto"

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
			Fields{Error: err.Error()}).
			Info(EndException)

		c.JSON(http.StatusInternalServerError, gin.H{"data": nil})
	} else {
		c.JSON(http.StatusOK, gin.H{"data": results})
	}
}
