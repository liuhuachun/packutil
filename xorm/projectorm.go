package xorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
	"time"
)

var engine *xorm.Engine

type Project struct {
	Id         int64
	Name       string    `xorm:"unique"`
	CreateTime time.Time `xorm:"created"`
	CreateUser string
	Desc       string
}

func init() {
	InitDb()
	if err := engine.Sync2(new(Project)); err != nil {
		log.Fatalf("Fail to Sync database %v", err)
	}
}

//数据库的连接测试
func TestDbCon() error {
	err := engine.Ping()
	return err
}

//保存一个项目，通过项目的属性
func SaveProject(name, createUser, desc string) error {
	_, err := engine.Insert(&Project{Name: name, Desc: desc})
	return err
}

//保存一个项目，参数为项目的结构体
func SaveProjectObject(proj Project) error {
	_, err := engine.Insert(&proj)
	return err
}

//通过id删除记录
func DeleteById(id int64) error {
	_, err := engine.Delete(&Project{Id: id})
	return err
}

//删除传过来的数据
func DeleteProjectByObj(obj *Project) error {
	_, err := engine.Delete(obj)
	return err
}

//获取数据库中项目的数量
func GetProjectTotal() (int64, error) {
	counts, err := engine.Count(&Project{})
	return counts, err
}

//查询所有的项目信息
func FindAllDate() ([]Project, error) {
	records := make([]Project, 0)
	err := engine.Desc("Id").Find(&records)
	if err != nil {
		log.Fatalf("query all project fail , %v", err)
		return records, err
	}
	return records, nil
}

//返回工程的combox的数据
func FindAllProjDataCombox() []*Species {
	total, err := GetProjectTotal()
	if err != nil {
		total = int64(100)
	}
	records := make([]*Species, total)
	recordDates, err := FindAllDate()
	if err != nil {
		log.Fatalf("query all project fail , %v", err)
		return records
	}
	fmt.Println(len(recordDates))
	for i := range records {
		records[i] = &Species{
			Key:  recordDates[i].Name,
			Name: recordDates[i].Name,
		}
	}
	fmt.Println(records)
	return records
}
