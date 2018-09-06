package models

import (
	"../services"
	"time"
	"database/sql"
	"fmt"
	"strings"
)

const insertUser = `INSERT INTO db_schema."user"(person, nick_name, avatar, uuid, created_at, country, phone)
	VALUES($1, $2, $3, $4, NOW(), $5, $6) RETURNING id`

const selectUser = `SELECT id, person, nick_name, avatar, uuid, created_at, country, phone, is_confirmed
	FROM db_schema."user" WHERE id = $1`

const selectUsers = `SELECT id, person, nick_name, avatar, created_at 
	FROM db_schema."user" WHERE is_confirmed = 1::bit`

const selectUserByUuid = `SELECT id, person, nick_name, avatar, uuid, created_at, country, phone 
	FROM db_schema."user" WHERE uuid = $1`

const findNickName = `SELECT count(*) AS count FROM db_schema."user" WHERE nick_name = $1`

const findPhone = `SELECT count(*) AS count FROM db_schema."user"
	WHERE position($1 in phone) > 0`

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
	Country		string		`json:"country"`
	Phone		string 		`json:"phone"`
	IsConfirmed rune		`json:"is_confirmed"`
}

func (u *User) Add() (err error) {
	NameExists, _ := u.NickNameExists()
	PhoneExists, _ := u.PhoneNumberExists()

	if !NameExists && !PhoneExists {
		u.Id, err = pg.Execute(
			insertUser,
			u.Person,
			u.NickName,
			u.Avatar,
			u.Uuid,
			u.Country,
			u.Phone,
		)
	}

	return err
}

func (u *User) PhoneNumberExists() (bool, error) {
	var count int
	var err error = nil
	var rows *sql.Rows
	fmt.Println(u.Phone)
	if u.Phone != "" {
		rows, err = pg.ExecuteSelect(findPhone, strings.Trim(u.Phone, " "))
		if err == nil {
			for rows.Next() {
				rows.Scan(&count)
			}
		}
	}

	fmt.Println(count)

	return count > 0, err
}

func (u *User) NickNameExists() (bool, error) {
	var count int
	var err error = nil
	var rows *sql.Rows
	if u.NickName != "" {
		rows, err = pg.ExecuteSelect(findNickName, u.NickName)
		if err == nil {
			for rows.Next() {
				rows.Scan(&count)
			}
		}
	}

	return count > 0, err
}

func (u *User) Exists() (bool, error) {
	var err error = nil
	var rows *sql.Rows
	if u.Uuid != "" {
		rows, err = pg.ExecuteSelect(selectUserByUuid, u.Uuid)

		if err == nil {
			for rows.Next() {
				rows.Scan(
					&u.Id,
					&u.Person,
					&u.NickName,
					&u.Avatar,
					&u.Uuid,
					&u.CreatedAt,
					&u.Country,
					&u.Phone,
				)
			}
		}
	}

	return u.Id != 0, err
}

func (u *User) One(id int64) (err error) {
	var rows *sql.Rows
	rows, err = pg.ExecuteSelect(selectUser, id)

	if err == nil {
		for rows.Next() {
			rows.Scan(
				&u.Id,
				&u.Person,
				&u.NickName,
				&u.Avatar,
				&u.Uuid,
				&u.CreatedAt,
				&u.Country,
				&u.Phone,
				&u.IsConfirmed,
			)
		}
	}

	return err
}

func (u *User) Get(limit, skip int64) (result []User, err error) {
	var rows *sql.Rows

	rows, err = pg.ExecuteSelect(selectUsers)

	if err == nil {
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
	}

	return result, err
}
