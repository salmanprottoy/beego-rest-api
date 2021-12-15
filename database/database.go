package database

import (
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

type UserForm struct {
	Id          int    `orm:"auto"; primary_key`
	FirstName   string `json:"firstname" orm:"size(40)`
	LastName    string `json:"lastname" orm:"size(40)`
	Phone       string `json:"phone" orm:"size(50)`
	Email       string `json:"email" orm:"size(64);unique"`
	Password    string `json:"password"`
	DateOfBrith string `json:"dob"`
}

type Orm struct {
	CreatedORM orm.Ormer
}

func InitDB() {
	orm.RegisterDriver("postgres", orm.DRPostgres)
	orm.RegisterDataBase("default", "postgres", "user=postgres password=admin host=127.0.0.1 port=5432 dbname=userdb sslmode=disable")
	orm.RegisterModel(new(UserForm))
	orm.RunSyncdb("default", false, true)
}
