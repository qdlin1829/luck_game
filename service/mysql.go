package service

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

const (
	driveName = "mysql"
	dsName    = "root:root@(127.0.0.1:3306)/gin-vue?charset=utf8"
	showSQL   = true
	maxCon    = 10
	maxIdleConn = 100
)

var Db *xorm.Engine

func init(){
	initMysql()
}

func initMysql () {
	var err error
	Db, err = xorm.NewEngine(driveName, dsName)
	if err != nil {
		log.Fatal(err)
		return
	}

	Db.ShowSQL(showSQL)
	Db.SetMaxIdleConns(maxIdleConn)
	Db.SetMaxOpenConns(maxCon)

	fmt.Println("init data base ok")
}
