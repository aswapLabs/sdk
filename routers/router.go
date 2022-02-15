package routers

import (
	"aswap-go/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/api/pair/reg", &controllers.PairsController{},"get,post:Register")
}
