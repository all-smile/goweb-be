package chapter02

import (
	"github.com/astaxie/beego"
)

type OtherTypeController struct {
	beego.Controller
}

type Teacher struct {
	Id   int
	Name string
	Age  int
}

func (o *OtherTypeController) Get() {
	teacher := Teacher{
		Id:   1,
		Name: "xiao",
		Age:  26,
	}
	o.Data["json"] = teacher
	o.ServeJSON()

	//o.Data["xml"] = teacher
	//o.ServeXML()

	//o.Data["yaml"] = teacher
	//o.ServeYAML()
}
