package models

import (
	"database/sql"
	"errors"
	"math/rand"
	"strconv"
	"time"

	"blog/services"
)

const insertUser = `INSERT INTO db_schema."user"(person, nick_name, avatar, uuid, created_at, country, phone)
	VALUES($1, $2, $3, $4, NOW(), $5, $6) RETURNING id`

const selectUser = `SELECT id, person, nick_name, avatar, uuid, created_at, country, phone, is_confirmed
	FROM db_schema."user" WHERE id = $1`

const selectUsers = `SELECT id, person, nick_name, avatar, created_at 
	FROM db_schema."user" WHERE is_confirmed = 1::bit`

const selectUserByUuid = `SELECT id, person, nick_name, avatar, uuid, created_at, country, phone, is_confirmed
	FROM db_schema."user" WHERE uuid = $1`

const selectUserByPhone = `SELECT id, person, nick_name, avatar, uuid, created_at, country, phone, is_confirmed
	FROM db_schema."user" WHERE position($1 in phone) > 0`

const findNickName = `SELECT count(*) AS count FROM db_schema."user" 
	WHERE nick_name = $1`

const findPhone = `SELECT count(*) AS count FROM db_schema."user"
	WHERE position($1 in phone) > 0`

const setConfirmed = `UPDATE db_schema."user" SET is_confirmed = 1::bit
	WHERE id = $1 AND phone = $2 RETURNING id`

func init() {
	pg = services.Pg{}
}

type User struct {
	Id          int32     `json:"id"`
	Person      string    `json:"person"`
	NickName    string    `json:"nick_name"`
	Avatar      string    `json:"avatar"`
	Uuid        string    `json:"uuid"`
	CreatedAt   time.Time `json:"created_at"`
	IsOwner     bool      `json:"is_owner"`
	Country     string    `json:"country"`
	Phone       string    `json:"phone"`
	IsConfirmed rune      `json:"is_confirmed"`
}

func (u *User) Add() (err error) {
	rand.Seed(time.Now().UTC().UnixNano())
	code := strconv.Itoa(rand.Intn(100))
	code += strconv.Itoa(rand.Intn(1000))

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

		if err == nil && u.Id > 0 {
			sender := services.IQSms{}
			errConfig := sender.SetFromConfig()
			_, errSending := sender.Send(u.Phone, code)

			if errConfig != nil {
				return errors.New("fail config sending confirmation code")
			}

			if errSending != nil {
				return errors.New("fail with sending confirmation code")
			}

			attemptConfirm := AttemptConfirm{
				Uid:   u.Id,
				Code:  string(code),
				Phone: u.Phone,
				Date:  time.Now(),
			}
			attemptConfirm.Add()
		}
	}

	return err
}

func (u *User) PhoneNumberExists() (bool, error) {
	var count int
	var err error = nil
	var rows *sql.Rows

	if u.Phone != "" {
		rows, err = pg.ExecuteSelect(findPhone, u.Phone)

		if err == nil {
			for rows.Next() {
				rows.Scan(&count)
			}
		}
	}

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
					&u.IsConfirmed,
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

func (u *User) OneByPhone(phone string) (err error) {
	var rows *sql.Rows
	rows, err = pg.ExecuteSelect(
		selectUserByPhone,
		phone,
	)

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

func (u *User) ConfirmPhone(code string) (confirmed bool, err error) {
	a := AttemptConfirm{
		Uid:   u.Id,
		Code:  code,
		Phone: u.Phone,
	}

	codeExists, err := a.CodeExists()

	if err == nil && codeExists {
		id, err := pg.Execute(
			setConfirmed,
			u.Id,
			u.Phone,
		)

		confirmed = id > 0 && err == nil
	}

	return confirmed, err
}
