package main

import (
	"flag"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Vars
var (
	AppName string
	BuiltAt string
	debug   = flag.Bool("d", false, "if yes, the app is in debug mode")
	window  *astilectron.Window
)

var updemiaUser *Updemia

type Updemia struct {
	Email      string
	Passphrase string
	Hash       string
	Imgs       []UpdemiaImg
}

type UpdemiaImg struct {
	Key string
	Url string
	Img string
}

type ImgsResponse struct {
	Collection []UpdemiaImg
}

func main() {
	// Init
	flag.Parse()
	astilog.FlagInit()

	// Run bootstrap
	if err := bootstrap.Run(bootstrap.Options{
		Asset: Asset,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/logo.icns",
			AppIconDefaultPath: "resources/logo.png",
		},
		Debug:    *debug,
		Homepage: "index.html",
		MenuOptions: []*astilectron.MenuItemOptions{
			{
				Label: astilectron.PtrStr(AppName),
				SubMenu: []*astilectron.MenuItemOptions{
					{
						Role: astilectron.MenuItemRoleClose,
					},
				},
			},
			{
				Label: astilectron.PtrStr("Style"),
				SubMenu: []*astilectron.MenuItemOptions{
					{
						Checked: astilectron.PtrBool(true),
						Label:   astilectron.PtrStr("Dark"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							// Send
							if err := window.Send(bootstrap.MessageOut{Name: "set.style", Payload: "dark"}); err != nil {
								astilog.Error(errors.Wrap(err, "setting dark style failed"))
								return
							}
							return
						},
						Type: astilectron.MenuItemTypeRadio,
					},
					{
						Label: astilectron.PtrStr("Light"),
						OnClick: func(e astilectron.Event) (deleteListener bool) {
							// Send
							if err := window.Send(bootstrap.MessageOut{Name: "set.style", Payload: "light"}); err != nil {
								astilog.Error(errors.Wrap(err, "setting dark style failed"))
								return
							}
							return
						},
						Type: astilectron.MenuItemTypeRadio,
					},
				},
			},
		},
		MessageHandler: handleMessages,
		OnWait: func(_ *astilectron.Astilectron, w *astilectron.Window, _ *astilectron.Menu, t *astilectron.Tray, _ *astilectron.Menu) error {
			// Store global variables
			window = w

			// Add listeners on tray
			t.On(astilectron.EventNameTrayEventClicked, func(e astilectron.Event) (deleteListener bool) { astilog.Info("Tray has been clicked!"); return })
			return nil
		},
		RestoreAssets: RestoreAssets,
		TrayOptions: &astilectron.TrayOptions{
			Image:   astilectron.PtrStr("resources/tray.png"),
			Tooltip: astilectron.PtrStr("Wow, what a beautiful tray!"),
		},
		WindowOptions: &astilectron.WindowOptions{
			BackgroundColor: astilectron.PtrStr("#333"),
			Center:          astilectron.PtrBool(true),
			Height:          astilectron.PtrInt(720),
			Width:           astilectron.PtrInt(370),
		},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}
