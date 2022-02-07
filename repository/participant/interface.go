package participant

import "sirclo/entities"

type Participant interface {
	GetParticipants(eventId int, limit int, offset int) ([]entities.User, error)
	GetTotalPageParticipants(eventId int) int
	CreateParticipant(eventId int, loginId int) error
	DeleteParticipant(eventId int, loginId int) error
}
