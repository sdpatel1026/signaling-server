package controllers

import (
	"encoding/json"
	"net/http"
	"signaling-server/configs"
	"signaling-server/helpers"

	"github.com/gorilla/websocket"
)

var AllRoom Rooms
var dataChannel = make(chan Data)

type Data struct {
	Message map[string]interface{}
	RoomID  string
	Client  *websocket.Conn
}

func CreateRoom(w http.ResponseWriter, r *http.Request) {
	helpers.Logger.Info("Got a request for the room creation.")
	roomID := AllRoom.CreateRoom()
	response := Response{}
	response.Result = map[string]interface{}{}
	response.Result[configs.KEY_ROOM_ID] = roomID
	err := json.NewEncoder(w).Encode(response)
	if err != nil {

		helpers.Logger.Errorf("error in encoding create room response:%s", err.Error())
	} else {
		helpers.Logger.Infof("room is succefully created with room_id %s", roomID)
	}
}

func JoinRoom(w http.ResponseWriter, r *http.Request) {
	helpers.Logger.Info("Got a request for joining into the room.")
	response := Response{}
	response.Result = make(map[string]interface{})
	roomIDs, ok := r.URL.Query()[configs.KEY_ROOM_ID]

	if !ok {
		response.Result[configs.KEY_ERROR] = configs.ERROR_EMPTY_ROOM_ID
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}
	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		helpers.Logger.Errorf("error in upgrading http connnection to the Web Socket:%s", err.Error())
		response.Result[configs.KEY_ERROR] = configs.SERVER_ERROR
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}
	roomID := roomIDs[0]
	room := AllRoom.GetRoom(roomID)
	member := Member{IsHost: false, Conn: ws}
	memberID := room.InsertIntoRoom(roomID, &member)
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}
