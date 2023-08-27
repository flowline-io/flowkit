package bot

import (
	"fyne.io/fyne/v2"
	"github.com/flowline-io/linkit/internal/pkg/types"
)

type Executor struct {
	Flag string
	Run  func(app fyne.App, window fyne.Window, data types.KV) error
}

var DoInstruct = map[string][]Executor{
	"dev":       dev,
	"clipboard": clipboard,
	"url":       url,
}
