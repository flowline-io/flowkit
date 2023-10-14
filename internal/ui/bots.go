package ui

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/flowline-io/flowkit/internal/pkg/client"
	"github.com/flowline-io/flowkit/internal/pkg/constant"
	"github.com/flowline-io/flowkit/internal/pkg/logs"
	"github.com/flowline-io/flowkit/internal/pkg/setting"
	"github.com/flowline-io/flowkit/internal/pkg/types"
)

type bots struct {
	client *client.Tinode

	app         fyne.App
	window      fyne.Window
	preferences fyne.Preferences
}

func newBotsTab(app fyne.App, window fyne.Window) *container.TabItem {
	b := bots{app: app, window: window, preferences: app.Preferences()}
	return &container.TabItem{
		Text:    "Bots",
		Icon:    theme.ComputerIcon(),
		Content: b.buildUI(),
	}
}

func (b *bots) buildUI() *container.Scroll {
	onOffOptions := []string{"On", "Off"}

	b.client = client.NewTinode(setting.Get().AccessToken)

	var options []string
	var co []fyne.CanvasObject
	co = append(co, &widget.Label{Text: "manage your bots instruct settings here."})

	res, err := b.client.Bots()
	if err != nil {
		logs.Error(err)
	}
	if res != nil {
		d := b.preferences.String(constant.InstructPreferenceKey)
		data := types.KV{}
		_ = json.Unmarshal([]byte(d), &data)
		for _, item := range res.Bots {
			selected := "Off"
			if s, ok := data.String(item.Id); ok {
				selected = s
			}
			id := item.Id
			options = append(options, item.Name)
			co = append(co, container.NewGridWithColumns(2,
				newBoldLabel(item.Name), &widget.RadioGroup{Options: onOffOptions, Horizontal: true, Required: true, Selected: selected, OnChanged: func(val string) {
					b.onSwitchChanged(id, val)
				}},
			))
		}
	}

	settingsContainer := &widget.Card{Title: "Bots Instruct Settings", Content: container.NewVBox(co...)}

	var data = [][]string{[]string{"top left", "top right"},
		[]string{"bottom left", "bottom right"}}

	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			o.(*widget.Label).SetText(data[i.Row][i.Col])
		},
	)
	infoContainer := container.NewVBox(
		widget.NewSelectEntry(options),
		&widget.Card{Title: "dev", Content: table},
	)

	panelContainer := container.NewGridWithColumns(2,
		settingsContainer, infoContainer,
	)

	return container.NewScroll(panelContainer)
}

func (b *bots) onSwitchChanged(id, val string) {
	d := b.preferences.String(constant.InstructPreferenceKey)
	data := types.KV{}
	_ = json.Unmarshal([]byte(d), &data)
	data[id] = val
	j, _ := json.Marshal(data)
	b.preferences.SetString(constant.InstructPreferenceKey, string(j))
}
