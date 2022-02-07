package comment

import "sirclo/entities"

type Comment interface {
	GetComments(eventId int, limit int, offset int) ([]entities.CommentResponse, error)
	GetComment(commentId int) (entities.CommentResponse, error)
	GetTotalPageComments(eventId int) int
	CreateComment(comment entities.Comment) error
	DeleteComment(eventId int, commentId int, loginId int) error
	EditComment(comment entities.Comment) error
}
