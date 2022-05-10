package chapter03

import "github.com/astaxie/beego"

type Temp01Controller struct {
	beego.Controller
}

func (t *Temp01Controller) Get() {
	t.Data["name"] = "xiao"
	t.Data["arr"] = [3]int{1, 2, 3}
	t.Data["html"] = `<a href="http://www.baidu.com">百度</a>`
	t.Data["username"] = "xiaoxiao"
	t.TplName = "chapter03/index.html"
}
