// UpdatePackUtil project main.go
package main

import (
	"log"
	"packutil/view"
)

func main() {

	var err error
	mw := new(view.SysMainWin)
	sysWin := view.InitSysMainWin(mw)

	err = sysWin.Create()
	if err != nil {
		log.Panic(err)
	}
	mw.Run()
}
