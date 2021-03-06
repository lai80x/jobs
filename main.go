package main

import (
	_ "jobs/routers"
	_ "jobs/models"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	dbAlias := beego.AppConfig.String("mysql::db_alias")
	dbName := beego.AppConfig.String("mysql::db_name")
	dbUser := beego.AppConfig.String("mysql::db_user")
	dbPwd := beego.AppConfig.String("mysql::db_pwd")
	//dbPort := beego.AppConfig.String("mysql::db_port")
	//dbHost := beego.AppConfig.String("mysql::db_host")
	dbInfo := dbUser + ":" + dbPwd + "@/" + dbName + "?charset=utf8"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase(dbAlias, "mysql", dbInfo)
}

func main() {
	beego.Run()
}

