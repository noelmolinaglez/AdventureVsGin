package model

type Department struct {
	DepartmentID uint   `json:"DepartmentID" gorm:"column:DepartmentID;primaryKey;autoIncrement:true"`
	Name         string `json:"Name" form:"name" gorm:"column:Name"`
	GroupName    string `json:"GroupName" form:"groupName" gorm:"column:GroupName"`
	ModifiedDate string `json:"ModifiedDate" gorm:"column:ModifiedDate"`
}

func (Department) TableName() string {
	return "HumanResources.Department"
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

type EmployeePayHistory struct {
	BusinessEntityId uint    `json:"BusinessEntityID"`
	RateChangeDate   string  `json:"RateChangeDate"`
	Rate             float32 `json:"Rate"`
	PayFrequency     uint    `json:"PayFrequency"`
	ModifiedDate     string  `json:"ModifiedDate"`
}

type JobCandidate struct {
	JobCandidateId   uint   `json:"JobCandidateID"`
	BusinessEntityId uint   `json:"BusinessEntityID"`
	Resume           string `json:"Resume"`
	ModifiedDate     string `json:"ModifiedDate"`
}

type Shift struct {
	ShiftId      uint   `json:"ShiftID"`
	StartTime    string `json:"StartTime"`
	EndTime      string `json:"EndTime"`
	ModifiedDate string `json:"ModifiedDate"`
}
