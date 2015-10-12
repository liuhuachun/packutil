package view

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"packutil/xorm"
)

/**
 *添加项目的方法
 */
func CreateProj_Triggered(owner walk.Form, model *xorm.ProjectModel) {
	var dlg *walk.Dialog
	var db *walk.DataBinder
	var ep walk.ErrorPresenter
	var project = new(xorm.Project)
	var acceptPB, cancelPB *walk.PushButton

	var dialog = Dialog{}
	dialog.AssignTo = &dlg
	dialog.Title = "创建工程"
	dialog.DataBinder = DataBinder{
		AssignTo:       &db,
		DataSource:     project,
		ErrorPresenter: ErrorPresenterRef{&ep},
	}
	dialog.MinSize = Size{300, 200}
	dialog.Layout = VBox{}
	dialog.DefaultButton = &acceptPB
	dialog.CancelButton = &cancelPB

	childrens := []Widget{
		Composite{
			Layout: Grid{Columns: 2},
			Children: []Widget{
				Label{
					Text: "名称:",
				},
				LineEdit{
					Text:      Bind("Name"),
					MaxLength: 10,
				},
				Label{
					Text: "创建者:",
				},
				LineEdit{
					Text:      Bind("CreateUser"),
					MaxLength: 10,
				},
				Label{
					Text: "描述:",
				},
				TextEdit{
					Text:    Bind("Desc"),
					MinSize: Size{300, 50},
				},
				LineErrorPresenter{
					AssignTo:   &ep,
					ColumnSpan: 2,
				},
			},
		},
		Composite{
			Layout: HBox{},
			Children: []Widget{
				HSpacer{},
				PushButton{
					AssignTo: &acceptPB,
					Text:     "保存",
					OnClicked: func() {
						if err := db.Submit(); err != nil {
							log.Print(err)
							return
						}
						xorm.SaveProjectObject(*project)
						model.ResetRows()
						dlg.Accept()
					},
				},
				PushButton{
					AssignTo:  &cancelPB,
					Text:      "取消",
					OnClicked: func() { dlg.Cancel() },
				},
			},
		},
	}
	dialog.Children = childrens
	dialog.Run(owner)
}

/**
** 查看所有项目的方法
**/
func CreateProj_Query(owner walk.Form) {
	var dlg *walk.Dialog
	var tv *walk.TableView
	var dialog = Dialog{}
	model := xorm.NewProjectModel()

	dialog.AssignTo = &dlg
	dialog.Title = "项目管理"
	dialog.Layout = VBox{}
	dialog.MinSize = Size{650, 300}
	dialog.Children = []Widget{
		TableView{
			AssignTo:              &tv,
			AlternatingRowBGColor: walk.RGB(255, 255, 224),
			ColumnsOrderable:      true,
			Columns: []TableViewColumn{
				{Title: "编号", Width: 50},
				{Title: "名称"},
				{Title: "创建者"},
				{Title: "创建时间", Format: "2006-01-02 15:04:05", Width: 130},
				{Title: "描述", Width: 200},
			},
			Model: model,
		},
		Composite{
			Layout: HBox{},
			Children: []Widget{
				HSpacer{},
				PushButton{
					Text: "创建",
					OnClicked: func() {
						CreateProj_Triggered(owner, model)
					},
				},
				PushButton{
					Text: "删除",
					OnClicked: func() {
						indexs := tv.SelectedIndexes()
						if indexs.Len() == 0 {
							walk.MsgBox(owner, "提示", "请选择要删除的数据", walk.MsgBoxIconError)
							return
						}
						walk.MsgBox(owner, "提示", "确认是否删除此工程", walk.MsgBoxOKCancel)

						obj := model.GetItemByindex(int64(indexs.At(0)))
						xorm.DeleteProjectByObj(obj)
						model.ResetRows()
					},
				},
			},
		},
	}
	dialog.Run(owner)
}
