package controllers

import (
	"CicilAppointmentApp/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

type ApiController struct {
	beego.Controller
}

func (this *ApiController) AllSchedule() {
	o := orm.NewOrm()
	o.Using("default")
	var appointment []*models.Appointment
	num, err := o.QueryTable("appointment").All(&appointment)

	if err != orm.ErrNoRows && num > 0 {
		this.Data["json"] = &appointment
		this.ServeJSON()
	}
}

func (this *ApiController) Schedule() {
	appointmentId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	o.Using("default")
	appointment := models.Appointment{Id: appointmentId}

	err := o.Read(&appointment)

	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
	}
	this.Data["json"] = &appointment
	this.ServeJSON()
}

func (this *ApiController) PostSchedule() {
	o := orm.NewOrm()
	o.Using("default")
	var appointment models.Appointment
	appointment.Title = this.GetString("title")
	appointment.Date = this.GetString("date")
	appointment.Time = this.GetString("time")
	appointment.Description = this.GetString("desc")
	appointment.Attendes = this.GetString("atten")
	appointment.Location = this.GetString("loc")

	id, err := o.Insert(&appointment)
	if err == nil {
		fmt.Println(id)
	}

	this.Data["json"] = "Success"
	this.ServeJSON()
}

func (this *ApiController) UpdateSchedule() {
	appointmentId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	o := orm.NewOrm()
	o.Using("default")
	appointment := models.Appointment{Id: appointmentId}
	if o.Read(&appointment) == nil {
		appointment.Title = this.GetString("title")
		appointment.Date = this.GetString("date")
		appointment.Time = this.GetString("time")
		appointment.Description = this.GetString("desc")
		appointment.Attendes = this.GetString("atten")
		appointment.Location = this.GetString("loc")
		if num, err := o.Update(&appointment); err == nil {
			fmt.Println(num)
		}
	}

	this.Data["json"] = "Success"
	this.ServeJSON()
}

func (this *ApiController) DeleteSchedule() {
	o := orm.NewOrm()
	o.Using("default")
	appointmentId, _ := strconv.Atoi(this.Ctx.Input.Param(":id"))
	if num, err := o.Delete(&models.Appointment{Id: appointmentId}); err == nil {
		fmt.Println(num)
	}
	this.Data["json"] = "Success"
	this.ServeJSON()
}
