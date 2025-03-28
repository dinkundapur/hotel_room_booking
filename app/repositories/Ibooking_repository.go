package repositories

import "hotel-room-booking/app/models"

type IBookingRepository interface {
	CreateBooking(room models.Booking)
	GetBookings() []models.Booking
}
