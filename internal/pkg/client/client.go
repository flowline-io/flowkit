package client

import (
	"encoding/json"
	"fmt"
	"github.com/flowline-io/flowkit/internal/pkg/constant"
	"github.com/flowline-io/flowkit/internal/pkg/setting"
	"github.com/flowline-io/flowkit/internal/pkg/types"
	"github.com/flowline-io/flowkit/internal/pkg/util"
	"github.com/go-resty/resty/v2"
	"net/http"
	"time"
)

type Tinode struct {
	c           *resty.Client
	accessToken string
}

func NewTinode(accessToken string) *Tinode {
	v := &Tinode{accessToken: accessToken}

	v.c = resty.New()
	v.c.SetBaseURL(util.FillScheme(setting.DefaultConfig().ServerHost))
	v.c.SetTimeout(time.Minute)

	return v
}

func (v *Tinode) fetcher(action types.Action, content any) ([]byte, error) {
	resp, err := v.c.R().
		SetHeader("Authorization", fmt.Sprintf("Bearer %s", v.accessToken)).
		SetResult(&types.ServerComMessage{}).
		SetBody(map[string]any{
			"action":  action,
			"version": 1,
			"content": content,
		}).
		Post("/flowkit")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		r := resp.Result().(*types.ServerComMessage)
		return json.Marshal(r.Data)
	} else {
		return nil, fmt.Errorf("%d, %s (%s)",
			resp.StatusCode(),
			resp.Header().Get("X-Error-Code"),
			resp.Header().Get("X-Error"))
	}
}

func (v *Tinode) Bots() (*BotsResult, error) {
	data, err := v.fetcher(constant.Bots, nil)
	if err != nil {
		return nil, err
	}
	var r BotsResult
	err = json.Unmarshal(data, &r.Bots)
	if err != nil {
		return nil, err
	}
	return &r, err
}

type BotsResult struct {
	Bots []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"bots"`
}

func (v *Tinode) Help() (*HelpResult, error) {
	data, err := v.fetcher(constant.Help, nil)
	if err != nil {
		return nil, err
	}
	var r HelpResult
	err = json.Unmarshal(data, &r.Bots)
	if err != nil {
		return nil, err
	}
	return &r, err
}

type HelpResult struct {
	Bots []struct {
		Id   string `json:"id"`
		Name string `json:"name"`
	} `json:"bots"`
}

func (v *Tinode) Pull() (*InstructResult, error) {
	data, err := v.fetcher(constant.Pull, nil)
	if err != nil {
		return nil, err
	}
	var r InstructResult
	err = json.Unmarshal(data, &r.Instruct)
	if err != nil {
		return nil, err
	}
	return &r, err
}

type InstructResult struct {
	Instruct []Instruct `json:"instruct"`
}

type Instruct struct {
	No       string `json:"no"`
	Bot      string `json:"bot"`
	Flag     string `json:"flag"`
	Content  any    `json:"content"`
	ExpireAt string `json:"expire_at"`
}

func (v *Tinode) Agent(content types.AgentContent) (string, error) {
	data, err := v.fetcher(constant.Agent, content)
	if err != nil {
		return "", err
	}
	return string(data), err
}
