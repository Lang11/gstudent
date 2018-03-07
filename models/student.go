package models

type GStudent struct {
	Id    int    `form:"id" valid:"Required" orm:"auto"`
	Name  string `form:"name" valid:"Required" orm:"size(100)"`
	Sex   int    `form:"sex" valid:"Required"`
	Class string `form:"class" valid:"Required"`
}
