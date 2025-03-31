package handlers_v1

import (
	"hotel-room-booking/app/dto"
	"hotel-room-booking/app/handlers"
	"hotel-room-booking/app/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

type HotelHandler struct {
	handlers.BaseHandler
	service services.IBookingService
}

func InitHotelHandler() *HotelHandler {
	return &HotelHandler{
		service: services.NewBookingService(),
	}
}

func (h *HotelHandler) CreateBooking(ctx *gin.Context) {
	var booking_data dto.CreateBookingRequest
	err := ctx.BindJSON(&booking_data)

	if err != nil { //still err will not be empty for right request also
		h.ErrorResponse(ctx, err)
		return
	}

	data, err := h.service.CreateBooking(booking_data)
	if err != nil {
		h.ErrorResponse(ctx, err)
		return
	}
	h.SuccessResponse(ctx, data)
	return
}

func (h *HotelHandler) GetBookings(ctx *gin.Context) {

	data, err := h.service.GetBookings()
	if err != nil {
		h.ErrorResponse(ctx, err)
		return
	}
	h.SuccessResponse(ctx, data)
}

func (h *HotelHandler) GetBookingByEmail(ctx *gin.Context) {

	email := ctx.Param("email")
	data, err := h.service.GetBookingByEmail(email)
	if err != nil {
		h.ErrorResponse(ctx, err)
		return
	}
	h.SuccessResponse(ctx, data)
}

func (h *HotelHandler) UpdateBooking(ctx *gin.Context) {
	var booking_data dto.UpdateBookingRequest
	err := ctx.BindJSON(&booking_data)

	if err != nil {
		h.ErrorResponse(ctx, err)
		return
	}

	room_number := ctx.Param("room-number")
	room_number_int, err := strconv.Atoi(room_number)

	if err != nil {
		h.ErrorResponse(ctx, err)
		return
	}

	err = h.service.UpdateBooking(room_number_int, booking_data)
	if err != nil {
		h.ErrorResponse(ctx, err)
		return
	}

	h.SuccessResponse(ctx, "Booking info updated successfully")
	return
}
