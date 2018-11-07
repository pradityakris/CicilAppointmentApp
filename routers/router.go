package routers

import (
	"CicilAppointmentApp/controllers"

	"github.com/astaxie/beego"
)

func init() {
	//CRUD Routes
	beego.Router("/", &controllers.AppointmentController{}, "get:View")
	beego.Router("/appointment/add", &controllers.AppointmentController{}, "get,post:Add")
	beego.Router("/appointment/view", &controllers.AppointmentController{}, "get:View")
	beego.Router("/appointment/edit/:id([0-9]+)", &controllers.AppointmentController{}, "get,post:Edit")
	beego.Router("/appointment/delete/:id([0-9]+)", &controllers.AppointmentController{}, "get,post:Delete")

	//API Routes
	beego.Router("/api/getallschedule", &controllers.ApiController{}, "get:AllSchedule")
	beego.Router("/api/schedule/:id([0-9]+)", &controllers.ApiController{}, "get:Schedule")
	beego.Router("/api/addschedule", &controllers.ApiController{}, "post:PostSchedule")
	beego.Router("/api/updateschedule/:id([0-9]+)", &controllers.ApiController{}, "post:UpdateSchedule")
	beego.Router("/api/deleteschedule/:id([0-9]+)", &controllers.ApiController{}, "get:DeleteSchedule")
}
