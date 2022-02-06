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

func (er *EventRepository) GetEventParam(param string) ([]entities.EventCat, error) {
	var events []entities.EventCat
	convParam := "%" + param + "%"
	result, err := er.db.Query(`SELECT e.id, e.user_id, e.category_id, e.title, e.host, e.date, e.location, e.description, e.image_url, c.id as category_id, c.category FROM events e JOIN categories c ON e.category_id = c.id WHERE e.title LIKE ? OR e.location LIKE ? OR c.category LIKE ? AND e.deleted_at IS NULL`, convParam, convParam, convParam)
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

func (er *EventRepository) GetMyEvents(loginId int) ([]entities.Event, error) {
	var events []entities.Event
	result, err := er.db.Query(`select id, user_id, category_id, title, host, date, location, description, image_url from events WHERE user_id = ? AND deleted_at IS NULL`, loginId)
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

func (er *EventRepository) GetEventByCatID(categoryID int) ([]entities.Event, error) {
	var events []entities.Event
	result, err := er.db.Query(`select id, user_id, category_id, title, host, date, location, description, image_url from events WHERE category_id = ? AND deleted_at IS NULL`, categoryID)
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

func (er *EventRepository) GetEventJoinedByUser(loginId int) ([]entities.Event, error) {
	var events []entities.Event
	// result, err := er.db.Query(`select e.id, e.user_id, e.category_id, e.title, e.host, e.date, e.location, e.description, e.image_url, p.user_id from events e JOIN participants p ON e.user_id = p.user_id WHERE p.user_id = ? AND p.deleted_at IS NULL`, loginId)
	// if err != nil {
	// 	return nil, err
	// }

	// for result.Next() {
	// 	var event entities.Event

	// 	err = result.Scan(&event.Id, &event.UserID, &event.CategoryId, &event.Title, &event.Host, &event.Date, &event.Location, &event.Description, &event.ImageUrl)

	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	events = append(events, event)
	// }
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

func (er *EventRepository) UpdateEvent(event entities.Event) (entities.Event, error) {
	query := `UPDATE events SET category_id = ?, title = ?, host = ?, date = ?, location = ?, description = ?, image_url = ?, updated_at = now() WHERE id = ? AND user_id = ?`

	statement, err := er.db.Prepare(query)
	if err != nil {
		return event, err
	}

	result, err_ex := statement.Exec(event.CategoryId, event.Title, event.Host, event.Date, event.Location, event.Description, event.ImageUrl, event.Id, event.UserID)
	if err_ex != nil {
		return event, err_ex
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return event, fmt.Errorf("event not found")
	}
	return event, nil
}

func (er *EventRepository) DeleteEvent(event entities.Event, loginId int) (entities.Event, error) {
	query := `Update events SET deleted_at = now() WHERE id = ? AND user_id = ?`

	statement, err := er.db.Prepare(query)
	if err != nil {
		return event, err
	}
	fmt.Println("ini ei & li: ", event.Id, loginId)
	result, err_exec := statement.Exec(event.Id, loginId)
	if err_exec != nil {
		return event, err_exec
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return event, fmt.Errorf("event not found")
	}

	return event, nil

}
