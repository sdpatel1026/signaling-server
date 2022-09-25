package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"signaling-server/controllers"
	"signaling-server/helpers"
)

func main() {
	helpers.Logger.Info("signaling server start...")
	log.Println("server start")
	controllers.AllRoom = new(controllers.Rooms)
	controllers.AllRoom.New()
	http.HandleFunc("/room", controllers.CreateRoom)
	http.HandleFunc("/join", controllers.JoinRoom)
	// appPort := configs.GetEnvWithKey(configs.KEY_APP_PORT, "8080")
	appPort := os.Getenv("PORT")
	severAddres := fmt.Sprintf(":%s", appPort)
	fmt.Println("port:", severAddres)
	err := http.ListenAndServe(severAddres, nil)
	if err != nil {
		log.Fatal(err)
	}
}
