package models

import "time"

type Booking struct {
	Id           string    `json:"id"`
	RoomNumber   int       `json:"room_number"`
	GuestName    string    `json:"guest_name"`
	MobileNumber int       `json:"mobile_number"`
	Email        string    `json:"email"`
	CheckIn      time.Time `json:"check_in"`
	CheckOut     time.Time `json:"check_out	"`
}
