package bot

import (
	"github.com/flowline-io/flowkit/internal/pkg/client"
	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/flowline-io/flowkit/internal/pkg/types"
	"time"
)

const (
	DevAgentVersion = 1
	ImportAgentId   = "import_agent"
)

func DevImport(c *client.Tinode) {
	_, err := c.Agent(types.AgentContent{
		Id:      ImportAgentId,
		Version: DevAgentVersion,
		Content: types.KV{
			"time": time.Now().String(),
		},
	})
	if err != nil {
		flog.Error(err)
	}
}
