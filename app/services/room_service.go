package services

import (
	"encoding/json"
	"hotel-room-booking/app/dto"
	"hotel-room-booking/app/models"
	"hotel-room-booking/app/repositories"
	"log"
)

type RoomService struct {
	repo repositories.IBookingRepository
}

func NewRoomService() *RoomService {
	return &RoomService{
		repo: repositories.NewBookingRepository(),
	}
}

func (r *RoomService) CreateBooking(booking dto.Booking) error {
	var booking_model models.Booking
	bk, err := json.Marshal(booking)

	if err != nil {
		log.Println("Error ", err)
		return err
	}

	json.Unmarshal(bk, &booking_model)
	r.repo.CreateBooking(booking_model)
	return nil
}

func (r *RoomService) GetBookings() ([]dto.Booking, error) {
	var booking_model []dto.Booking
	data := r.repo.GetBookings()

	bk, err := json.Marshal(data)

	if err != nil {
		log.Println("Error ", err)
		return booking_model, err
	}
	json.Unmarshal(bk, &booking_model)
	return booking_model, nil

}
