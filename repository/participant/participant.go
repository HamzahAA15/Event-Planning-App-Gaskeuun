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

func (pr *ParticipantRepository) GetParticipants(eventId int, limit int, offset int) ([]entities.User, error) {
	var users []entities.User
	result, err := pr.db.Query(`
	select users.id, users.name, users.email, users.image_url from users
	JOIN participants ON users.Id = participants.user_id
	where participants.deleted_at is null AND users.deleted_at is null AND participants.event_id = ? LIMIT ? OFFSET ?`, eventId, limit, offset)
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
	defer result_check.Close()
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

func (pr *ParticipantRepository) GetTotalPageParticipants(eventId int) int {
	var hasil int
	result_check, _ := pr.db.Query("select COUNT(id) from participants where event_id = ? AND deleted_at IS null", eventId)
	defer result_check.Close()
	for result_check.Next() {
		err := result_check.Scan(&hasil) // sql null string kalau mau skip
		if err != nil {
			return 1
		}
		return hasil
	}
	return 1
}

func (pr *ParticipantRepository) GetParticipantStatus(eventId int, loginId int) bool {
	var id int
	result := pr.db.QueryRow(`SELECT id FROM participants WHERE event_id = ? AND user_id = ? AND deleted_at IS NULL`, eventId, loginId)
	err := result.Scan(&id)
	if err != nil {
		return false
	}
	return true
}
