package chapter02

import (
	"fmt"
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

type Student struct {
	Id       int
	Name     string `form:"user_name"`
	Password string `form:"password"`
}

func (u *UserController) Get() {
	//id := u.GetString(":id")
	id := u.GetString("id")
	//id, err := u.GetInt("id")
	//if err != nil {
	//	u.Ctx.WriteString("id只能是int类型")
	//}
	fmt.Println("id=", id)
	u.Data["id"] = id
	u.TplName = "chapter02/user.html"
}

// add_user post
func (u *UserController) Post() {
	// 获取post
	//user_name := u.GetString("user_name")
	//password := u.GetString("password")
	//fmt.Println("user_name=", user_name, "password=", password)

	// form表单解析到结构体
	//student := Student{}
	//u.ParseForm(&student)
	//fmt.Println("student=", student, student.Name, student.Password)

	// 获取Ajax数据
	user_name := u.GetString("user_name")
	password := u.GetString("password")
	fmt.Println("user_name=", user_name, "password=", password)
	//u.Ctx.WriteString("提交成功")
	res := map[string]interface{}{
		"code": 200,
		"msg":  "添加成功",
	}
	u.Data["json"] = res
	u.ServeJSON()
}
