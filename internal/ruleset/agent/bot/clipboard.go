package bot

import (
	"fyne.io/fyne/v2"
	"github.com/allegro/bigcache/v3"
	"github.com/flowline-io/linkit/internal/pkg/client"
	"github.com/flowline-io/linkit/internal/pkg/logs"
	"github.com/flowline-io/linkit/internal/pkg/types"
)

const (
	ClipboardAgentVersion = 1
	UploadAgentID         = "clipboard_upload"
)

func ClipboardUpload(window fyne.Window, cache *bigcache.BigCache, c *client.Tinode) {
	old, _ := cache.Get("clipboard")
	now := window.Clipboard().Content()
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
		logs.Error(err)
	}
	_ = cache.Set("clipboard", []byte(now))
}
