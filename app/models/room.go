package models

import "hotel-room-booking/app/enum"

type Room struct {
	RoomNumber int
	Status     enum.RoomStatus
}
