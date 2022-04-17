package model

type Test struct {
	Id        int
	Balance   int
	Frozen    int
	Available int
}

func (Test) TableName() string {
	return "t_test"
}
