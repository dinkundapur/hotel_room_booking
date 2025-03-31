package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"hotel-room-booking/app/dto"
	"hotel-room-booking/app/handlers"
	handlers_v1 "hotel-room-booking/app/handlers/v1"
	"hotel-room-booking/app/utils"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateBooking(t *testing.T) {
	// Create a new Gin engine
	gin.SetMode(gin.TestMode)
	r := gin.Default()

	// Initialize rooms
	utils.InitializeRooms()

	// Create a new booking handler
	hotelHandler := handlers_v1.InitHotelHandler()

	// Register the CreateBooking route
	r.POST("/api/v1/room/book", hotelHandler.CreateBooking)

	// Create a new booking request
	bookingRequest := dto.CreateBookingRequest{
		GuestName:    "Test",
		Email:        "test@example.com",
		MobileNumber: 9844038112,
		CheckOut:     "2024-03-17 13:12:20",
		CheckIn:      "2024-03-14 13:12:20",
	}

	// Convert the booking request to JSON
	jsonBytes, err := json.Marshal(bookingRequest)

	assert.NoError(t, err)
	// Create a new HTTP request
	req, err := http.NewRequest("POST", "/api/v1/room/book", bytes.NewBuffer(jsonBytes))
	assert.NoError(t, err)

	// Create a new response recorder
	w := httptest.NewRecorder()

	// Call the CreateBooking route
	r.ServeHTTP(w, req)

	// Unmarshal the response JSON into a Booking struct
	var response handlers.APIResponse

	err = json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	booking := response.Data.(map[string]any)

	// Assert that the booking guest name matches the request guest name
	assert.Equal(t, bookingRequest.GuestName, booking["guest_name"])
}
