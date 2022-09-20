package controllers

import (
	"sync"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

//Member represent single client.
type Member struct {
	IsHost bool
	Conn   *websocket.Conn
}

//Room has different members, each have unique memberID.
type Room struct {
	Mutex sync.Mutex
	Map   map[string]Member
}

//Rooms contain multiple Room with unique roomID.
type Rooms struct {
	Mutex sync.RWMutex
	Map   map[string]Room
}

//New initialise Rooms.
func (rooms *Rooms) New() {
	rooms.Map = make(map[string]Room)
}

//New initialise Room.
func (room *Room) New() {
	room.Map = make(map[string]Member)
}

//Get will return Room associate with given roomID.
func (rooms *Rooms) GetRoom(roomID string) Room {
	rooms.Mutex.RLock()
	defer rooms.Mutex.RUnlock()
	return rooms.Map[roomID]
}

func (rooms *Rooms) CreateRoom() string {
	rooms.Mutex.Lock()
	defer rooms.Mutex.Unlock()
	roomID := uuid.New().String()
	room := Room{}
	room.New()
	rooms.Map[roomID] = room
	return roomID
}

func (room *Room) InsertIntoRoom(roomID string, member *Member) (memberID string) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()
	memberID = uuid.New().String()
	room.Map[memberID] = *member
	return memberID
}
func (room *Room) DeleteMember(memberID string) {
	room.Mutex.Lock()
	defer room.Mutex.Unlock()
	delete(room.Map, memberID)
}
func (rooms *Rooms) deleteRoom(roomID string) {
	rooms.Mutex.Lock()
	defer rooms.Mutex.Unlock()
	delete(rooms.Map, roomID)
}
