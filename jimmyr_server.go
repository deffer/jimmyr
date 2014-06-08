// jimmyr_server
package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/deffer/jimmyr/routers"
)

func main() {
	fmt.Println("Starting server...")
	beego.SetStaticPath("/css", "static/css")
	beego.SetStaticPath("/img", "static/img")
	beego.SetStaticPath("/robots.txt", "static/robots.txt")
	beego.Run()
}
