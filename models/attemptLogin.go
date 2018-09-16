package models

import (
	"time"
	"../services"
	"errors"
)

const selectAttemptsGetCode = `SELECT count(*) FROM db_schema.attempt_get_code 
	WHERE position($1 in phone) > 0 AND ip = $2 AND date + INTERVAL '1 hour' > $3`

const insertSendCode = `INSERT INTO db_schema.attempt_get_code(phone, code, ip, date)
	VALUES($1, $2, $3, $4) RETURNING id`

const lastAttempt = `SELECT id FROM  db_schema.attempt_get_code
	WHERE position($1 in phone) > 0 AND code = $2 ORDER BY id LIMIT 1`

func init() {
	pg = services.Pg{}
}

type AttemptLogin struct {
	Id    int
	Uid   int32
	Code  string
	Phone string
	Ip    string
	Date  time.Time
}

func (a *AttemptLogin) SendCode() error {
	var maxAttempts int32 = 3
	count, err := a.CountAttemptsGetCodeLastHour()

	if count < maxAttempts && err == nil {
		sender := new(services.IQSms)
		sender.SetFromConfig()
		_, err := sender.Send(a.Phone, a.Code)

		if err == nil {
			a.AddSendCode(a.Phone, a.Code, a.Ip)
		}
	} else if count >= maxAttempts {
		err = errors.New("too many requests, hold a hour")
	}

	return err
}

func (a *AttemptLogin) Last() (id int64, err error) {
	rows, err := pg.ExecuteSelect(
		lastAttempt,
		a.Phone,
		a.Code,
	)

	if err == nil {
		for rows.Next() {
			rows.Scan(&id)
		}
	}

	return id, err
}

func (a *AttemptLogin) AddSendCode(phone, code, ip string) (id int32, err error) {
	return pg.Execute(
		insertSendCode,
		phone,
		code,
		ip,
		time.Now(),
	)
}

func (a *AttemptLogin) CountAttemptsGetCodeLastHour() (count int32, err error) {
	rows, err := pg.ExecuteSelect(
		selectAttemptsGetCode,
		a.Phone,
		a.Ip,
		time.Now(),
	)

	if err == nil {
		for rows.Next() {
			rows.Scan(&count)
		}
	}

	return count, err
}