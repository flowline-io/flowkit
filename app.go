package main

import (
	"context"
	"fmt"
	"github.com/flowline-io/flowkit/internal/pkg/setting"
	"github.com/gen2brain/beeep"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"time"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// beforeClose is a function that is called before closing the App.
//
// It displays a message dialog asking the user if they want to quit.
// The function takes a context.Context parameter.
// It returns a bool indicating whether the close action should be prevented.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	dialog, err := runtime.MessageDialog(ctx, runtime.MessageDialogOptions{
		Type:    runtime.QuestionDialog,
		Title:   "Quit?",
		Message: "Are you sure you want to quit?",
		Buttons: []string{"Ok", "Cancel"},
	})

	if err != nil {
		return false
	}
	return dialog != "Ok"
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {

	// store
	store, err := setting.NewConfigStore()
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return ""
	}

	c, err := store.Config()
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return ""
	}
	runtime.LogPrintf(a.ctx, "%+v", c)

	err = store.Save(setting.Config{
		ServerHost: time.Now().String(),
	})
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return ""
	}

	// notify
	err = beeep.Beep(beeep.DefaultFreq, beeep.DefaultDuration)
	if err != nil {
		runtime.LogError(a.ctx, err.Error())
		return ""
	}
	err = beeep.Notify("Title", "Message body", "")
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a App) Example(p Person) error {
	fmt.Println(p)
	return nil
}

type Person struct {
	Name    string   `json:"name"`
	Age     uint8    `json:"age"`
	Address *Address `json:"address"`
}

type Address struct {
	Street   string `json:"street"`
	Postcode string `json:"postcode"`
}
