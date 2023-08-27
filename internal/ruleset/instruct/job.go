package instruct

import (
	"fmt"
	"fyne.io/fyne/v2"
	"github.com/allegro/bigcache/v3"
	"github.com/flowline-io/linkit/internal/pkg/client"
	"github.com/flowline-io/linkit/internal/pkg/logs"
	"github.com/flowline-io/linkit/internal/pkg/setting"
	"github.com/flowline-io/linkit/internal/pkg/types"
	"github.com/flowline-io/linkit/internal/ruleset/instruct/bot"
	"time"
)

type instructJob struct {
	app    fyne.App
	window fyne.Window
	cache  *bigcache.BigCache
	client *client.Tinode
}

func (j *instructJob) Run() {
	res, err := j.client.Pull()
	if err != nil {
		logs.Error(err)
		return
	}
	if res == nil {
		return
	}
	// get preference
	switcher := setting.Get().InstructSwitch
	// instruct loop
	for _, item := range res.Instruct {
		// check switch
		s, ok := switcher.String(item.Bot)
		if !ok || s == "" || s == "Off" {
			continue
		}
		// check has been run
		has, _ := j.cache.Get(item.No)
		if len(has) > 0 {
			continue
		}
		// check expired
		expiredAt, err := time.Parse("2006-01-02T15:04:05Z", item.ExpireAt)
		if err != nil {
			continue
		}
		if time.Now().After(expiredAt) {
			continue
		}
		err = RunInstruct(j.app, j.window, j.cache, item)
		if err != nil {
			logs.Error(fmt.Errorf("instruct run job failed %s %s %s", item.Bot, item.No, err))
		}
	}
}

func RunInstruct(app fyne.App, window fyne.Window, cache *bigcache.BigCache, item client.Instruct) error {
	for id, dos := range bot.DoInstruct {
		if item.Bot != id {
			continue
		}
		for _, do := range dos {
			if item.Flag != do.Flag {
				continue
			}
			// run instruct
			logs.Info("[instruct] %s %s", item.Bot, item.No)
			data := types.KV{}
			if v, ok := item.Content.(map[string]any); ok {
				data = v
			}
			err := do.Run(app, window, data)
			if err != nil {
				return err
			}
			err = cache.Set(item.No, []byte("1"))
			if err != nil {
				return err
			}
		}
	}
	return nil
}
