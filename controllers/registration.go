package controllers

import (
	"beego-rest-api/models"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type RegistrationController struct {
	beego.Controller
}

func (this *RegistrationController) Post() {
	//var req database.UserForm
	var req models.UserData
	req.FirstName = this.GetString("firstname")
	req.LastName = this.GetString("lastname")
	req.Phone = this.GetString("phone")
	req.Email = this.GetString("email")
	req.Password = this.GetString("password")
	req.DateOfBrith = this.GetString("dob")

	o := orm.NewOrm()

	//o.Using("default")

	if _, err := o.Insert(&req); err != nil {
		this.Ctx.Output.SetStatus(309)
		this.Data["json"] = "Email already used"
		this.ServeJSON()
		return
	}
	this.Ctx.Output.SetStatus(200)
	this.Data["json"] = "Registration Successful"
	this.ServeJSON()
	return
}
