package repositories

import (
	"hotel-room-booking/app/enum"
	"hotel-room-booking/app/models"
)

type IRoomRepository interface {
	InitializeRooms(rooms []models.Room)
	GetAvailableRooms() []models.Room
	UpdateRoomStatus(room_number int, status enum.RoomStatus)
}
