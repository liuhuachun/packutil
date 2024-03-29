package xorm

import (
	"github.com/lxn/walk"
	"log"
)

type ProjectModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*Project
}

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

//通过索引获取数据
func (m *ProjectModel) GetItemByindex(index int64) *Project {
	return m.items[index]
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
	}
	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}
