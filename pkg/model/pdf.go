package model

import (
	"github.com/noelmolinaglez/simplePdf/pkg/dto"
	"github.com/noelmolinaglez/simplePdf/pkg/helpers"
	"strconv"
)

func (Department) GetHeaders() []dto.CellStruct {
	columns := []string{
		"DepartmentId",
		"Name",
		"GroupName",
		"ModifiedDate",
	}
	headers := make([]dto.CellStruct, len(columns))
	for i, _ := range headers {
		headers[i] = helpers.CreateSimpleCell(columns[i])
	}

	return headers
}

func (d Department) GetValues() []dto.CellStruct {

	values := make([]dto.CellStruct, 4)
	values[0] = helpers.CreateSimpleCell(strconv.Itoa(int(d.DepartmentID)))
	values[1] = helpers.CreateSimpleCell(d.Name)
	values[2] = helpers.CreateSimpleCell(d.GroupName)
	values[3] = helpers.CreateSimpleCell(d.ModifiedDate)

	return values
}
