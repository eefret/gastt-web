package controllers

import (
	"github.com/astaxie/beego"
)

type ApiController struct {
	beego.Controller
}

func (a *ApiController) Get() {
	a.Ctx.WriteString("Api Controller Get")
}
