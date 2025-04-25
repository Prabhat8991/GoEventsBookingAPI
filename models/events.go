package models

import "time"

type Event struct {
	ID          int
	Name        string    `binding:"required"` // For gin to validate incoming request body 
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Time        time.Time `binding:"required"`
	UserID      int
}

var events = []Event{}

func (e Event) Save() {
	//TODO: Save event to db
	events = append(events, e)
}

func GetAllEvents() []Event {
	return events
}
