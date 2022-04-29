package chapter02

import (
	"fmt"
	"github.com/astaxie/beego"
	"time"
)

type UploadController struct {
	beego.Controller
}

func (u *UploadController) Get() {
	u.TplName = "chapter02/upload.html"
}

func (u *UploadController) Post() {
	// 1. 获取上传文件
	f, h, err := u.GetFile("file")
	defer f.Close()
	fmt.Println(f)
	fmt.Println(h.Filename)
	fmt.Println(err)

	// 2.保存
	time_unix := time.Now().Unix()
	filePath := fmt.Sprintf("%d_%s", time_unix, h.Filename)
	fmt.Println("time_unix=", time_unix, "upload/"+filePath)
	err = u.SaveToFile("file", "upload/"+filePath)
	if err != nil {
		fmt.Println("保存失败", err)
	}
	u.Ctx.WriteString("ok")
}
