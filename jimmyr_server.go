// jimmyr_server
package main

import (
	"fmt"
	"github.com/astaxie/beego"
	_ "github.com/deffer/jimmyr/routers"
)

func main() {
	fmt.Println("Starting server...")
	beego.Run()
}
