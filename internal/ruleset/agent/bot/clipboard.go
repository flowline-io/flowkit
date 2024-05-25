package bot

import (
	"github.com/allegro/bigcache/v3"
	"github.com/flowline-io/flowkit/internal/pkg/client"
	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/flowline-io/flowkit/internal/pkg/types"
)

const (
	ClipboardAgentVersion = 1
	UploadAgentID         = "clipboard_upload"
)

func ClipboardUpload(cache *bigcache.BigCache, c *client.Flowbot) {
	old, _ := cache.Get("clipboard")
	now := "..." // todo
	if string(old) == now {
		return
	}
	_, err := c.Agent(types.AgentContent{
		Id:      UploadAgentID,
		Version: ClipboardAgentVersion,
		Content: types.KV{
			"txt": now,
		},
	})
	if err != nil {
		flog.Error(err)
	}
	_ = cache.Set("clipboard", []byte(now))
}
