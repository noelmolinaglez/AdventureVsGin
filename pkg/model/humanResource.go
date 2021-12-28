package model

type Department struct {
	DepartmentId uint   `json:"DepartmentID" gorm:"primary_key"`
	Name         string `json:"Name"`
	GroupName    string `json:"GroupName"`
	ModifiedDate string `json:"ModifiedDate"`
}

type Employee struct {
	BusinessEntityId  uint   `json:"BusinessEntityID" gorm:"primary_key"`
	NationalIdNumber  string `json:"NationalIDNumber"`
	LoginId           string `json:"loginID"`
	OrganizationNode  string `json:"OrganizationNode"`
	OrganizationLevel string `json:"OrganizationLevel"`
	JobTitle          string `json:"JobTitle"`
	BirthDate         string `json:"BirthDate"`
	MaritalStatus     string `json:"MaritalStatus"`
	Gender            string `json:"Gender"`
	HiredDate         string `json:"HiredDate"`
	SalariedFlag      bool   `json:"SalariedFlag"`
	VacationHours     uint   `json:"VacationHours"`
	SickLeaveHours    uint   `json:"SickLeaveHours"`
	CurrentFlag       bool   `json:"CurrentFlag"`
	RowGuid           string `json:"rowguid"`
	ModifiedDate      string `json:"ModifiedDate"`
}

type EmployeeDepartmentHistory struct {
	BusinessEntityId uint   `json:"BusinessEntityID"`
	DepartmentId     uint   `json:"DepartmentID" `
	ShiftId          uint   `json:"ShiftID"`
	StartDate        string `json:"StartDate"`
	EndDate          string `json:"EndDate"`
	ModifiedDate     string `json:"ModifiedDate"`
}
