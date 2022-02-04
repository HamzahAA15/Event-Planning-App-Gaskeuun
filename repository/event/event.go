package event

import (
	"database/sql"
	"sirclo/entities"
)

type EventRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (er *EventRepository) GetEvents() ([]entities.Event, error) {
	var events []entities.Event
	result, err := er.db.Query(`select id, user_id, category_id, title, host, date, location, description, image_url from events WHERE deleted_at IS NULL`)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var event entities.Event

		err = result.Scan(&event.Id, &event.UserID, &event.CategoryId, &event.Title, &event.Host, &event.Date, &event.Location, &event.Description, &event.ImageUrl)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (er *EventRepository) GetEvent(eventID int) (entities.Event, error) {
	var event entities.Event
	result := er.db.QueryRow(`SELECT id, user_id, category_id, title, host, date, location, description, image_url FROM events WHERE id = ? AND deleted_at IS NULL`, eventID)

	err := result.Scan(&event.Id, &event.UserID, &event.CategoryId, &event.Title, &event.Host, &event.Date, &event.Location, &event.Description, &event.ImageUrl)

	if err != nil {
		return event, err
	}
	return event, nil
}

func (er *EventRepository) CreateEvent(event entities.Event) (entities.Event, error) {
	query := `INSERT INTO events (user_id, category_id, title, host, date, location, description, image_url) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	statement, err := er.db.Prepare(query)
	if err != nil {
		return event, err
	}

	_, err = statement.Exec(event.UserID, event.CategoryId, event.Title, event.Host, event.Date, event.Location, event.Description, event.ImageUrl)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (er *EventRepository) UpdateEvent(event entities.Event) (entities.Event, error) {
	query := `UPDATE events SET category_id = ?, title = ?, host = ?, date = ?, location = ?, description = ?, image_url = ?, updated_at = now() WHERE id = ?`

	statement, err := er.db.Prepare(query)
	if err != nil {
		return event, err
	}

	_, err = statement.Exec(event.CategoryId, event.Title, event.Host, event.Date, event.Location, event.Description, event.ImageUrl, event.Id)
	if err != nil {
		return event, err
	}
	return event, nil
}

func (er *EventRepository) DeleteEvent(event entities.Event) (entities.Event, error) {
	query := `Update events SET deleted_at = now() WHERE id = ?`

	statement, err := er.db.Prepare(query)
	if err != nil {
		return event, err
	}

	_, err = statement.Exec(event.Id)
	if err != nil {
		return event, err
	}

	return event, nil

}
