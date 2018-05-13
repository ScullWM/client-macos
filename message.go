package main

import (
	"crypto/md5"
	"encoding/json"
	"encoding/hex"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

type Auth struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "hash":
		// Unmarshal payload
		var email string
		if len(m.Payload) > 0 {
			// Unmarshal payload
			if err = json.Unmarshal(m.Payload, &email); err != nil {
				payload = err.Error()
				return
			}
		}

		// Explore
		if payload, err = hash(email); err != nil {
			payload = err.Error()
			return
		}

		// start watching directory
		go watchUploadFolder()
	}
	return
}

func hash(e string) (a Auth, err error) {
    data := []byte(e)
    dataHash := md5.Sum(data)

    hashString := hex.EncodeToString(dataHash[:])
    updemiaUser = Auth{Email: e, Hash: hashString}

	return updemiaUser, nil
}

// PayloadDir represents a dir payload
type Dir struct {
	Name string `json:"name"`
	Path string `json:"path"`
}
