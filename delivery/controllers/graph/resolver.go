package graph

import (
	"sirclo/repository/auth"
	"sirclo/repository/comment"
	"sirclo/repository/event"
	"sirclo/repository/participant"
	"sirclo/repository/user"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	authRepo        auth.Auth
	userRepo        user.User
	commentRepo     comment.Comment
	participantRepo participant.Participant
	eventRepo       event.Event
}

func NewResolver(ur user.User, ar auth.Auth, cr comment.Comment, pr participant.Participant, er event.Event) *Resolver {
	return &Resolver{
		commentRepo:     cr,
		participantRepo: pr,
		authRepo:        ar,
		userRepo:        ur,
		eventRepo:       er,
	}
}
