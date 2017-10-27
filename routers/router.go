package routers

import (
	"jobs/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/emp", &controllers.EmpController{})
    beego.Router("/emp/create", &controllers.EmpController{},"*:Create")
    beego.Router("/emp/getall", &controllers.EmpController{},"*:GetAll")
    beego.Router("/emp/show", &controllers.EmpController{},"*:Show")
    beego.Router("/mold", &controllers.MoldController{})
    beego.Router("/mold/create", &controllers.MoldController{},"*:Create")
    beego.Router("/mold/getall", &controllers.MoldController{},"*:GetAll")
    beego.Router("/mold/show", &controllers.MoldController{},"*:Show")
    beego.Router("/logs", &controllers.LogsController{})
    beego.Router("/logs/create", &controllers.LogsController{},"*:Create")
    beego.Router("/logs/search", &controllers.LogsController{},"*:SearchByMold")
    beego.Router("/logs/searchemp", &controllers.LogsController{},"*:SearchByEmp")
    beego.Router("/logs/show", &controllers.LogsController{},"*:Show")
    beego.Router("/codes", &controllers.CodesController{})
    beego.Router("/codes/getall", &controllers.CodesController{},"*:GetAll")
}
