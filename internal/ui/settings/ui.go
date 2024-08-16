package settings

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/flowline-io/flowkit/internal/pkg/constant"
	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/flowline-io/flowkit/internal/pkg/preferences"
)

func Show(app fyne.App, win fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		hostInput(app),
		tokenInput(app),
	)
}

func hostInput(app fyne.App) fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("input api url")
	entry.SetText(preferences.AppConfig().ServerHost)
	entry.OnChanged = func(s string) {
		app.Preferences().SetString(constant.ServerPreferenceKey, s)
		err := preferences.Load(app)
		if err != nil {
			flog.Error(err)
		}
	}

	return container.NewVBox(
		widget.NewLabel("Flowbot API URL"),
		entry,
	)
}

func tokenInput(app fyne.App) fyne.CanvasObject {
	entry := widget.NewEntry()
	entry.SetPlaceHolder("input access token")
	entry.SetText(preferences.AppConfig().AccessToken)
	entry.OnChanged = func(s string) {
		app.Preferences().SetString(constant.TokenPreferenceKey, s)
		err := preferences.Load(app)
		if err != nil {
			flog.Error(err)
		}
	}

	return container.NewVBox(
		widget.NewLabel("Access token"),
		entry,
	)
}
