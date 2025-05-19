package models

import (
	"time"
    "example.com/testserver/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	DateTime    time.Time `json:"date_time" binding:"required"`
	UserID      int64       `json:"user_id"`
	
}

var events = []Event{}

func (e *Event) Save() string {
	query := `
	INSERT INTO events (name, description, location, dateTime, userID)
	 VALUES (?, ?, ?, ?, ?)
	 `
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return "Error preparing statement"
	}
	defer stmt.Close()
	// Execute the statement with the event data
	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		return "Error executing statement"	
	}
	id, err := result.LastInsertId()
	if err != nil {
		return "Error getting last insert ID"
	}
	e.ID = id	
	return "Event saved successfully"
}

func (e *Event) Update() string {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, dateTime = ?, userID = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return "Error preparing statement"
	}
	defer stmt.Close()
	// Execute the statement with the event data
	_, err = stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID, e.ID)
	if err != nil {
		return "Error executing statement"
	}
	return "Event updated successfully"
}

func (e *Event) Delete() error {
	
	query := `
	DELETE FROM events
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)
	
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)
	return err

}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}	
	defer rows.Close()
	var events []Event

	// Iterate through the rows and scan the data into the events slice
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		
		if err != nil {	
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventByID(id int64) (*Event, error) {
//func GetEventByID(id int64) (Event, error) {

	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	// The error from QueryRow (e.g., sql.ErrNoRows) is checked here, after Scan.
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		//return Event{}, err
		return nil, err
	}
	return &event, nil
}
