package services

import "hotel-room-booking/app/dto"

type IRoomService interface {
	CreateBooking(booking dto.Booking) error
	GetBookings() ([]dto.Booking, error)
}
