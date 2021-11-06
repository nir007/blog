package services

import (
	"github.com/nir007/blog/internal/contracts"
	"time"
)

type AttemptConfirmService struct {
	db contracts.DatabaseFucker
}

func NewAttemptConfirmService(db contracts.DatabaseFucker) *AttemptConfirmService{
	return &AttemptConfirmService{
		db: db,
	}
}

const insert = `INSERT INTO db_schema.attempt_confirm (uid, code, phone, date)
	VALUES($1, $2, $3, $4)`

const oneAttempt = `SELECT id FROM db_schema.attempt_confirm 
	WHERE uid = $1 AND code = $2 AND phone = $3 ORDER BY id LIMIT 1`

const selectCountLastHour = `SELECT count(*) FROM db_schema.attempt_confirm 
	WHERE uid = $1 AND phone = $2 AND date + INTERVAL '1 hour' > $3`


func (s *AttemptConfirmService) Add(a *contracts.AttemptConfirm) (id int32, err error) {
	return s.db.Execute(
		insert,
		a.Uid,
		a.Code,
		a.Phone,
		a.Date,
	)
}

func (s *AttemptConfirmService) One(attempt *contracts.AttemptConfirm) (id int32, err error) {

	rows, err := s.db.ExecuteSelect(
		oneAttempt,
		attempt.Uid,
		attempt.Code,
		attempt.Phone,
	)

	for rows.Next() {
		rows.Scan(&id)
	}

	return id, err
}

func (s *AttemptConfirmService) CountLastHour(uid int32, phone string) (count int32, err error) {

	rows, err := s.db.ExecuteSelect(
		selectCountLastHour,
		uid,
		phone,
		time.Now(),
	)

	if err == nil {
		for rows.Next() {
			rows.Scan(&count)
		}
	}

	return count, err
}

func (s *AttemptConfirmService) CodeExists(attempt *contracts.AttemptConfirm) (confirmed bool, err error) {
	id, err := s.One(attempt)
	return id > 0, err
}
