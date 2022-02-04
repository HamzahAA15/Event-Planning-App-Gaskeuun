package event

import "sirclo/entities"

type Event interface {
	GetEvents() ([]entities.Event, error)
	GetEvent(eventID int) (entities.Event, error)
	GetEventParam(param string) ([]entities.EventCat, error)
	CreateEvent(Event entities.Event) (entities.Event, error)
	UpdateEvent(Event entities.Event) (entities.Event, error)
	DeleteEvent(Event entities.Event, loginId int) (entities.Event, error)
}
