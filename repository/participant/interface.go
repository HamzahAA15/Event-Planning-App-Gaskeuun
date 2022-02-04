package participant

import "sirclo/entities"

type Participant interface {
	GetParticipants(eventId int) ([]entities.User, error)
	CreateParticipant(eventId int, loginId int) error
	DeleteParticipant(eventId int, loginId int) error
}
