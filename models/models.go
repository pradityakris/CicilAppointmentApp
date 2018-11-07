package models

import "github.com/astaxie/beego/orm"

type Appointment struct {
	Id          int    `form:"-"`
	Title       string `form:"title,text,title:"`
	Date        string `form:"date,text,date:"`
	Time        string `form:"time,text,time:"`
	Description string `form:"desc,text,desc:"`
	Attendes    string `form:"atten,text,atten:"`
	Location    string `form:"loc,text,loc:"`
}

func init() {
	orm.RegisterModel(new(Appointment))
}

func (a *Appointment) TableName() string {
	return "appointment"
}
