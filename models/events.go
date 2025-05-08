package models

import (
	"api/db"
	"log"
	"time"
)

type Event struct {
	ID          int64
	Name        string    `binding:"required"` // For gin to validate incoming request body
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	Time        time.Time `binding:"required"`
	UserID      int
}

func (e Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, time, userId)
	VALUES (?, ?, ?, ?, ?)
	`
	//Note: Preparing complexing query has better performance
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.Time, e.UserID)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	e.ID = id
	return err

}

func GetAllEvents() ([]Event, error) {
	query := `
	  SELECT * FROM events
	`
	//Notes: 'Use exec when query is changing something in DB and Query when just getting'
	rows, err := db.DB.Query(query)
	log.Println("Error in query", err)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event
	for rows.Next() {
		var event Event
		//Note: Scan gives you current row
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Time, &event.UserID)
		log.Println("Error in for loop", err)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := `
     SELECT * FROM events WHERE ID = ?
   `
	row := db.DB.QueryRow(query, id)

	var event Event

	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.Time, &event.UserID)
	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (event Event) UpdateEventById(id int64) error {
	query := `
	  UPDATE events
	  SET name = ?, description = ?, time = ?, location = ?
	  WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.Name, event.Description, event.Time, event.Location, id)

	return err
}

func DeleteEventById(id int64) error {
	query := `
	  DELETE FROM events
	  WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}
