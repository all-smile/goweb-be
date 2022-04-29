package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

type User struct {
	Id   int
	Name string
	Age  int
}

func (c *MainController) Get() {
	// 直接渲染字符串到浏览器， 模板文件不会渲染了
	//c.Ctx.WriteString("直接渲染字符串到浏览器")

	//c.Data["Website"] = "beego.me"
	//c.Data["Email"] = "astaxie@gmail.com"
	//c.TplName = "index.tpl"

	var user = User{
		Id:   1,
		Name: "xiaoming",
		Age:  16,
	}
	c.Data["user"] = user

	arr := [5]int{1, 2, 3, 4, 5}
	c.Data["arr"] = arr

	users := [3]User{
		{Id: 1, Name: "xixi", Age: 16},
		{Id: 2, Name: "haha", Age: 15},
		{Id: 3, Name: "lala", Age: 16},
	}
	c.Data["users"] = users

	map_data := map[string]interface{}{
		"name":   "Go",
		"author": "idontkonw",
		"age":    18,
	}
	c.Data["map_data"] = map_data

	map_struct := map[string]User{
		"JS": {
			Id:   1,
			Name: "javascript",
			Age:  12,
		},
		"TS": {
			Id:   2,
			Name: "typescript",
			Age:  10,
		},
	}
	c.Data["map_struct"] = map_struct

	slice := []int{1, 2, 3, 4, 5}
	c.Data["slice"] = slice

	c.TplName = "index.html"
}

func (c *MainController) Post() {

}
