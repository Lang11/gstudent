package gsdb

import (
	"github.com/astaxie/beego/orm"
	"github.com/Lang11/gstudent/models"
	_ "github.com/lib/pq"
)

func init()  {
	orm.RegisterDriver("postgres", orm.DRPostgres) // 注册驱动
	orm.RegisterDataBase("default","postgres","user=postgres password=pg dbname=test host=127.0.0.1 port=5432 sslmode=disable")
	orm.RegisterModel(new(models.GStudent))
	orm.RunSyncdb("default",false,true)
	orm.Debug = true
}