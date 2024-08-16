package preferences

import (
	"encoding/json"
	"fyne.io/fyne/v2"
	"github.com/flowline-io/flowkit/internal/pkg/constant"
	"github.com/flowline-io/flowkit/internal/pkg/types"
)

type Config struct {
	ServerHost     string   `json:"server_host"`
	AccessToken    string   `json:"access_token"`
	LogPath        string   `json:"log_path"`
	InstructSwitch types.KV `json:"instruct_switch"`
}

var config Config

func AppConfig() Config {
	return config
}

func Load(app fyne.App) error {
	config.ServerHost = app.Preferences().String(constant.ServerPreferenceKey)
	config.AccessToken = app.Preferences().String(constant.TokenPreferenceKey)
	config.LogPath = app.Preferences().String(constant.LogPreferenceKey)

	data := app.Preferences().String(constant.InstructPreferenceKey)
	if data == "" {
		config.InstructSwitch = make(types.KV)
	} else {
		err := json.Unmarshal([]byte(data), &config.InstructSwitch)
		if err != nil {
			return err
		}
	}
	return nil
}
