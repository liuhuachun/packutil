package view

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

var isSpecialMode = walk.NewMutableCondition()

type SysMainWin struct {
	*walk.MainWindow
}

func InitSysMainWin(mw *SysMainWin) *MainWindow {
	MustRegisterCondition("isSpecialMode", isSpecialMode)

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
						CreateVersion_Query(mw)
					},
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
						CreatePatch_Query(mw)
					},
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
	return &sysWin
}
