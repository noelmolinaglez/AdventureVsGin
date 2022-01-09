package model

import (
	"github.com/noelmolinaglez/simplePdf/pkg/dto"
)

func (Department) GetHeaders() []dto.CellStruct {

	headers := make([]dto.CellStruct, 4)

	headers[0] = dto.CellStruct{
		Width:   2,
		Height:  10,
		Content: "DepartmentId",
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    false,
		Link:    0,
		LinkStr: "",
	}
	headers[1] = dto.CellStruct{
		Width:   2,
		Height:  10,
		Content: "Name",
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    false,
		Link:    0,
		LinkStr: "",
	}
	headers[2] = dto.CellStruct{
		Width:   3,
		Height:  10,
		Content: "GroupName",
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    false,
		Link:    0,
		LinkStr: "",
	}
	headers[3] = dto.CellStruct{
		Width:   5,
		Height:  10,
		Content: "ModifiedDate",
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    false,
		Link:    0,
		LinkStr: "",
	}

	return headers
}

func (d Department) GetValues() []dto.CellStruct {

	values := make([]dto.CellStruct, 4)
	//values[0] = helpers.CreateSimpleCell(strconv.Itoa(int(d.DepartmentID)))
	//values[1] = helpers.CreateSimpleCell(d.Name)
	//values[2] = helpers.CreateSimpleCell(d.GroupName)
	//values[3] = helpers.CreateSimpleCell(d.ModifiedDate)

	return values
}
