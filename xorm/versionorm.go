package xorm

import (
	"log"
	"time"
)

type Version struct {
	Id          int64
	ProjectName string
	Name        string    `xorm:"unique"`
	CreateTime  time.Time `xorm:"created"`
	CreateUser  string
	Path        string
	Desc        string
}

func init() {
	if err := engine.Sync2(new(Version)); err != nil {
		log.Fatalf("Fail to Sync database %v", err)
	}
}

//保存一个版本，通过项目的属性
func SaveVersionByParam(name, createUser, desc string) error {
	_, err := engine.Insert(&Version{Name: name, CreateUser: createUser, Desc: desc})
	return err
}

//查询所有的版本信息
func FindAllVersionInfo() ([]Version, error) {
	records := make([]Version, 0)
	err := engine.Desc("Id").Find(&records)
	if err != nil {
		log.Fatalf("query all project fail , %v", err)
		return records, err
	}
	return records, nil
}

//删除传过来的数据
func DeleteVersionByObj(obj *Version) error {
	_, err := engine.Delete(obj)
	return err
}

//获取数据库中项目的数量
func GetVersionTotal() (int64, error) {
	counts, err := engine.Count(&Version{})
	return counts, err
}

//返回工程的combox的数据
func FindAllVersionDataCombox() []*Species {
	total, err := GetVersionTotal()
	if err != nil {
		total = int64(100)
	}
	records := make([]*Species, total)
	recordDates, err := FindAllVersionInfo()
	if err != nil {
		log.Fatalf("query all project fail , %v", err)
		return records
	}
	for i := range records {
		records[i] = &Species{
			Key:  recordDates[i].Name,
			Name: recordDates[i].Name,
		}
	}
	return records
}
