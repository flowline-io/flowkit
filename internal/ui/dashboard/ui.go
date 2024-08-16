package dashboard

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"github.com/flowline-io/flowkit/internal/pkg/constant"
	"image/color"
	"time"
)

func Show(app fyne.App, win fyne.Window) fyne.CanvasObject {
	app.Preferences().SetBool(constant.InstructPreferenceKey, true)
	fmt.Printf("App: %+v\n", app.Metadata())

	// app.SendNotification(&fyne.Notification{Title: "Hello", Content: "This is a notification"})

	str := binding.NewString()
	go func() {
		dots := "....."
		for i := 5; i >= 0; i-- {
			_ = str.Set("Count down" + dots[:i])
			time.Sleep(time.Second)
			_ = str.Set("Blast off")
		}
	}()

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText(fmt.Sprintf("%v", app.Preferences().Bool(constant.InstructPreferenceKey)), green)
	text2.Move(fyne.NewPos(20, 20))
	content := container.NewWithoutLayout(text1, text2, widget.NewLabelWithData(str))

	return content
}
