package models

type Sample struct {
	Id     int
	Value1 string
	Value2 string
}

func (m *Sample) TableName() string {
	return "sample"
}
