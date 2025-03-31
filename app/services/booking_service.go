package services

import (
	"encoding/json"
	"errors"
	"hotel-room-booking/app/dto"
	"hotel-room-booking/app/enum"
	"hotel-room-booking/app/models"
	"hotel-room-booking/app/repositories"
	room_service "hotel-room-booking/app/services/room"
	"hotel-room-booking/app/utils"
	"log"
)

type BookingService struct {
	repo      repositories.IBookingRepository
	room_serv room_service.IRoomService
}

func NewBookingService() *BookingService {
	return &BookingService{
		repo:      repositories.NewBookingRepository(),
		room_serv: room_service.NewRoomService(),
	}
}

func (r *BookingService) CreateBooking(booking dto.CreateBookingRequest) (dto.CreateBookingResponse, error) {
	var data dto.CreateBookingResponse

	//Check Rooms available or not
	available_rooms := r.room_serv.GetAvailableRooms()
	if len(available_rooms) == 0 {
		return data, errors.New("No rooms available")
	}

	var booking_model models.Booking
	bk, err := json.Marshal(booking)
	if err != nil {
		log.Println("Error ", err)
		return data, err
	}

	json.Unmarshal(bk, &booking_model)
	booking_model.Id = "ROOM" + utils.GenerateRandomAlphaNum()
	booking_model.RoomNumber = available_rooms[0].RoomNumber
	booking_model.Status = enum.ACTIVE
	r.repo.CreateBooking(booking_model)
	r.room_serv.UpdateRoomStatus(available_rooms[0].RoomNumber, enum.BOOKED)

	data.RoomNumber = booking_model.RoomNumber
	data.GuestName = booking_model.GuestName
	data.CheckIn = booking_model.CheckIn
	data.CheckOut = booking_model.CheckOut

	return data, nil
}

func (r *BookingService) GetBookings() ([]dto.GetBookingResponse, error) {
	var booking_model []dto.GetBookingResponse
	data := r.repo.GetBookings()

	bk, err := json.Marshal(data)

	if err != nil {
		log.Println("Error ", err)
		return booking_model, err
	}
	json.Unmarshal(bk, &booking_model)
	return booking_model, nil

}

func (r *BookingService) GetBookingByEmail(email string) ([]dto.GetBookingResponse, error) {
	var bookings_array []dto.GetBookingResponse
	data := r.repo.GetBookings()

	for _, booking_data := range data {
		if booking_data.Email == email {
			bk, err := json.Marshal(booking_data)

			if err != nil {
				log.Println("Error ", err)
				return bookings_array, err
			}
			var booking dto.GetBookingResponse
			json.Unmarshal(bk, &booking)

			bookings_array = append(bookings_array, booking)
		}
	}

	return bookings_array, nil
}

// Update Booking only if status is Booked
func (r *BookingService) UpdateBooking(room_number int, booking_data dto.UpdateBookingRequest) error {

	bookings := r.repo.GetAllBookingsByRoomNumber(room_number, booking_data.Email)

	if (len(bookings)) > 0 {
		for _, booking := range bookings {
			if booking.Status == enum.ACTIVE { //Booking is active
				if booking_data.Status != "" {
					booking.Status = booking_data.Status

				}

				if booking_data.CheckIn != "" {
					booking.CheckIn = booking_data.CheckIn

				}
				if booking_data.CheckOut != "" {
					booking.CheckOut = booking_data.CheckOut
				}

				r.repo.UpdateBooking(booking)
				return nil
			}
		}
	}

	return errors.New("Booking not found")
}
