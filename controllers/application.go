package controllers

import "signaling-server/helpers"

type Response struct {
	Result map[string]interface{}
}

func brodCaster() {
	for {
		data := <-dataChannel
		room := AllRoom.GetRoom(data.RoomID)
		for memberID, member := range room.Map {
			if member.Conn != data.Client {
				err := member.Conn.WriteJSON(data.Message)
				if err != nil {
					helpers.Logger.Errorf("error in sending message to member with member_id %s is %s", memberID, err.Error())
					member.Conn.Close()
				}
			}
		}
	}
}
