package models

import (
	"github.com/astaxie/beego/orm"
)

type Sample struct {
	Id          int
	Value1        string
	Value2        string
	Profile     *Sample `orm:"rel(one)"` // OneToOne relation
}

func init() {
	// Need to register model in init
	orm.RegisterModel(new(Sample))
}
