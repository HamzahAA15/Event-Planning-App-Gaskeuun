package event

import (
	"database/sql"
	"fmt"
	"sirclo/entities"
)

type EventRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (er *EventRepository) GetEvents(limit, offset int) ([]entities.Event, error) {
	var events []entities.Event
	result, err := er.db.Query(`select id, user_id, category_id, title, host, date, location, description, image_url from events WHERE deleted_at IS NULL LIMIT ? OFFSET ?`, limit, offset)
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

func (er *EventRepository) GetEvent(eventID int) (entities.EventIdResponse, error) {
	var event entities.EventIdResponse
	result := er.db.QueryRow(`SELECT id, user_id, category_id, title, host, date, location, description, image_url FROM events WHERE id = ? AND deleted_at IS NULL`, eventID)

	err := result.Scan(&event.Id, &event.UserID, &event.CategoryId, &event.Title, &event.Host, &event.Date, &event.Location, &event.Description, &event.ImageUrl)

	if err != nil {
		return event, err
	}
	return event, nil
}

func (er *EventRepository) GetEventParam(param string, limit, offset int) ([]entities.EventCat, error) {
	var events []entities.EventCat
	convParam := "%" + param + "%"
	result, err := er.db.Query(`SELECT e.id, e.user_id, e.category_id, e.title, e.host, e.date, e.location, e.description, e.image_url, c.id as category_id, c.category FROM events e JOIN categories c ON e.category_id = c.id WHERE e.title LIKE ? OR e.location LIKE ? OR c.category LIKE ? AND e.deleted_at IS NULL LIMIT ? OFFSET ?`, convParam, convParam, convParam, limit, offset)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var event entities.EventCat
		err = result.Scan(&event.Id, &event.UserID, &event.CategoryId, &event.Title, &event.Host, &event.Date, &event.Location, &event.Description, &event.ImageUrl, &event.Categories.Id, &event.Categories.Category)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func (er *EventRepository) GetMyEvents(loginId int, limit, offset int) ([]entities.Event, error) {
	var events []entities.Event
	result, err := er.db.Query(`select id, user_id, category_id, title, host, date, location, description, image_url from events WHERE user_id = ? AND deleted_at IS NULL LIMIT ? OFFSET ?`, loginId, limit, offset)
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

func (er *EventRepository) GetEventByCatID(categoryID int, param string, limit, offset int) ([]entities.Event, error) {
	convParam := "%" + param + "%"
	var events []entities.Event
	result, err := er.db.Query(`select id, user_id, category_id, title, host, date, location, description, image_url from events WHERE category_id = ? AND deleted_at IS NULL AND title LIKE ? OR location LIKE ? LIMIT ? OFFSET ?`, categoryID, convParam, convParam, limit, offset)
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

func (er *EventRepository) GetEventJoinedByUser(loginId int, limit, offset int) ([]entities.JoinedEvent, error) {
	var events []entities.JoinedEvent
	result, err := er.db.Query(`select e.id, e.category_id, e.title, e.host, e.date, e.location, e.description, e.image_url, p.event_id, p.user_id as participant from events e JOIN participants p ON e.id = p.event_id WHERE p.user_id = ? AND p.deleted_at IS NULL LIMIT ? OFFSET ?`, loginId, limit, offset)
	if err != nil {
		return nil, err
	}

	for result.Next() {
		var event entities.JoinedEvent

		err = result.Scan(&event.Id, &event.CategoryId, &event.Title, &event.Host, &event.Date, &event.Location, &event.Description, &event.ImageUrl, &event.EventId, &event.UserId)

		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
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

func (er *EventRepository) UpdateEvent(event entities.Event, eventID, loginId int) error {
	query := `UPDATE events SET category_id = ?, title = ?, host = ?, date = ?, location = ?, description = ?, image_url = ?, updated_at = now() WHERE id = ? AND user_id = ? AND deleted_at IS NULL`

	statement, err := er.db.Prepare(query)
	if err != nil {
		return err
	}

	result, err_ex := statement.Exec(event.CategoryId, event.Title, event.Host, event.Date, event.Location, event.Description, event.ImageUrl, eventID, loginId)
	if err_ex != nil {
		return err_ex
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("event not found")
	}
	return nil
}

func (er *EventRepository) DeleteEvent(eventId int, loginId int) error {
	query := `Update events SET deleted_at = now() WHERE id = ? AND user_id = ? AND deleted_at IS NULL`

	statement, err := er.db.Prepare(query)
	if err != nil {
		return err
	}
	fmt.Println("ini ei & li: ", eventId, loginId)
	result, err_exec := statement.Exec(eventId, loginId)
	if err_exec != nil {
		return err_exec
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("event not found")
	}

	return nil

}

func (er *EventRepository) GetTotalEvents(param string) int {
	var hasil int
	convParam := "%" + param + "%"
	result_check, _ := er.db.Query("SELECT COUNT(id) FROM events WHERE title LIKE ? OR location LIKE ? AND deleted_at IS null", convParam, convParam)
	for result_check.Next() {
		err := result_check.Scan(&hasil) // sql null string kalau mau skip
		if err != nil {
			return 1
		}
		return hasil
	}
	return 1
}

func (er *EventRepository) GetTotalMyEvents(loginId int) int {
	var hasil int
	result_check, _ := er.db.Query("SELECT COUNT(id) FROM events WHERE user_id = ?  AND deleted_at IS null", loginId)
	for result_check.Next() {
		err := result_check.Scan(&hasil) // sql null string kalau mau skip
		if err != nil {
			return 1
		}
		return hasil
	}
	return 1
}

func (er *EventRepository) GetTotalJoinedEvents(loginId int) int {
	var hasil int
	result_check, _ := er.db.Query("SELECT COUNT(e.id) FROM events e JOIN participants p ON e.id = p.event_id WHERE p.user_id = ? AND p.deleted_at IS null", loginId)
	for result_check.Next() {
		err := result_check.Scan(&hasil) // sql null string kalau mau skip
		if err != nil {
			return 1
		}
		return hasil
	}
	return 1
}

func (er *EventRepository) GetTotalEventsByCatId(categoryID int, param string) int {
	var hasil int
	convParam := "%" + param + "%"
	result_check, _ := er.db.Query("SELECT COUNT(id) FROM events WHERE category_id = ? AND title LIKE ? OR location LIKE ? AND deleted_at IS null", categoryID, convParam, convParam)
	for result_check.Next() {
		err := result_check.Scan(&hasil) // sql null string kalau mau skip
		if err != nil {
			return 1
		}
		return hasil
	}
	return 1
}
