package main

import (
	"github.com/murlokswarm/app"
	"github.com/murlokswarm/log"
)

// AppMainMenu implements app.Componer interface.
type AppMainMenu struct {
	CustomTitle string
	Disabled    bool
}

func (m *AppMainMenu) Render() string {
	return `
<menu>
    <menu label="app">
        <menuitem label="Quit" shortcut="meta+q" selector="terminate:" />
    </menu>
    <WindowMenu />
</menu>
    `
}

// OnCustomMenuClick is the handler called when an onclick event occurs in a menuitem.
func (m *AppMainMenu) OnCustomMenuClick() {
	log.Info("OnCustomMenuClick")
}

// WindowMenu implements app.Componer interface.
// It's another component which will be nested inside the AppMenu component.
type WindowMenu struct {
}

func (m *WindowMenu) Render() string {
	return `
<menu label="Window">
    <menuitem label="Close" selector="performClose:" shortcut="meta+w" />
</menu>
    `
}

func init() {
	app.RegisterComponent(&AppMainMenu{})
	app.RegisterComponent(&WindowMenu{})
}
