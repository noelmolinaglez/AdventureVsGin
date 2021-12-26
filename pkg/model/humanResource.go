package model

type Department struct {
	DepartmentId uint   `json:"DepartmentID" gorm:"primary_key"`
	Name         string `json:"Name"`
	GroupName    string `json:"GroupName"`
	ModifiedDate string `json:"ModifiedDate"`
}
