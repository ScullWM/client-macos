package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

type AuthCredential struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// handleMessages handles messages
func handleMessages(w *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "user.auth":

		var a AuthCredential
		_ = json.Unmarshal(m.Payload, &a)

		hash := md5.Sum([]byte(a.Email))
		hashString := hex.EncodeToString(hash[:])

		w.Send(bootstrap.MessageOut{Name: "list", Payload: hashString})
	}
	return
}
