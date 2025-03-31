package dto

import "hotel-room-booking/app/enum"

type CreateBookingRequest struct {
	GuestName    string `json:"guest_name"`
	MobileNumber int    `json:"mobile_number"`
	Email        string `json:"email"`
	CheckIn      string `json:"check_in"`
	CheckOut     string `json:"check_out"`
}

type CreateBookingResponse struct {
	RoomNumber int    `json:"room_number"`
	GuestName  string `json:"guest_name"`
	CheckIn    string `json:"check_in"`
	CheckOut   string `json:"check_out"`
}

type GetBookingResponse struct {
	Id           string             `json:"id"`
	RoomNumber   int                `json:"room_number"`
	GuestName    string             `json:"guest_name"`
	MobileNumber int                `json:"mobile_number"`
	Email        string             `json:"email"`
	CheckIn      string             `json:"check_in"`
	CheckOut     string             `json:"check_out"`
	Status       enum.BookingStatus `json:"status"`
}

type UpdateBookingRequest struct {
	Email    string             `json:"email"`
	Status   enum.BookingStatus `json:"status,omitempty"`
	CheckIn  string             `json:"check_in,omitempty"`
	CheckOut string             `json:"check_out,omitempty"`
}
