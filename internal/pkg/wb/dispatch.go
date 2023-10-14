package wb

import (
	"fmt"
	"github.com/flowline-io/flowkit/internal/pkg/types"
)

func (s *Session) dispatch(msg *types.ServerComMessage) {
	fmt.Println(msg)
	// todo msg.data to instruct
	// todo instruct.RunInstruct(app, window, cache, item)
}
