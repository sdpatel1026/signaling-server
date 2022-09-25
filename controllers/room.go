package controllers

import (
	"fmt"
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

//Member represent single client.
type Member struct {
	Mutex         sync.RWMutex
	IsHost        bool
	WebSocketConn *websocket.Conn
}

//Room has different members, each have unique memberID.
type Room struct {
	Mutex sync.RWMutex
	Map   map[string]*Member
}

//Rooms contain multiple Room with unique roomID.
type Rooms struct {
	Mutex sync.RWMutex
	Map   map[string]*Room
}

//New initialise Rooms.
func (rooms *Rooms) New() {
	rooms.Map = make(map[string]*Room)
}

//New initialise Room.
func (room *Room) New() {
	room.Map = make(map[string]*Member)
}

//Get will return Room associate with given roomID.
func (rooms *Rooms) GetRoom(roomID string) *Room {
	rooms.Mutex.RLock()
	defer rooms.Mutex.RUnlock()
	fmt.Printf("roomID: %v\n", roomID)
	fmt.Printf("rooms.Map[roomID]: %v\n", rooms.Map[roomID])
	return rooms.Map[roomID]
}

//CreateRoom create room and put  it into the Rooms.
func (rooms *Rooms) CreateRoom() string {
	rooms.Mutex.Lock()
	defer rooms.Mutex.Unlock()
	roomID := uuid.New().String()
	room := &Room{}
	room.New()
	rooms.Map[roomID] = room
	fmt.Printf("rooms: %v\n", rooms)
	return roomID
}

//InsertIntoRoom insert member into the Room based on roomID.
func (room *Room) InsertIntoRoom(roomID string, member *Member) (memberID string) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()
	memberID = uuid.New().String()
	room.Map[memberID] = member
	return memberID
}

//DeleteMember delete member from the Room.
func (room *Room) DeleteMember(memberID string) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()
	delete(room.Map, memberID)
}

//deleteRoom deletes the Room from Rooms.
func (rooms *Rooms) deleteRoom(roomID string) {
	rooms.Mutex.Lock()
	defer rooms.Mutex.Unlock()
	delete(rooms.Map, roomID)
}
