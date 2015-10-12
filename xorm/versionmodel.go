package xorm

import (
	"github.com/lxn/walk"
	"log"
)

type VersionModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*Version
}

func NewVersionModel() *VersionModel {
	m := new(VersionModel)
	m.ResetRows()
	return m
}

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *VersionModel) RowCount() int {
	return len(m.items)
}

//通过索引获取数据
func (m *VersionModel) GetVersionByindex(index int64) *Version {
	return m.items[index]
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *VersionModel) Value(row, col int) interface{} {
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
	}

	panic("unexpected col")
}

// Called by the TableView to sort the model.
func (m *VersionModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	//	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

func (m *VersionModel) ResetRows() {
	versions, err := FindAllVersionInfo()
	if err != nil {
		log.Fatalf("query all project fail, %v", err)
		versions = make([]Version, 0)
	}

	m.items = make([]*Version, len(versions))
	for i, _ := range versions {
		m.items[i] = &versions[i]
	}
	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}
