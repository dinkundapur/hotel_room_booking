package models

import "hotel-room-booking/app/enum"

type Booking struct {
	Id           string             `json:"id"`
	RoomNumber   int                `json:"room_number"`
	GuestName    string             `json:"guest_name"`
	MobileNumber int                `json:"mobile_number"`
	Email        string             `json:"email"`
	CheckIn      string             `json:"check_in"`
	CheckOut     string             `json:"check_out"`
	Status       enum.BookingStatus `json:"status"`
}
