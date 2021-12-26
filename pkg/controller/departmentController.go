package controller

import (
	"adventureVsModule/pkg/Repository/humanResources"
	"github.com/gin-gonic/gin"
)

func ListDepartments(c *gin.Context) {

	humanResources.ListDepartments(c)

}
