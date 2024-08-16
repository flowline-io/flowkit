package info

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/flowline-io/flowkit/internal/pkg/constant"
	"github.com/flowline-io/flowkit/version"
)

func Show(app fyne.App, win fyne.Window) fyne.CanvasObject {
	return container.NewVBox(
		widget.NewLabel(constant.AppName),
		widget.NewLabel(constant.AppId),
		widget.NewLabel(fmt.Sprintf("%s (%s)", version.Buildtags, version.Buildstamp)),
	)
}
