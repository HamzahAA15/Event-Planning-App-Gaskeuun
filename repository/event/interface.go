package event

import "sirclo/entities"

type Event interface {
	GetEvents(limit, offset int) ([]entities.Event, error)
	GetEvent(eventID int) (entities.EventIdResponse, error)
	GetEventParam(param string, limit, offset int) ([]entities.EventCat, error)
	GetMyEvents(loginId int, limit, offset int) ([]entities.Event, error)
	GetEventByCatID(categoryID int, param string, limit, offset int) ([]entities.Event, error)
	GetEventJoinedByUser(loginId int, limit, offset int) ([]entities.JoinedEvent, error)
	GetTotalEvents(param string) int
	GetTotalMyEvents(loginId int) int
	GetTotalJoinedEvents(loginId int) int
	GetTotalEventsByCatId(categoryID int, param string) int
	CreateEvent(Event entities.Event) (entities.Event, error)
	UpdateEvent(Event entities.Event, eventID, loginId int) error
	DeleteEvent(eventId int, loginId int) error
}
