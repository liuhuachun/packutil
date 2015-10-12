package xorm

import (
	"fmt"
	"log"
	"time"
)

type Patch struct {
	Id          int64
	ProjectName string
	VersionName string
	PatchName   string    `xorm:"unique"`
	CreateTime  time.Time `xorm:"created"`
	CreateUser  string
	Path        string
	Desc        string
}
type Species struct {
	Key  string
	Name string
}

func init() {
	InitDb()
	if err := engine.Sync2(new(Patch)); err != nil {
		log.Fatalf("Fail to Sync database %v", err)
	}
}

//保存一个补丁，通过补丁的属性
func SavePatchByParam(name, createUser, desc string) error {
	_, err := engine.Insert(&Patch{PatchName: name, CreateUser: createUser, Desc: desc})
	return err
}

//保存一个补丁，参数为补丁的结构体
func SavePatchByObj(patch *Patch) error {
	fmt.Println(patch)
	_, err := engine.Insert(patch)
	return err
}

//查询所有的补丁信息
func FindAllPatchInfo() ([]Patch, error) {
	records := make([]Patch, 0)
	err := engine.Desc("Id").Find(&records)
	if err != nil {
		log.Fatalf("query all project fail , %v", err)
		return records, err
	}
	return records, nil
}

//删除传过来的数据
func DeletePatchByObj(obj *Patch) error {
	_, err := engine.Delete(obj)
	return err
}
