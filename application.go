package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/murlokswarm/app"
	"github.com/sadlil/go-trigger"
	"log"
	"net/http"
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

func (h *Updemia) Render() string {
	return `
<div class="WindowLayout" oncontextmenu="OnContextMenu">
    {{if .Hash}}
    <div class="Updemia">
        <img src="http://www.updemia.com/images/logo.png" alt="" />
        <img src="folder.png" alt="" onclick="OnOpenFolder" />
    </div>
    <div class="Gravatar">
        <img src="https://www.gravatar.com/avatar/{{html .Hash}}" title="logout" onclick="OnUserLogout" alt="" />
    </div>
    <div class="">
      <div class="listMedias">
      {{ range $index, $element := .Imgs }}
          <div class="media">
            <a href="{{html .Url}}" target="_blank" style="background-image: url('{{ $element.Img }}');"> </a>
          </div>
      {{ end }}
      </div>
    </div>
    {{else}}
    <div class="HelloBox center">

        <img src="http://www.updemia.com/images/logo.png" alt="" />
        <h1>
            Hello, you
        </h1>
        <input type="email"
               value="{{html .Email}}"
               placeholder="What's your user email?"
               autofocus="true"
               onkeydown="Email"
               onkeyup="Email" required="required" />
        <input type="password"
               value="{{html .Passphrase}}"
               placeholder="Pass phrase"
               onkeydown="Passphrase"
               onkeyup="Passphrase" required="required" />

         <input type="submit" value="Start" class="btn" onclick="OnUserLog" />
    </div>
    {{end}}
</div>
    `
}

func (h *Updemia) OnUserLogout(arg app.ChangeArg) {
	h.Hash = ""
	app.Render(updemiaUser)
}

func (h *Updemia) OnUserLog(arg app.ChangeArg) {

	if len(h.Email) > 0 {
		hash := md5.Sum([]byte(h.Email))
		h.Hash = hex.EncodeToString(hash[:])

		h.updateImgs()
		updemiaUser = h

		// update screenshot capture
		updateScreencaptureDirectory(getDestinationPath())

		// start watching directory
		go watchUploadFolder()
	}

	app.Render(h)

	trigger.On("user-newfile-success", func() {
		h.updateImgs()
		app.Render(updemiaUser)
	})
}

func (h *Updemia) OnOpenFolder(arg app.ChangeArg) {
	openDirectory(getDestinationPath())
}

func (h *Updemia) updateImgs() {
	url := fmt.Sprintf("http://www.updemia.com/api/v1/get?hash=%s", h.Hash)
	req, _ := http.NewRequest("GET", url, nil)
	client := &http.Client{}
	resp, _ := client.Do(req)
	defer resp.Body.Close()

	keys := make([]UpdemiaImg, 0)
	if err := json.NewDecoder(resp.Body).Decode(&keys); err != nil {
		log.Println(err)
	}

	h.Imgs = keys
}

func (h *Updemia) OnContextMenu() {
	ctxmenu := app.NewContextMenu()
	ctxmenu.Mount(&AppMainMenu{})
}

func init() {
	app.RegisterComponent(&Updemia{})
}
