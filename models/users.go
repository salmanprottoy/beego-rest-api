package models

import (
	"github.com/beego/beego/v2/client/orm"
)

type Users struct {
	Id          int
	Name        string
	First_Name  string
	Last_Name   string
	Email       string
	Phone       string
	Password    string
	DateOfBirth string
}

func init() {
	orm.RegisterModel(new(Users))
}
