// UpdatePackUtil project main.go
package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
	"log"
	"packutil/xorm"
)

const (
	DEFAULT_PROJECT_NUM = 50
)

var isSpecialMode = walk.NewMutableCondition()

type SysMainWin struct {
	*walk.MainWindow
}

func main() {

	var err error

	MustRegisterCondition("isSpecialMode", isSpecialMode)

	mw := new(SysMainWin)

	sysWin := MainWindow{}

	sysWin.Title = "JAVA_WEB升级包制作工具"
	sysWin.MinSize = Size{600, 400}

	sysWin.MenuItems = []MenuItem{
		Menu{
			Text: "工程",
			Items: []MenuItem{
				Action{
					Text:     "工程管理",
					Shortcut: Shortcut{walk.ModControl, walk.KeyQ},
					OnTriggered: func() {
						CreateProj_Query(mw)
					},
				},
				Separator{},
				Action{
					Text:        "退出",
					Shortcut:    Shortcut{walk.ModControl, walk.KeyE},
					OnTriggered: func() { mw.Close() },
				},
			},
		},
		Menu{
			Text: "版本",
			Items: []MenuItem{
				Action{
					Text:     "版本管理",
					Shortcut: Shortcut{walk.ModControl, walk.KeyQ},
					OnTriggered: func() {
						CreateProj_Query(mw)
					},
				},
				Separator{},
				Action{
					Text:        "退出",
					Shortcut:    Shortcut{walk.ModControl, walk.KeyE},
					OnTriggered: func() { mw.Close() },
				},
			},
		},
		Menu{
			Text: "补丁",
			Items: []MenuItem{
				Action{
					Text:     "补丁管理",
					Shortcut: Shortcut{walk.ModControl, walk.KeyQ},
					OnTriggered: func() {
						CreateProj_Query(mw)
					},
				},
				Separator{},
				Action{
					Text:        "退出",
					Shortcut:    Shortcut{walk.ModControl, walk.KeyE},
					OnTriggered: func() { mw.Close() },
				},
			},
		},
		Menu{
			Text: "帮助",
			Items: []MenuItem{
				Action{
					//AssignTo:    &showAboutBoxAction,
					Text:     "关于",
					Shortcut: Shortcut{walk.ModControl, walk.KeyA},
					//OnTriggered: mw.showAboutBoxAction_Triggered,
				},
			},
		},
	}

	sysWin.AssignTo = &mw.MainWindow
	err = sysWin.Create()
	if err != nil {
		log.Panic(err)
	}
	mw.Run()
}

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
	var dialog = Dialog{}
	model := xorm.NewProjectModel()

	dialog.AssignTo = &dlg
	dialog.Title = "项目管理"
	dialog.Layout = VBox{}
	dialog.MinSize = Size{650, 300}
	dialog.Children = []Widget{
		TableView{
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
					Text:      "删除",
					OnClicked: func() { dlg.Cancel() },
				},
			},
		},
	}
	dialog.Run(owner)
}
