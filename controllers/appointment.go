package controllers

import (
	"CicilAppointmentApp/models"
	"fmt"
	"strconv"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/validation"
)

type AppointmentController struct {
	beego.Controller
}

func (a *AppointmentController) Home() {
	a.Layout = "layout.tpl"
	a.LayoutSections = make(map[string]string)
	a.LayoutSections["Header"] = "header.tpl"
	a.LayoutSections["Footer"] = "footer.tpl"
	a.TplName = "appointment/home.tpl"
}

func (a *AppointmentController) Add() {
	a.Data["Form"] = &models.Appointment{}
	a.Layout = "layout.tpl"
	a.LayoutSections = make(map[string]string)
	a.LayoutSections["Header"] = "header.tpl"
	a.LayoutSections["Footer"] = "footer.tpl"
	a.TplName = "appointment/add.tpl"

	flash := beego.ReadFromRequest(&a.Controller)

	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		a.Data["flash"] = ok
	}

	o := orm.NewOrm()
	o.Using("default")

	appointment := models.Appointment{}

	if err := a.ParseForm(&appointment); err != nil {
		beego.Error("Err. Reason: ", err)
	} else {
		a.Data["Appointment"] = appointment
		valid := validation.Validation{}
		isValid, _ := valid.Valid(appointment)

		if a.Ctx.Input.Method() == "POST" {
			if !isValid {
				a.Data["Errors"] = valid.ErrorsMap
				beego.Error("Form didn't validate.")
			} else {
				searchAppointment := models.Appointment{Title: appointment.Title}
				beego.Debug("appointment name supplied:", appointment.Title)
				err = o.Read(&searchAppointment)
				beego.Debug("Err:", err)
				flash := beego.NewFlash()

				if err == orm.ErrNoRows || err == orm.ErrMissPK {
					beego.Debug("No appointment found matching details supplied. Attempting to insert appointment: ", appointment)
					id, err := o.Insert(&appointment)
					if err == nil {
						fmt.Println(id)
						a.Redirect("/appointment/view", 302)
					} else {
						msg := fmt.Sprint("Couldn't insert new appointment. Reason: ", err)
						beego.Debug(msg)
						flash.Error(msg)
						flash.Store(&a.Controller)
					}
				} else {
					beego.Debug("appointment found matching details supplied. Cannot insert")
				}
			}
		}
	}
}

func (a *AppointmentController) View() {
	a.Layout = "layout.tpl"
	a.LayoutSections = make(map[string]string)
	a.LayoutSections["Header"] = "header.tpl"
	a.LayoutSections["Footer"] = "footer.tpl"
	a.TplName = "appointment/view.tpl"

	flash := beego.ReadFromRequest(&a.Controller)

	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		a.Data["errors"] = ok
	}

	if ok := flash.Data["notice"]; ok != "" {
		// Display error messages
		a.Data["notices"] = ok
	}

	o := orm.NewOrm()
	o.Using("default")

	var appointment []*models.Appointment
	num, err := o.QueryTable("appointment").All(&appointment)

	if err != orm.ErrNoRows && num > 0 {
		a.Data["records"] = appointment
	}
}

func (a *AppointmentController) Edit() {
	a.Layout = "layout.tpl"
	a.LayoutSections = make(map[string]string)
	a.LayoutSections["Header"] = "header.tpl"
	a.LayoutSections["Footer"] = "footer.tpl"
	a.TplName = "appointment/edit.tpl"

	flash := beego.ReadFromRequest(&a.Controller)

	if ok := flash.Data["error"]; ok != "" {
		// Display error messages
		a.Data["errors"] = ok
	}

	if ok := flash.Data["notice"]; ok != "" {
		// Display error messages
		a.Data["notices"] = ok
	}

	appointmentId, _ := strconv.Atoi(a.Ctx.Input.Param(":id"))

	if a.Ctx.Input.Method() == "POST" {
		o := orm.NewOrm()
		appointment := models.Appointment{Id: appointmentId}
		if o.Read(&appointment) == nil {
			appointment.Title = a.GetString("title")
			appointment.Date = a.GetString("date")
			appointment.Time = a.GetString("time")
			appointment.Description = a.GetString("desc")
			appointment.Attendes = a.GetString("atten")
			appointment.Location = a.GetString("loc")
			if num, err := o.Update(&appointment); err == nil {
				fmt.Println(num)
			}
		}
		a.Redirect("/appointment/view", 302)
	}

	o := orm.NewOrm()
	o.Using("default")
	appointment := models.Appointment{Id: appointmentId}

	err := o.Read(&appointment)

	if err == orm.ErrNoRows {
		fmt.Println("No result found.")
	} else if err == orm.ErrMissPK {
		fmt.Println("No primary key found.")
	} else {
		fmt.Println(appointment.Id)
	}
	//appointment := o.Raw("select * from appointment where Id = ?", appointmentId)
	//appointment := models.Appointment{Id: appointmentId}
	a.Data["record"] = appointment
}

func (a *AppointmentController) Delete() {
	o := orm.NewOrm()
	appointmentId, _ := strconv.Atoi(a.Ctx.Input.Param(":id"))
	if num, err := o.Delete(&models.Appointment{Id: appointmentId}); err == nil {
		fmt.Println(num)
	}
	a.Redirect("/appointment/view", 302)
}
