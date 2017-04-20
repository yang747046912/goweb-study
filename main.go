package main

import (
	_ "demo/routers"
	_"demo/initial"
	"github.com/astaxie/beego"
)

func main() {
	beego.Run()
}

