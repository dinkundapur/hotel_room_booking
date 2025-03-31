package utils

import (
	"hotel-room-booking/app/dto"
	room_service "hotel-room-booking/app/services/room"
	"math/rand"
	"reflect"
)

// this functions accepts the struct object and call all the exported methods defined in the structs
func ExecuteMethods(obj interface{}) {
	// Calling all the struct methods
	structValue := reflect.ValueOf(obj)
	for i := 0; i < structValue.NumMethod(); i++ {
		method := structValue.Method(i)
		method.Call(nil)
	}
}

const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func GenerateRandomAlphaNum() string {
	b := make([]byte, 16)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func InitializeRooms() {
	//Initialize rooms
	rooms := []dto.Room{
		{RoomNumber: 1111, Status: "Available"},
		{RoomNumber: 1112, Status: "Available"},
		{RoomNumber: 1113, Status: "Available"},
		{RoomNumber: 2222, Status: "Available"},
		{RoomNumber: 2223, Status: "Available"},
		{RoomNumber: 2224, Status: "Available"},
	}

	room_service := room_service.NewRoomService()
	room_service.InitializeRooms(rooms)
}
