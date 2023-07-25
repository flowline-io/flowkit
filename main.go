package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/sysatom/linkit/internal/assets"
	"github.com/sysatom/linkit/internal/pkg/theme"
	"github.com/sysatom/linkit/internal/ruleset/agent"
	"github.com/sysatom/linkit/internal/ruleset/instruct"
	"github.com/sysatom/linkit/internal/ui"
)

func main() {
	// app
	a := app.NewWithID("com.github.sysatom.linkit")
	assets.SetIcon(a)
	w := a.NewWindow("Linkit")

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
