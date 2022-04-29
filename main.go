package main

import (
	"github.com/astaxie/beego"
	_ "goweb-be/routers"
)

func main() {
	//beego.SetViewsPath("front")
	//beego.SetStaticPath("/front", "front") // url是路径别名，path是项目目录 eg:/front/css/public.css
	beego.Run()
}
