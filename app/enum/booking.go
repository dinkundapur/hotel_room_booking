package enum

type BookingStatus string

const (
	ACTIVE     BookingStatus = "Active"
	CANCELLED  BookingStatus = "Cancelled"
	CHECKEDOUT BookingStatus = "CheckedOut"
)

type RoomStatus string

const (
	AVAILABLE RoomStatus = "Available"
	BOOKED    RoomStatus = "Booked"
)
