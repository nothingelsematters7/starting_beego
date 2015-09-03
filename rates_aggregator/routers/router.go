package routers

import (
	"github.com/nothingelsematters7/starting_beego/rates_aggregator/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
