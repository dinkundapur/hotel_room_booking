# hotel-room-booking

A hotel room booking system build using go


## Walkthrough

The following is an in-depth walkthrough of this project. This go API repo is responsible for 
- Initializing 6 rooms with room numbers 1111, 1112, 1113, 2222, 2223, 2224 with status Available
- 6 bookings can be made, room numbers will be assigned automatically
- If all 6 rooms are availble then we get message "No rooms available" for booking API
- Updated booking API can be called to update check_in, checkout_date or even update status to Cancelled or Checkedout for active bookings.
- Rooms will be available for booking if the current booking is cancelled or checkedout.


--------
## Requirements

- [Go](https://go.dev/)

--------

## Getting Started

The following are basic instructions to get started. 

### Step 1 - Get the code

Clone the code:

```shell
$ git clone https://github.com/dinkundapur/hotel_room_booking.git
Cloning into 'hotel_room_booking'...
```


### Step 2 - Create .env file
create .env file in project root, and copy .env_example file containt to .env


### Step 3 - Run the Web Server
```shell
$ go run main.go
```
#### Build binary
```shell
$ go build
```

## API contract
### Book a room
POST localhost:8089/api/v1/room/book

Request Body:
```
{

    "guest_name": "Dinesh",
    "mobile_number":9844038112,
    "email": "dinkundapur@gmail.com",
    "check_in": "2024-01-09 10:12:12",
    "check_out": "2024-01-10 11:12:12"
}
```

Response Body:
```
{
    "success": true,
    "data": {
        "room_number": 1111,
        "guest_name": "Dinesh12",
        "check_in": "2024-01-09 10:12:12",
        "check_out": "2024-01-10 11:12:12"
    },
    "message": ""
}
```

Response Body when all 6 rooms booked:

```
{
    "success": true,
    "data": null,
    "message": "No rooms available"
}
```


### API to get all Bookings
GET localhost:8089/api/v1/room/bookings

```
{
    "success": true,
    "data": [
        {
            "id": "ROOMA9OSIOM2L97T7OS9",
            "room_number": 1111,
            "guest_name": "Dinesh",
            "mobile_number":9844038112,
            "email": "dinkundapur@gmail.com",
            "check_in": "2024-01-09 10:12:12",
            "check_out": "2024-01-10 11:12:12",
            "status": "Active"
        },
        {
            "id": "ROOM03XM2GIRYTEGZEC4",
            "room_number": 1113,
            "guest_name": "Mahesh",
            "mobile_number": 9844038115,
            "email": "test@gmail.com",
            "check_in": "2024-01-09 10:12:12",
            "check_out": "2024-01-10 11:12:12",
            "status": "Active"
        }
    ],
    "message": ""
}
```


### API to get specific user bookings
GET localhost:8089/api/v1/room/booking/dinkundapur@gmail.com
```
{
    "success": true,
    "data": [
        {
            "id": "ROOMA9OSIOM2L97T7OS9",
            "room_number": 1111,
            "guest_name": "Dinesh",
            "mobile_number":9844038112,
            "email": "dinkundapur@gmail.com",
            "check_in": "2024-01-09 10:12:12",
            "check_out": "2024-01-10 11:12:12",
            "status": "Active"
        }
    ],
    "message": ""
}
```

### API to update booking status as Cancelled or CheckedOut or updating check_in and Check_out
POST localhost:8089/api/v1/room/update/booking/1111

```
{   
    "email": "dinkundapur@gmail.com",
    "status":"CANCELLED", //optional
    "check_in": "2024-01-23 10:12:12",//optional
    "check_out": "2024-01-24 11:12:12" //optional
}
```
```
{
    "success": true,
    "data": "Booking info updated successfully",
    "message": ""
}
```
