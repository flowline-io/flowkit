package info

import (
	"fmt"
	"github.com/flowline-io/flowkit/internal/pkg/flog"
	"github.com/flowline-io/flowkit/internal/pkg/setting"
	"github.com/gen2brain/beeep"
	"math"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
)

type clockLayout struct {
	hour, minute, second     *canvas.Line
	pips                     [12]*canvas.Line
	hourDot, secondDot, face *canvas.Circle

	canvas fyne.CanvasObject
	stop   bool
}

func (c *clockLayout) rotate(hand *canvas.Line, middle fyne.Position, facePosition float64, offset, length float32) {
	rotation := math.Pi * 2 / 60 * facePosition
	x2 := length * float32(math.Sin(rotation))
	y2 := -length * float32(math.Cos(rotation))

	offX := float32(0)
	offY := float32(0)
	if offset > 0 {
		offX += offset * float32(math.Sin(rotation))
		offY += -offset * float32(math.Cos(rotation))
	}

	hand.Position1 = fyne.NewPos(middle.X+offX, middle.Y+offY)
	hand.Position2 = fyne.NewPos(middle.X+offX+x2, middle.Y+offY+y2)
}

func (c *clockLayout) Layout(_ []fyne.CanvasObject, size fyne.Size) {
	diameter := fyne.Min(size.Width, size.Height)
	radius := diameter / 2
	dotRadius := radius / 12
	smallDotRadius := dotRadius / 8

	stroke := diameter / 40
	midStroke := diameter / 90
	smallStroke := diameter / 200

	size = fyne.NewSize(diameter, diameter)
	middle := fyne.NewPos(size.Width/2, size.Height/2)
	topleft := fyne.NewPos(middle.X-radius, middle.Y-radius)

	c.face.Resize(size)
	c.face.Move(topleft)

	c.hour.StrokeWidth = stroke
	c.rotate(c.hour, middle, float64((time.Now().Hour()%12)*5)+(float64(time.Now().Minute())/12), dotRadius, radius/2)
	c.minute.StrokeWidth = midStroke
	c.rotate(c.minute, middle, float64(time.Now().Minute())+(float64(time.Now().Second())/60), dotRadius, radius*.9)
	c.second.StrokeWidth = smallStroke
	c.rotate(c.second, middle, float64(time.Now().Second()), 0, radius-3)

	c.hourDot.StrokeWidth = stroke
	c.hourDot.Resize(fyne.NewSize(dotRadius*2, dotRadius*2))
	c.hourDot.Move(fyne.NewPos(middle.X-dotRadius, middle.Y-dotRadius))
	c.secondDot.StrokeWidth = smallStroke
	c.secondDot.Resize(fyne.NewSize(smallDotRadius*2, smallDotRadius*2))
	c.secondDot.Move(fyne.NewPos(middle.X-smallDotRadius, middle.Y-smallDotRadius))
	c.face.StrokeWidth = smallStroke

	for i, p := range c.pips {
		c.rotate(p, middle, float64((i)*5), radius/8*7, radius/8)
		p.StrokeWidth = smallStroke
	}
}

func (c *clockLayout) MinSize(_ []fyne.CanvasObject) fyne.Size {
	return fyne.NewSize(200, 200)
}

func (c *clockLayout) render() *fyne.Container {
	c.hourDot = &canvas.Circle{StrokeColor: theme.ForegroundColor(), StrokeWidth: 5}
	c.secondDot = &canvas.Circle{StrokeColor: theme.PrimaryColor(), StrokeWidth: 3}
	c.face = &canvas.Circle{StrokeColor: theme.DisabledColor(), StrokeWidth: 1}

	c.hour = &canvas.Line{StrokeColor: theme.ForegroundColor(), StrokeWidth: 5}
	c.minute = &canvas.Line{StrokeColor: theme.ForegroundColor(), StrokeWidth: 3}
	c.second = &canvas.Line{StrokeColor: theme.PrimaryColor(), StrokeWidth: 1}

	container := container.NewWithoutLayout(c.hourDot, c.secondDot)
	for i := range c.pips {
		pip := &canvas.Line{StrokeColor: theme.DisabledColor(), StrokeWidth: 1}
		container.Add(pip)
		c.pips[i] = pip
	}
	container.Objects = append(container.Objects, c.face, c.hour, c.minute, c.second)
	container.Layout = c

	c.canvas = container
	return container
}

func (c *clockLayout) animate(co fyne.CanvasObject) {
	tick := time.NewTicker(time.Second)
	go func() {
		for !c.stop {
			c.Layout(nil, co.Size())
			c.hour.Refresh()
			c.minute.Refresh()
			c.second.Refresh()
			<-tick.C
		}
	}()
}

func (c *clockLayout) applyTheme(_ fyne.Settings) {
	c.hourDot.StrokeColor = theme.ForegroundColor()
	c.secondDot.StrokeColor = theme.PrimaryColor()
	c.face.StrokeColor = theme.DisabledColor()

	c.hour.StrokeColor = theme.ForegroundColor()
	c.minute.StrokeColor = theme.ForegroundColor()
	c.second.StrokeColor = theme.PrimaryColor()

	for _, p := range c.pips {
		p.StrokeColor = theme.DisabledColor()
	}
}

// Show loads a clock example window for the specified app context
func Show(win fyne.Window) fyne.CanvasObject {
	clock := &clockLayout{}
	//clockWindow.SetOnClosed(func() {
	//	clock.stop = true
	//})

	content := clock.render()
	go clock.animate(content)

	listener := make(chan fyne.Settings)
	fyne.CurrentApp().Settings().AddChangeListener(listener)
	go func() {
		for {
			settings := <-listener
			clock.applyTheme(settings)
		}
	}()

	return content
}

func f() string {
	// store
	store, err := setting.NewConfigStore()
	if err != nil {
		flog.Error(err)
		return ""
	}

	c, err := store.Config()
	if err != nil {
		flog.Error(err)
		return ""
	}
	flog.Info(fmt.Sprintf("%v", c))

	err = store.Save(setting.Config{
		ServerHost: time.Now().String(),
	})
	if err != nil {
		flog.Error(err)
		return ""
	}

	// notify
	err = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		flog.Error(err)
		return ""
	}
	err = beeep.Notify("Title", "Message body", "")
	if err != nil {
		panic(err)
	}

	return ""
}
