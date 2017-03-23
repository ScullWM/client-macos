package main

import (
	"github.com/murlokswarm/app"
	_ "github.com/murlokswarm/mac"
)

var (
	win app.Contexter
)

func main() {
	app.OnLaunch = func() {
		appMenu := &AppMainMenu{}
		app.MenuBar().Mount(appMenu)

		appMenuDock := &AppMainMenu{}
		app.Dock().Mount(appMenuDock)

		win = newMainWindow()
	}

	app.OnReopen = func(hasVisibleWindow bool) {
		if win != nil {
			return
		}

		win = newMainWindow()
	}

	defer app.Run()
}

func newMainWindow() app.Contexter {
	// Creates a window context.
	win := app.NewWindow(app.Window{
		Title:          "Updemia - Client",
		Width:          320,
		Height:         720,
		MinWidth:       320,
		MinHeight:      720,
		MaxWidth:       320,
		TitlebarHidden: true,
		OnClose: func() bool {
			win = nil
			return true
		},
	})

	hello := &Updemia{} // Creates a Hello component.
	win.Mount(hello)    // Mounts the Hello component into the window context.
	return win
}
