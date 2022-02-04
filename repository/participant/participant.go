package participant

import (
	"database/sql"
	"fmt"
	"sirclo/entities"
)

type ParticipantRepository struct {
	db *sql.DB
}

func New(db *sql.DB) *ParticipantRepository {
	return &ParticipantRepository{db: db}
}

func (pr *ParticipantRepository) GetParticipants(eventId int) ([]entities.User, error) {
	var users []entities.User
	result, err := pr.db.Query(`
	select users.id, users.name, users.email, users.image_url from users
	JOIN participants ON users.Id = participants.user_id
	where participants.deleted_at is null AND users.deleted_at is null AND participants.event_id = ?`, eventId)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	for result.Next() {
		var user entities.User
		err := result.Scan(&user.Id, &user.Name, &user.Email, &user.ImageUrl) // sql null string kalau mau skip
		if err != nil {
			return nil, fmt.Errorf("user not found")
		}
		users = append(users, user)
	}
	return users, nil
}

func (pr *ParticipantRepository) CreateParticipant(eventId int, loginId int) error {
	result_check, _ := pr.db.Query("select id from participants where event_id = ? AND user_id = ? AND deleted_at IS null", eventId, loginId)
	for result_check.Next() {
		return fmt.Errorf("anda sudah terdaftar")
	}
	result, err := pr.db.Exec("INSERT INTO participants(event_id, user_id, updated_at) VALUES(?,?, now())", eventId, loginId)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("participant not created")
	}
	return nil
}

func (pr *ParticipantRepository) DeleteParticipant(eventId int, loginId int) error {
	result, err := pr.db.Exec("UPDATE participants SET deleted_at = now() where event_id = ? AND user_id = ? AND deleted_at IS null", eventId, loginId)
	if err != nil {
		return err
	}
	mengubah, _ := result.RowsAffected()
	if mengubah == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
