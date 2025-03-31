package repositories

import (
	"hotel-room-booking/app/enum"
	"hotel-room-booking/app/models"
	"sync"
)

type RoomRepository struct {
	Rooms map[int]models.Room
	mu    sync.RWMutex
}

var (
	roomRepoInstance *RoomRepository
	once             sync.Once
)

func NewRoomRepository() *RoomRepository {
	once.Do(func() { //Implemeting Singleton instance
		roomRepoInstance = &RoomRepository{
			Rooms: make(map[int]models.Room),
		}
	})
	return roomRepoInstance
}

func (r *RoomRepository) InitializeRooms(rooms []models.Room) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, room := range rooms {
		r.Rooms[room.RoomNumber] = room
	}
}

func (r *RoomRepository) GetAvailableRooms() []models.Room {
	r.mu.RLock()
	defer r.mu.RUnlock()
	var availableRooms []models.Room
	for _, room := range r.Rooms {
		if room.Status == "Available" {
			availableRooms = append(availableRooms, room)
		}

	}
	return availableRooms
}

func (r *RoomRepository) UpdateRoomStatus(room_number int, status enum.RoomStatus) {
	r.mu.Lock()
	defer r.mu.Unlock()
	room, ok := r.Rooms[room_number]
	if ok {
		room.Status = status
		r.Rooms[room_number] = room
	}
}
