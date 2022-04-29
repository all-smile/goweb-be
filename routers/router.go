package routers

import (
	"github.com/astaxie/beego"
	"goweb-be/controllers"
	"goweb-be/controllers/chapter02"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	// user/123
	// beego.Router("/user/?:id:int", &controllers.UserController{})
	// user?id=123
	beego.Router("/user", &chapter02.UserController{})
	beego.Router("/add_user", &chapter02.UserController{})
	beego.Router("/add_user_ajax", &chapter02.UserController{})
	beego.Router("/upload", &chapter02.UploadController{})
	beego.Router("/upload_ajax", &chapter02.UploadController{})
}
