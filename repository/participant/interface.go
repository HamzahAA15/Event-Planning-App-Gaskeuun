package participant

import "sirclo/entities"

type Participant interface {
	GetParticipants(eventId int, limit int, offset int) ([]entities.User, error)
	GetParticipantStatus(eventId int, loginId int) bool
	GetTotalPageParticipants(eventId int) int
	CreateParticipant(eventId int, loginId int) error
	DeleteParticipant(eventId int, loginId int) error
}
