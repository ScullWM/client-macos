package main

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

type Auth struct {
	Email string `json:"email"`
	Hash  string `json:"hash"`
}

type HashQuery struct {
	Email string `json:"email"`
	Hash  string `json:"hashKey"`
}

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (payload interface{}, err error) {
	switch m.Name {
	case "hash":
		// Unmarshal payload
		var hashQuery HashQuery
		if len(m.Payload) > 0 {
			// Unmarshal payload
			if err = json.Unmarshal(m.Payload, &hashQuery); err != nil {
				payload = err.Error()
				return
			}
		}

		// Explore
		if payload, err = hash(hashQuery); err != nil {
			payload = err.Error()
			return
		}

		// start watching directory
		go watchUploadFolder()
	}
	return
}

func hash(e HashQuery) (a Auth, err error) {
	data := []byte(e.Email)
	dataHash := md5.Sum(data)

	hashString := hex.EncodeToString(dataHash[:])
	updemiaUser = Auth{Email: e.Email, Hash: hashString}

	return updemiaUser, nil
}
