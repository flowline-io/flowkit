package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/flowline-io/flowkit/internal/pkg/constant"
	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/flowline-io/flowkit/internal/pkg/preferences"
	"github.com/flowline-io/flowkit/internal/pkg/server"
	"github.com/flowline-io/flowkit/internal/pkg/util"
	"github.com/flowline-io/flowkit/internal/ruleset/agent"
	"github.com/flowline-io/flowkit/internal/ruleset/instruct"
	"github.com/flowline-io/flowkit/internal/ui/bots"
	"github.com/flowline-io/flowkit/internal/ui/dashboard"
	"github.com/flowline-io/flowkit/internal/ui/info"
	"github.com/flowline-io/flowkit/internal/ui/settings"
)

type appInfo struct {
	name string
	icon fyne.Resource
	run  func(fyne.App, fyne.Window) fyne.CanvasObject
}

var apps = []appInfo{
	{"Dashboard", theme.BrokenImageIcon(), dashboard.Show},
	{"Bots", theme.AccountIcon(), bots.Show},
	{"Settings", theme.SettingsIcon(), settings.Show},
	{"Info", theme.InfoIcon(), info.Show},
}

func main() {
	// logger
	flog.Init()

	// check singleton
	util.CheckSingleton()

	// embed server
	server.EmbedServer(constant.EmbedServerPort)

	a := app.NewWithID(constant.AppId)
	a.SetIcon(resourceIconPng)

	// main ui
	content := container.New(layout.NewCustomPaddedLayout(20, 20, 20, 20))
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
		content.Objects = []fyne.CanvasObject{apps[id].run(a, w)}
	}
	appList.Select(0)
	split := container.NewHSplit(appList, content)
	split.Offset = 0.2
	w.CenterOnScreen()
	w.SetContent(split)
	w.Resize(fyne.NewSize(1080, 720))
	w.SetCloseIntercept(func() {
		w.Hide()
	})

	// system tray
	if desk, ok := a.(desktop.App); ok {
		m := fyne.NewMenu(constant.AppName,
			fyne.NewMenuItem("Show", func() { w.Show() }),
			fyne.NewMenuItem("Quit", func() { a.Quit() }),
		)
		desk.SetSystemTrayMenu(m)
	}

	// lifecycle hook
	a.Lifecycle().SetOnStarted(func() {
		flog.Info("app started")

		// load
		err := preferences.Load(a)
		if err != nil {
			flog.Panic(err.Error())
		}

		// cron
		instruct.Cron()
		agent.Cron()
	})

	w.ShowAndRun()
}
