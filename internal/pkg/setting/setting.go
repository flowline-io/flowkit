package setting

import (
	"github.com/flowline-io/flowkit/internal/pkg/types"
)

var s = &Setting{}

type Setting struct {
	ServerHost      string
	LogPath         string
	AccessToken     string
	RequestInterval int
	InstructSwitch  types.KV
}

func LoadPreferences(_ any) {
	//s.ServerHost = p.String(constant.ServerPreferenceKey)
	//s.LogPath = p.String(constant.LogPreferenceKey)
	//s.AccessToken = p.String(constant.TokenPreferenceKey)
	//s.RequestInterval = p.Int(constant.IntervalPreferenceKey)

	//data := p.String(constant.InstructPreferenceKey)
	//instructSwitch := types.KV{}
	//_ = json.Unmarshal([]byte(data), &instructSwitch)
	//s.InstructSwitch = instructSwitch
}

func Get() *Setting {
	return s
}
