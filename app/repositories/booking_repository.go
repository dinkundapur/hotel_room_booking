package repositories

import (
	"fmt"
	"hotel-room-booking/app/models"
	"sync"
)

type BookingRepository struct {
	Bookings map[string]models.Booking
	mu       sync.RWMutex
}

func NewBookingRepository() *BookingRepository {
	return &BookingRepository{
		Bookings: make(map[string]models.Booking),
	}

}

func (r *BookingRepository) CreateBooking(room models.Booking) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.Bookings[room.Id] = room

	fmt.Println(r.Bookings)
}

func (r *BookingRepository) GetBookings() []models.Booking {
	r.mu.RLock()
	defer r.mu.RUnlock()
	bookings := make([]models.Booking, 0, len(r.Bookings))
	for _, booking := range r.Bookings {
		bookings = append(bookings, booking)
	}

	fmt.Println(bookings)
	return bookings
}
