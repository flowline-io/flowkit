package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/flowline-io/flowkit/internal/assets"
	"github.com/flowline-io/flowkit/internal/pkg/constant"
	"github.com/flowline-io/flowkit/internal/pkg/logs"
	"github.com/flowline-io/flowkit/internal/pkg/server"
	"github.com/flowline-io/flowkit/internal/pkg/setting"
	"github.com/flowline-io/flowkit/internal/pkg/theme"
	"github.com/flowline-io/flowkit/internal/pkg/util"
	"github.com/flowline-io/flowkit/internal/pkg/wb"
	"github.com/flowline-io/flowkit/internal/ruleset/agent"
	"github.com/flowline-io/flowkit/internal/ruleset/instruct"
	"github.com/flowline-io/flowkit/internal/ui"
)

func main() {
	// app
	a := app.NewWithID(constant.AppId)
	assets.SetIcon(a)
	w := a.NewWindow(constant.AppTitle)

	// load preferences
	setting.LoadPreferences(a.Preferences())

	// logger
	logs.Init()

	// check singleton
	util.CheckSingleton()

	// embed server
	server.EmbedServer(constant.EmbedServerPort)

	// websocket
	wb.Init(a, w) // todo app.context

	// cron
	instruct.Cron(a, w)
	agent.Cron(a, w)

	// theme
	t := theme.NewAppTheme()
	a.Settings().SetTheme(t)

	// systray
	if desk, ok := a.(desktop.App); ok {
		ui.SetupSystray(desk, w)
	}

	// main window
	w.SetContent(ui.Create(a, w))
	w.Resize(fyne.NewSize(1000, 600))
	w.SetMaster()
	w.ShowAndRun()
}
