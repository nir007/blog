package models

import (
	"time"
	"../services"
	)

const selectAttemptsGetCode = `SELECT count(*) FROM db_schema.attempt_confirm 
	WHERE uid = $1 AND phone = $2 AND date + INTERVAL '1 hour' > $3`

const insertSendCode = `INSERT INTO `

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
	sender := new(services.IQSms)
	_, err := sender.Send(a.Phone, a.Code)

	if err == nil {
		a.AddSendCode(a.Phone, a.Code, a.Ip)
	}

	return err
}

func (a *AttemptLogin) Login (phone, code string) (err error) {


	return err
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

func (a *AttemptLogin) One() (id int32, err error) {
	rows, err := pg.ExecuteSelect(
		oneAttempt,
		a.Uid,
		a.Code,
		a.Phone,
	)

	for rows.Next() {
		rows.Scan(&id)
	}

	return id, err
}

func (a *AttemptLogin) CountAttemptsGetCodeLastHour() (count int32, err error) {

	rows, err := pg.ExecuteSelect(
		selectAttemptsGetCode,
		a.Ip,
		a.Phone,
		time.Now(),
	)

	if err == nil {
		for rows.Next() {
			rows.Scan(&count)
		}
	}

	return count, err
}

func (a *AttemptLogin) CodeExists() (confirmed bool, err error) {
	id, err := a.One()
	return id > 0, err
}