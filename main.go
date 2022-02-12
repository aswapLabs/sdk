package main

import (
	_ "aswap-go/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

