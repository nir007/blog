package models

import (
	"time"

	"blog/services"
)

const insert = `INSERT INTO db_schema.attempt_confirm (uid, code, phone, date)
	VALUES($1, $2, $3, $4)`

const oneAttempt = `SELECT id FROM db_schema.attempt_confirm 
	WHERE uid = $1 AND code = $2 AND phone = $3 ORDER BY id LIMIT 1`

const selectCountLastHour = `SELECT count(*) FROM db_schema.attempt_confirm 
	WHERE uid = $1 AND phone = $2 AND date + INTERVAL '1 hour' > $3`

var pg services.Pg

func init() {
	pg = services.Pg{}
}

type AttemptConfirm struct {
	Id    int
	Uid   int32
	Code  string
	Phone string
	Date  time.Time
}

func (a *AttemptConfirm) Add() (id int32, err error) {
	return pg.Execute(
		insert,
		a.Uid,
		a.Code,
		a.Phone,
		a.Date,
	)
}

func (a *AttemptConfirm) One() (id int32, err error) {
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

func (a *AttemptConfirm) CountLastHour() (count int32, err error) {

	rows, err := pg.ExecuteSelect(
		selectCountLastHour,
		a.Uid,
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

func (a *AttemptConfirm) CodeExists() (confirmed bool, err error) {
	id, err := a.One()
	return id > 0, err
}
