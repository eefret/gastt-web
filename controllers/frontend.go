package controllers

import (
	"github.com/astaxie/beego"
)

type FrontendController struct {
	beego.Controller
}

func (f *FrontendController) Get() {
	f.TplNames = "google_btn.tpl"
	f.Ctx.WriteString("Hello World!")
}
