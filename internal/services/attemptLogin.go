package services

import (
	"errors"
	"time"

	"github.com/nir007/blog/internal/contracts"
)

const maxAttempts = 3

type AttemptLoginService struct {
	smsClient contracts.Sms
	db contracts.DatabaseFucker
}

func NewAttemptLoginService(smsClient contracts.Sms, db contracts.DatabaseFucker) *AttemptLoginService {
	return &AttemptLoginService{
		db: db,
		smsClient: smsClient,
	}
}

const selectAttemptsGetCode = `SELECT count(*) FROM db_schema.attempt_get_code 
	WHERE position($1 in phone) > 0 AND ip = $2 AND date + INTERVAL '1 hour' > $3`

const insertSendCode = `INSERT INTO db_schema.attempt_get_code(phone, code, ip, date)
	VALUES($1, $2, $3, $4) RETURNING id`

const lastAttempt = `SELECT id FROM  db_schema.attempt_get_code
	WHERE position($1 in phone) > 0 AND code = $2 ORDER BY id LIMIT 1`

type AttemptLogin struct {
	Id    int
	Uid   int32
	Code  string
	Phone string
	Ip    string
	Date  time.Time
}

func (s *AttemptLoginService) Last(phone, code string) (id int64, err error) {
	rows, err := s.db.ExecuteSelect(
		lastAttempt,
		phone,
		code,
	)

	if err == nil {
		for rows.Next() {
			rows.Scan(&id)
		}
	}

	return id, err
}

func (s *AttemptLoginService) SendCode(phone, code, ip string) error {

	count, err := s.CountAttemptsGetCodeLastHour(phone, ip)

	if count < maxAttempts && err == nil {
		_, err = s.smsClient.Send(phone, code)
		if err != nil {
			return err
		}

		_, err = s.AddSendCode(phone, code, ip)
		if err != nil {
			return err
		}

	} else if count >= maxAttempts {
		err = errors.New("too many requests, wait a hour")
	}

	return err
}

func (s *AttemptLoginService) lastLoginAttempt(phone, code string) (id int64, err error) {
	rows, err := s.db.ExecuteSelect(
		lastAttempt,
		phone,
		code,
	)

	if err == nil {
		for rows.Next() {
			rows.Scan(&id)
		}
	}

	return id, err
}

func (s *AttemptLoginService) AddSendCode(phone, code, ip string) (id int32, err error) {
	return s.db.Execute(
		insertSendCode,
		phone,
		code,
		ip,
		time.Now(),
	)
}

func (s *AttemptLoginService) CountAttemptsGetCodeLastHour(phone, ip string) (count int32, err error) {
	rows, err := s.db.ExecuteSelect(
		selectAttemptsGetCode,
		phone,
		ip,
		time.Now(),
	)

	if err == nil {
		for rows.Next() {
			rows.Scan(&count)
		}
	}

	return count, err
}
