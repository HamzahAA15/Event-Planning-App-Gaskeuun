package comment

import (
	"database/sql"
	"fmt"
	"sirclo/entities"
)

type CommentRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *CommentRepository {
	return &CommentRepository{db: db}
}

func (cr *CommentRepository) GetComments(eventId int) ([]entities.CommentResponse, error) {
	var comments []entities.CommentResponse
	result, err := cr.db.Query(`
	select comments.id, users.id, users.name, users.email, users.image_url, comments.comment from users
	JOIN comments ON users.Id = comments.user_id
	where comments.deleted_at is null AND users.deleted_at is null AND comments.event_id = ?`, eventId)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var comment entities.CommentResponse
		err := result.Scan(&comment.Id, &comment.User.Id, &comment.User.Name, &comment.User.Email, &comment.User.ImageUrl, &comment.Comment)
		if err != nil {
			return nil, fmt.Errorf("comment not found")
		}
		comments = append(comments, comment)
	}
	return comments, nil
}

func (cr *CommentRepository) GetComment(commentId int) (entities.CommentResponse, error) {
	var comment entities.CommentResponse
	result, err := cr.db.Query(`
	select comments.id, users.id, users.name, users.email, users.image_url, comments.comment from users
	JOIN comments ON users.Id = comments.user_id
	where comments.deleted_at is null AND users.deleted_at is null AND comments.id = ?`, commentId)
	if err != nil {
		return comment, err
	}
	defer result.Close()
	for result.Next() {
		err := result.Scan(&comment.Id, &comment.User.Id, &comment.User.Name, &comment.User.Email, &comment.User.ImageUrl, &comment.Comment)
		if err != nil {
			return comment, fmt.Errorf("comment not found")
		}
		return comment, nil
	}
	return comment, fmt.Errorf("comment not found")
}

func (cr *CommentRepository) CreateComment(comment entities.Comment) error {
	result, err := cr.db.Exec("INSERT INTO comments(user_id, event_id, comment, updated_at) VALUES(?,?,?,now())", comment.UserId, comment.EventId, comment.Comment)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("comment not created")
	}
	return nil
}

func (cr *CommentRepository) EditComment(comment entities.Comment) error {
	result, err := cr.db.Exec("UPDATE comments SET comment = ?, updated_at = now() WHERE id = ? AND event_id = ? AND user_id = ? AND deleted_at IS null", comment.Comment, comment.Id, comment.EventId, comment.UserId)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("comment not found")
	}
	return nil
}

func (cr *CommentRepository) DeleteComment(eventId int, commentId int, loginId int) error {
	result, err := cr.db.Exec("UPDATE comments SET deleted_at = now() where event_id = ? AND id = ? AND user_id = ? AND deleted_at IS null", eventId, commentId, loginId)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("comment not found")
	}
	return nil
}
