package controllers

import (
	"signaling-server/helpers"

	"github.com/gorilla/websocket"
)

type Response struct {
	Result map[string]interface{}
}

//broadCaster broad-cast message to all member present into the room, except sender.
func brodCaster() {
	for {
		data := <-dataChannel
		helpers.Logger.Infof("Trying to brodcast the message:%v", data)
		room := AllRoom.GetRoom(data.RoomID)
		for memberID, member := range room.Map {
			member.Mutex.Lock()
			if member.WebSocketConn != data.WebSocketConn {
				err := member.WebSocketConn.WriteJSON(data.Message)
				if err != nil {
					helpers.Logger.Errorf("error in sending message to member with member_id %s is %s", memberID, err.Error())
					member.WebSocketConn.Close()
				}
			}
			member.Mutex.Unlock()
		}
	}
}

//receiver receive message from webSocket and put messafe into the dataChannet for broadCasting.
func receiver(wsConn *websocket.Conn, roomID, memberID string) {
	for {
		var data Data
		err := wsConn.ReadJSON(&data.Message)
		if err != nil {
			helpers.Logger.Errorf("error in receiving message from member_id %s is %s", memberID, err.Error())
			wsConn.Close()
			break
		}
		data.WebSocketConn = wsConn
		data.RoomID = roomID
		dataChannel <- data
	}
	//needs to add the code that will remove this client from the room and
	//break the for loop in order to  complete the execution of this thread.
	//additionally If this room has no member than we may also remove this Room from the Rooms.
}
