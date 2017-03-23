package main

import (
	"github.com/deckarep/gosx-notifier"
	"github.com/sadlil/go-trigger"
	"io"
	"log"
	"net/http"
	"os"
)

const userSuccess string = "Upload complete"
const userFail string = "Upload fail"
const notifTitle string = "On updemia.com"
const distantLogo string = "http://www.updemia.com/images/logo.png"
const localLogo string = "/tmp/m1UIjW1.png"

func notifyUserSuccess(url string) {
	note := gosxnotifier.NewNotification(userSuccess)
	note.AppIcon = localLogo
	note.Title = notifTitle
	note.Link = url
	note.Group = "com.unique.updemia.identifier"
	err := note.Push()

	if err != nil {
		log.Println("notification error")
	}

	trigger.Fire("user-newfile-success")
}

func notifyUserFail() {
	note := gosxnotifier.NewNotification(userFail)
	note.AppIcon = localLogo
	note.Title = notifTitle
	note.Group = "com.unique.updemia.identifier"
	err := note.Push()

	if err != nil {
		log.Println("notification error")
	}
}

func saveNotificationLogo() {
	img, _ := os.Create(localLogo)
	defer img.Close()

	resp, _ := http.Get(distantLogo)
	defer resp.Body.Close()

	_, err := io.Copy(img, resp.Body)
	if err != nil {
		log.Println("Error getting logo")
	}
}
