package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "hello/models"
	_ "hello/routers"
)

func init() {
	_ = orm.RegisterDriver("mysql", orm.DRMySQL)
	_ = orm.RegisterDataBase("default", "mysql", "root:a123456789@/orm_test?charset=utf8")

}

func main() {
	//自动建表
	//orm.RunCommand()
	//// 数据库别名
	//name := "default"
	//
	//// drop table 后再建表
	//force := true
	//
	//// 打印执行过程
	//verbose := true
	//// 遇到错误立即返回
	//err := orm.RunSyncdb(name, force, verbose)
	//if err != nil {
	//	fmt.Println(err)
	//}

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.Run()
	//orm.Debug = true
	//o := orm.NewOrm()
	//_ = o.Using("default")
	//
	//profile := new(models.Profile)
	//profile.Age = 26
	//user := new(models.User)
	//user.Profile = profile
	//user.Name = "lecon"
	//
	//fmt.Println(o.Insert(profile))
	//fmt.Println(o.Insert(user))

}
