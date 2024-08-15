package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/flowline-io/flowkit/internal/pkg/constant"
	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/flowline-io/flowkit/internal/pkg/server"
	"github.com/flowline-io/flowkit/internal/pkg/setting"
	"github.com/flowline-io/flowkit/internal/pkg/util"
	"github.com/flowline-io/flowkit/internal/ruleset/agent"
	"github.com/flowline-io/flowkit/internal/ruleset/instruct"
	"github.com/flowline-io/flowkit/internal/ui/dashboard"
	"github.com/flowline-io/flowkit/internal/ui/info"
)

type appInfo struct {
	name string
	icon fyne.Resource
	run  func(fyne.Window) fyne.CanvasObject
}

var apps = []appInfo{
	{"Dashboard", theme.DocumentIcon(), dashboard.Show},
	{"Bots", theme.AccountIcon(), info.Show},
	{"Settings", theme.SettingsIcon(), info.Show},
	{"Info", theme.InfoIcon(), info.Show},
}

func main() {

	// load preferences
	err := setting.Init()
	if err != nil {
		flog.Panic(err.Error())
	}

	// logger
	flog.Init()

	// check singleton
	util.CheckSingleton()

	// embed server
	server.EmbedServer(constant.EmbedServerPort)

	// cron
	instruct.Cron()
	agent.Cron()

	a := app.New()
	a.SetIcon(resourceIconPng)

	content := container.NewStack()
	w := a.NewWindow(constant.AppName)

	appList := widget.NewList(
		func() int {
			return len(apps)
		},
		func() fyne.CanvasObject {
			icon := &canvas.Image{}
			label := widget.NewLabel("Text Editor")
			labelHeight := label.MinSize().Height
			icon.SetMinSize(fyne.NewSize(labelHeight, labelHeight))
			return container.NewBorder(nil, nil, icon, nil, label)
		},
		func(id widget.ListItemID, obj fyne.CanvasObject) {
			img := obj.(*fyne.Container).Objects[1].(*canvas.Image)
			text := obj.(*fyne.Container).Objects[0].(*widget.Label)
			img.Resource = apps[id].icon
			img.Refresh()
			text.SetText(apps[id].name)
		})
	appList.OnSelected = func(id widget.ListItemID) {
		content.Objects = []fyne.CanvasObject{apps[id].run(w)}
	}

	split := container.NewHSplit(appList, content)
	split.Offset = 0.1
	w.SetContent(split)
	w.Resize(fyne.NewSize(640, 720))
	w.SetCloseIntercept(func() {
		w.Hide()
	})

	// set system tray
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu(constant.AppName,
			fyne.NewMenuItem("Show", func() { w.Show() }),
			fyne.NewMenuItem("Quit", func() { a.Quit() }),
		)
		desk.SetSystemTrayMenu(m)
	}

	w.ShowAndRun()
}
