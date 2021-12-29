package dto

type Request struct {
	Start int `form:"start"`
	Limit int `form:"limit"`
}
