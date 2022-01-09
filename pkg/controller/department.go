package controller

import (
	"adventureVsModule/config"
	"adventureVsModule/pkg/model"
	"github.com/gin-gonic/gin"
	"github.com/noelmolinaglez/simplePdf/pkg/dto"
	"github.com/noelmolinaglez/simplePdf/pkg/helpers"
	"github.com/noelmolinaglez/simplePdf/pkg/interfaces"
	"github.com/noelmolinaglez/simplePdf/pkg/templates"
)

func ExportPDf(c *gin.Context) {

	db := config.DB
	query := db.Order("DepartmentID")
	query.Offset(0).Limit(20)
	var results []model.Department
	if err := query.Find(&results).Error; err != nil {

	}

	doc := helpers.CreateSimpleDoc()
	title := helpers.CreateSimpleTitle("Sample")
	table := dto.TableStruct{
		Font: dto.FontStruct{
			Family: "Times",
			Style:  "",
			Size:   12,
		},
		Color: dto.ColorStruct{
			R: 255,
			G: 255,
			B: 255,
		},
	}
	data := make([]interfaces.SimpleDoc, len(results))
	for i, k := range results {
		data[i] = k
	}

	templates.GenerateSimpleTable(doc, title, data, table)

}
