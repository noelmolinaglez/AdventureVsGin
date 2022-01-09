package model

import (
	"github.com/noelmolinaglez/simplePdf/pkg/dto"
	"strconv"
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
		Width:   3,
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
		Width:   4,
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
		Width:   3,
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
	values[0] = dto.CellStruct{
		Width:   2,
		Height:  10,
		Content: strconv.Itoa(int(d.DepartmentID)),
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    false,
		Link:    0,
		LinkStr: "",
	}

	values[1] = dto.CellStruct{
		Width:   3,
		Height:  10,
		Content: d.Name,
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    false,
		Link:    0,
		LinkStr: "",
	}

	values[2] = dto.CellStruct{
		Width:   4,
		Height:  10,
		Content: d.GroupName,
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    false,
		Link:    0,
		LinkStr: "",
	}

	values[3] = dto.CellStruct{
		Width:   3,
		Height:  10,
		Content: d.ModifiedDate,
		Border:  "1",
		Ln:      0,
		Align:   "C",
		Fill:    false,
		Link:    0,
		LinkStr: "",
	}

	return values
}
