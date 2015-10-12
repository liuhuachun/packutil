package xorm

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"log"
	"os"
	"packutil/util"
	"strings"
)

func InitDb() {
	err := util.InitConf()
	if err != nil {
		return
	}
	err = initDb()

	if err != nil {
		log.Fatalf("Fail to create engine %v.", err)
	}

	engine.ShowSQL = false   //则会在控制台打印出生成的SQL语句
	engine.ShowDebug = false //则会在控制台打印调试信息；
	engine.ShowErr = true    //则会在控制台打印错误信息；
	engine.ShowWarn = false  //则会在控制台打印警告信息；

	f, err := os.Create("./logs/sql.log")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()
	engine.Logger = xorm.NewSimpleLogger(f)
}

/**
**	初始化数据库连接
**/
func initDb() error {
	userName, err := util.GetValue("db", "userName")
	if err != nil {
		log.Fatalf("get value of userName from conf, %v", err)
		return err
	}
	pass, err := util.GetValue("db", "dbpass")
	if err != nil {
		log.Fatalf("get value of dbpass from conf, %v", err)
		return err
	}

	dbhost, err := util.GetValue("db", "dbhost")
	if err != nil {
		log.Fatalf("get value of dbhost from conf, %v", err)
		return err
	}
	dbport, err := util.GetValue("db", "dbport")
	if err != nil {
		log.Fatalf("get value of dbport from conf, %v", err)
		return err
	}
	dbname, err := util.GetValue("db", "dbname")
	if err != nil {
		log.Fatalf("get value of dbname from conf, %v", err)
		return err
	}

	dsn := []string{userName, ":", pass, "@tcp(", dbhost, ":", dbport, ")/", dbname}
	dsnStr := strings.Join(dsn, "")
	engine, err = xorm.NewEngine("mysql", dsnStr)
	return err
}
