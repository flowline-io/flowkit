package instruct

import (
	"fmt"
	"github.com/flowline-io/flowkit/internal/pkg/preferences"
	"time"

	"github.com/allegro/bigcache/v3"
	"github.com/flowline-io/flowkit/internal/pkg/client"
	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/flowline-io/flowkit/internal/pkg/types"
	"github.com/flowline-io/flowkit/internal/ruleset/instruct/bot"
)

type instructJob struct {
	app    any
	window any
	cache  *bigcache.BigCache
	client *client.Flowbot
}

func (j *instructJob) Run() {
	res, err := j.client.Pull()
	if err != nil {
		flog.Error(err)
		return
	}
	if res == nil {
		return
	}
	// get preference
	switcher := preferences.AppConfig().InstructSwitch
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
			flog.Error(fmt.Errorf("instruct run job failed %s %s %s", item.Bot, item.No, err))
		}
	}
}

func RunInstruct(app any, window any, cache *bigcache.BigCache, item client.Instruct) error {
	for id, dos := range bot.DoInstruct {
		if item.Bot != id {
			continue
		}
		for _, do := range dos {
			if item.Flag != do.Flag {
				continue
			}
			// run instruct
			flog.Info("[instruct] %s %s", item.Bot, item.No)
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
