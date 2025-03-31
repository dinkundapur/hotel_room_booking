package route_v1

import (
	handlers_v1 "hotel-room-booking/app/handlers/v1"
	utils "hotel-room-booking/app/utils"

	"github.com/gin-gonic/gin"
)

type route struct {
	RouteGroup *gin.RouterGroup
	ctx        *gin.Context
}

func Route(c *gin.Context, rGroup *gin.RouterGroup) {
	r := route{}
	r.RouteGroup = rGroup.Group("v1")

	// Calling all the route methods
	utils.ExecuteMethods(r)
}

func (r route) RoomRoutes() {
	hotelHandler := handlers_v1.InitHotelHandler()

	roomRoutes := r.RouteGroup.Group("room") // v1/room
	roomRoutes.POST("/book", hotelHandler.CreateBooking)
	roomRoutes.GET("/bookings", hotelHandler.GetBookings)
	roomRoutes.GET("/booking/:email", hotelHandler.GetBookingByEmail)
	roomRoutes.POST("/update/booking/:room-number", hotelHandler.UpdateBooking)

}
