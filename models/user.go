package models

import (
	"../services"
	"time"
	"database/sql"
)

const insertUser = `INSERT INTO public."user"(person, nick_name, avatar, uuid, created_at)
	VALUES($1, $2, $3, $4, NOW()) RETURNING id`

const selectUser = `SELECT id, person, nick_name, avatar, uuid, created_at 
	FROM public."user" WHERE id = $1`

const selectUsers = `SELECT id, person, nick_name, avatar, created_at 
	FROM public."user"`

const selectUserByUuid = `SELECT id, person, nick_name, avatar, uuid, created_at 
	FROM public."user" WHERE uuid = $1`

const findNickName = `SELECT count(*) AS count FROM public."user" WHERE nick_name = $1`

var pg services.Pg

func init() {
	pg = services.Pg{}
}

type User struct {
	Id 			int32		`json:"id"`
	Person 		string		`json:"person"`
	NickName 	string		`json:"nick_name"`
	Avatar 		string		`json:"avatar"`
	Uuid 		string		`json:"uuid"`
	CreatedAt 	time.Time	`json:"created_at"`
	IsOwner		bool		`json:"is_owner"`
}

func (u *User) Add() (err error) {
	if exists, _ := u.NickNameExists(); !exists {
		u.Id, err = pg.Execute(
			insertUser,
			u.Person,
			u.NickName,
			u.Avatar,
			u.Uuid,
		)
	}

	return err
}

func (u *User) NickNameExists() (bool, error) {
	var count int
	var err error = nil
	var rows *sql.Rows
	if u.NickName != "" {
		rows, err = pg.ExecuteSelect(findNickName, u.NickName)
		for rows.Next() {
			rows.Scan(&count)
		}
	}

	return count > 0, err
}

func (u *User) Exists() bool {
	if u.Uuid != "" {
		rows, _:= pg.ExecuteSelect(selectUserByUuid, u.Uuid)

		for rows.Next() {
			rows.Scan(
				&u.Id,
				&u.Person,
				&u.NickName,
				&u.Avatar,
				&u.Uuid,
				&u.CreatedAt,
			)
		}

		if u.Id != 0 {
			return true
		}
	}

	return false
}

func (u *User) One(id int64) {
	rows, _:= pg.ExecuteSelect(selectUser, id)

	for rows.Next() {
		rows.Scan(
			&u.Id,
			&u.Person,
			&u.NickName,
			&u.Avatar,
			&u.Uuid,
			&u.CreatedAt,
		)
	}
}

func (u *User) Get(limit, skip int64) (result []User) {
	rows, _:= pg.ExecuteSelect(selectUsers)

	for rows.Next() {
		u := User{}

		rows.Scan(
			&u.Id,
			&u.Person,
			&u.NickName,
			&u.Avatar,
			&u.CreatedAt,
		)

		result = append(result, u)
	}

	return result
}
