package model

type Student struct {
	StuNo     string `json:"stu_no"`
	Name      string `json:"name"`
	ClassName string `json:"class_name"`
	Grade     string `json:"grade"`
	Score     int    `json:"score"`
}
