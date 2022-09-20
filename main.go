package main

import (
	"fmt"
	"log"
	"net/http"
	"signaling-server/configs"
	"signaling-server/helpers"
)

func main() {
	helpers.Logger.Info("signaling server start...")

	http.HandleFunc("/create", CreateRoom)
	http.HandleFunc("/join", JoinRoom)
	appPort := configs.GetEnvWithKey(configs.KEY_APP_PORT, "")
	severAddres := fmt.Sprintf(":%s", appPort)
	err := http.ListenAndServe(severAddres, nil)
	if err != nil {
		log.Fatal(err)
	}
}
