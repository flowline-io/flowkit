package client

import (
	"fmt"
	"github.com/flowline-io/flowkit/internal/pkg/constant"
	"github.com/flowline-io/flowkit/internal/pkg/preferences"
	"github.com/flowline-io/flowkit/internal/pkg/types"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"net/http"
	"time"
)

type Flowbot struct {
	c           *resty.Client
	accessToken string
}

func NewFlowbot(accessToken string) *Flowbot {
	v := &Flowbot{accessToken: accessToken}

	v.c = resty.New()
	v.c.SetBaseURL(preferences.AppConfig().ServerHost)
	v.c.SetTimeout(time.Minute)

	return v
}

func (v *Flowbot) fetcher(action types.Action, content any) ([]byte, error) {
	resp, err := v.c.R().
		SetAuthToken(v.accessToken).
		SetResult(&types.Response{}).
		SetBody(map[string]any{
			"action":  action,
			"version": constant.ApiVersion,
			"content": content,
		}).
		Post("/flowkit")
	if err != nil {
		return nil, err
	}

	if resp.StatusCode() == http.StatusOK {
		r := resp.Result().(*types.Response)
		return jsoniter.Marshal(r.Data)
	} else {
		return nil, fmt.Errorf("%d", resp.StatusCode())
	}
}

func (v *Flowbot) Bots() (*BotsResult, error) {
	data, err := v.fetcher(constant.Bots, nil)
	if err != nil {
		return nil, err
	}
	var r BotsResult
	err = jsoniter.Unmarshal(data, &r.Bots)
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

func (v *Flowbot) Help() (*HelpResult, error) {
	data, err := v.fetcher(constant.Help, nil)
	if err != nil {
		return nil, err
	}
	var r HelpResult
	err = jsoniter.Unmarshal(data, &r.Bots)
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

func (v *Flowbot) Pull() (*InstructResult, error) {
	data, err := v.fetcher(constant.Pull, nil)
	if err != nil {
		return nil, err
	}
	var r InstructResult
	err = jsoniter.Unmarshal(data, &r.Instruct)
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

func (v *Flowbot) Agent(content types.AgentContent) (string, error) {
	data, err := v.fetcher(constant.Agent, content)
	if err != nil {
		return "", err
	}
	return string(data), err
}
