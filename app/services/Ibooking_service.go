package services

import "hotel-room-booking/app/dto"

type IBookingService interface {
	CreateBooking(booking dto.CreateBookingRequest) (dto.CreateBookingResponse, error)
	GetBookings() ([]dto.GetBookingResponse, error)
	GetBookingByEmail(email string) ([]dto.GetBookingResponse, error)
	UpdateBooking(room_number int, booking_data dto.UpdateBookingRequest) error
}
