package xorm

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/lxn/walk"
	"log"
	"os"
	"packutil/util"
	"strings"
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

type ProjectModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*Project
}

func init() {

	err := util.InitConf()
	if err != nil {
		return
	}
	err = initDb()

	if err != nil {
		log.Fatalf("Fail to create engine %v.", err)
	}

	if err := engine.Sync2(new(Project)); err != nil {
		log.Fatalf("Fail to Sync database %v", err)
	}

	engine.ShowSQL = true   //则会在控制台打印出生成的SQL语句
	engine.ShowDebug = true //则会在控制台打印调试信息；
	engine.ShowErr = true   //则会在控制台打印错误信息；
	engine.ShowWarn = true  //则会在控制台打印警告信息；

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
		return nil
	}
	pass, err := util.GetValue("db", "dbpass")
	if err != nil {
		log.Fatalf("get value of dbpass from conf, %v", err)
		return nil
	}

	dbhost, err := util.GetValue("db", "dbhost")
	if err != nil {
		log.Fatalf("get value of dbhost from conf, %v", err)
		return nil
	}
	dbport, err := util.GetValue("db", "dbport")
	if err != nil {
		log.Fatalf("get value of dbport from conf, %v", err)
		return nil
	}
	dbname, err := util.GetValue("db", "dbname")
	if err != nil {
		log.Fatalf("get value of dbname from conf, %v", err)
		return nil
	}

	dsn := []string{userName, ":", pass, "@tcp(", dbhost, ":", dbport, ")/", dbname}
	dsnStr := strings.Join(dsn, "")
	fmt.Println(dsnStr)
	engine, err = xorm.NewEngine("mysql", dsnStr)
	return err
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

//****************
func NewProjectModel() *ProjectModel {
	m := new(ProjectModel)
	m.ResetRows()
	return m
}

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *ProjectModel) RowCount() int {
	return len(m.items)
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *ProjectModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Id

	case 1:
		return item.Name

	case 2:
		return item.CreateUser

	case 3:
		return item.CreateTime
	case 4:
		return item.Desc
	case 5:
		return item.Id
	}

	panic("unexpected col")
}

// Called by the TableView to sort the model.
func (m *ProjectModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	//	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

func (m *ProjectModel) ResetRows() {
	projects, err := FindAllDate()
	if err != nil {
		log.Fatalf("query all project fail, %v", err)
		projects = make([]Project, 0)
	}

	m.items = make([]*Project, len(projects))
	for i, _ := range projects {
		m.items[i] = &projects[i]
		fmt.Println(m.items[i])
	}
	m.Sort(m.sortColumn, m.sortOrder)
}
