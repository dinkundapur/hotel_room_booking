package room_service

import (
	"hotel-room-booking/app/dto"
	"hotel-room-booking/app/enum"
	"hotel-room-booking/app/models"
	"hotel-room-booking/app/repositories"
)

type RoomService struct {
	repo repositories.IRoomRepository
}

func NewRoomService() *RoomService {
	return &RoomService{
		repo: repositories.NewRoomRepository(),
	}
}

func (r *RoomService) InitializeRooms(rooms []dto.Room) {
	var rooms_model []models.Room
	for _, room := range rooms {
		rooms_model = append(rooms_model, models.Room{
			RoomNumber: room.RoomNumber,
			Status:     room.Status,
		})
	}
	r.repo.InitializeRooms(rooms_model)
}

func (r *RoomService) GetAvailableRooms() []models.Room {
	return r.repo.GetAvailableRooms()

}

func (r *RoomService) UpdateRoomStatus(room_number int, status enum.RoomStatus) {
	r.repo.UpdateRoomStatus(room_number, status)
}
