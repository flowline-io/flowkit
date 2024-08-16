package bots

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/flowline-io/flowkit/internal/pkg/client"
	"github.com/flowline-io/flowkit/internal/pkg/constant"
	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/flowline-io/flowkit/internal/pkg/preferences"
	jsoniter "github.com/json-iterator/go"
)

func Show(app fyne.App, win fyne.Window) fyne.CanvasObject {
	return container.NewGridWrap(fyne.NewSquareSize(100), botsBox(app)...)
}

func botsBox(app fyne.App) []fyne.CanvasObject {
	//c := client.NewFlowbot(preferences.AppConfig().AccessToken)
	//result, err := c.Bots()
	//if err != nil {
	//	flog.Error(err)
	//	return nil
	//}

	result := client.BotsResult{
		Bots: []struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		}{{"A1", "A1"}, {"A2", "A2"},
			{"A3", "A3"}, {"A4", "A4"},
			{"A5", "A5"}, {"A6", "A6"}},
	}

	list := make([]fyne.CanvasObject, 0, len(result.Bots))
	for _, bot := range result.Bots {
		list = append(list, botBox(app, bot.Id, bot.Name))
	}
	return list
}

func botBox(app fyne.App, id, name string) fyne.CanvasObject {
	check := widget.NewCheck("Enable", func(b bool) {
		flog.Info("bot %s enable %v", id, b)

		kv := preferences.AppConfig().InstructSwitch
		kv[id] = fmt.Sprintf("%v", b)

		j, err := jsoniter.Marshal(kv)
		if err != nil {
			flog.Error(err)
			return
		}
		app.Preferences().SetString(constant.InstructPreferenceKey, string(j))
		err = preferences.Load(app)
		if err != nil {
			flog.Error(err)
		}
	})
	b, ok := preferences.AppConfig().InstructSwitch.String(id)
	if ok {
		check.SetChecked(b == "true")
	}
	c := container.NewVBox(
		widget.NewLabel(name),
		check,
	)
	return c
}
