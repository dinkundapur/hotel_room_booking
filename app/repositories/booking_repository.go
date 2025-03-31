package repositories

import (
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

func (b *BookingRepository) CreateBooking(room models.Booking) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Bookings[room.Id] = room
}

func (b *BookingRepository) GetBookings() []models.Booking {
	b.mu.RLock()
	defer b.mu.RUnlock()
	bookings := make([]models.Booking, 0, len(b.Bookings))
	for _, booking := range b.Bookings {
		bookings = append(bookings, booking)
	}

	return bookings
}

// It will give cancelled, active, completed bookings for given room number and email
func (b *BookingRepository) GetAllBookingsByRoomNumber(room_number int, email string) []models.Booking {
	b.mu.RLock()
	defer b.mu.RUnlock()
	bookings := make([]models.Booking, 0)
	for _, booking := range b.Bookings {
		if booking.RoomNumber == room_number && booking.Email == email {
			bookings = append(bookings, booking)
		}
	}

	return bookings
}

func (b *BookingRepository) UpdateBooking(booking models.Booking) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Bookings[booking.Id] = booking

}
