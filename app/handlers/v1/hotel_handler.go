package handlers_v1

import (
	"hotel-room-booking/app/dto"
	"hotel-room-booking/app/services"

	"github.com/gin-gonic/gin"
)

type HotelHandler struct {
	service services.IRoomService
}

func InitHotelHandler() *HotelHandler {
	return &HotelHandler{
		service: services.NewRoomService(),
	}
}

func (h *HotelHandler) CreateBooking(ctx *gin.Context) {
	var booking_data dto.Booking
	err := ctx.BindJSON(&booking_data)

	if err != nil { //still err will not be empty for right request also
		ctx.JSON(502, err)
	}

	h.service.CreateBooking(booking_data)
	ctx.JSON(200, "Booking created successfully")

}

func (h *HotelHandler) GetBookings(ctx *gin.Context) { //can pass email in query param

	data, err := h.service.GetBookings()
	if err != nil {
		ctx.JSON(502, err)
	}
	ctx.JSON(200, data)

}
