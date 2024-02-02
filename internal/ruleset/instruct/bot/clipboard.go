package bot

import (
	"github.com/flowline-io/flowkit/internal/pkg/types"
)

var clipboard = []Executor{
	{
		Flag: "clipboard_share",
		Run: func(app any, window any, data types.KV) error {
			txt, _ := data.String("txt")
			if txt != "" {
				// app.SendNotification(fyne.NewNotification("clipboard", "share text from chat"))
				// window.Clipboard().SetContent(txt)
			}
			return nil
		},
	},
}
