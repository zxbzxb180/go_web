package main

import (
	_ "beego_web/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	orm.RegisterDataBase("default", "mysql", "web:cx6222580@tcp(139.198.181.33:13306)/test")
	beego.Run()
}
