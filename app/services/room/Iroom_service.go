package room_service

import (
	"hotel-room-booking/app/dto"
	"hotel-room-booking/app/enum"
	"hotel-room-booking/app/models"
)

type IRoomService interface {
	InitializeRooms(rooms []dto.Room)
	GetAvailableRooms() []models.Room
	UpdateRoomStatus(room_number int, status enum.RoomStatus)
}
