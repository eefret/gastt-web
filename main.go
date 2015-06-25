package main

import (
	"github.com/astaxie/beego"
	_ "github.com/eefret/gastt-web/routers"
)

func main() {
	beego.Run()
}
