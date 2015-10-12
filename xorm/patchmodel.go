package xorm

import (
	"github.com/lxn/walk"
	"log"
)

type PatchModel struct {
	walk.TableModelBase
	walk.SorterBase
	sortColumn int
	sortOrder  walk.SortOrder
	items      []*Patch
}

func NewPatchModel() *PatchModel {
	m := new(PatchModel)
	m.ResetRows()
	return m
}

// Called by the TableView from SetModel and every time the model publishes a
// RowsReset event.
func (m *PatchModel) RowCount() int {
	return len(m.items)
}

//通过索引获取数据
func (m *PatchModel) GetPatchByindex(index int64) *Patch {
	return m.items[index]
}

// Called by the TableView when it needs the text to display for a given cell.
func (m *PatchModel) Value(row, col int) interface{} {
	item := m.items[row]

	switch col {
	case 0:
		return item.Id

	case 1:
		return item.PatchName

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
func (m *PatchModel) Sort(col int, order walk.SortOrder) error {
	m.sortColumn, m.sortOrder = col, order

	//	sort.Stable(m)

	return m.SorterBase.Sort(col, order)
}

func (m *PatchModel) ResetRows() {
	records, err := FindAllPatchInfo()
	if err != nil {
		log.Fatalf("query all project fail, %v", err)
		records = make([]Patch, 0)
	}

	m.items = make([]*Patch, len(records))
	for i, _ := range records {
		m.items[i] = &records[i]
	}
	m.PublishRowsReset()

	m.Sort(m.sortColumn, m.sortOrder)
}
