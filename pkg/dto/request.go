package dto

type Request struct {
	Action string `form:"action"`
	Type   string `form:"type"`
	Start  int    `form:"start"`
	Limit  int    `form:"limit"`
	Field  string `form:"field"`
	Dir    string `form:"dir"`
}
