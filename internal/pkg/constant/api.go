package constant

import "github.com/flowline-io/flowkit/internal/pkg/types"

const ApiVersion = 1

const (
	Info  types.Action = "info"
	Pull  types.Action = "pull"
	Agent types.Action = "agent"
	Bots  types.Action = "bots"
	Help  types.Action = "help"
	Ack   types.Action = "ack"
)
