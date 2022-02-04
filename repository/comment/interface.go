package comment

import "sirclo/entities"

type Comment interface {
	GetComments(eventId int) ([]entities.CommentResponse, error)
	GetComment(commentId int) (entities.CommentResponse, error)
	CreateComment(comment entities.Comment) error
	DeleteComment(eventId int, commentId int, loginId int) error
	EditComment(comment entities.Comment) error
}
