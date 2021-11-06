package services

import (
	"database/sql"
	"errors"
	"github.com/nir007/blog/internal/contracts"
	"math/rand"
	"strconv"
	"time"
)

type UserService struct {
	db                    contracts.DatabaseFucker
	smsClient             contracts.Sms
	attemptConfirmService contracts.AttemptConformer
	loggedUsers           map[string]User
}

func NewUserService(
	db contracts.DatabaseFucker,
	smsClient contracts.Sms,
	attemptConfirmService contracts.AttemptConformer,
) *UserService {
	return &UserService{
		db:                    db,
		smsClient:             smsClient,
		attemptConfirmService: attemptConfirmService,
		loggedUsers:           make(map[string]User, 20),
	}
}

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

func (s *UserService) Add(create *User) (err error) {
	rand.Seed(time.Now().UTC().UnixNano())
	code := strconv.Itoa(rand.Intn(100))
	code += strconv.Itoa(rand.Intn(1000))

	nameExists, err := s.NickNameExists(create.NickName)
	if err != nil {
		return err
	}

	phoneExists, err := s.PhoneNumberExists(create.Phone)
	if err != nil {
		return err
	}

	if !nameExists && !phoneExists {

		create.Id, err = s.db.Execute(
			insertUser,
			create.Person,
			create.NickName,
			create.Avatar,
			create.Uuid,
			create.Country,
			create.Phone,
		)

		if err == nil && create.Id > 0 {
			_, err = s.smsClient.Send(create.Phone, code)
			if err != nil {
				return errors.New("fail with sending confirmation code")
			}

			_, err = s.attemptConfirmService.Add(&contracts.AttemptConfirm{
				Uid:   create.Id,
				Code:  code,
				Phone: create.Phone,
				Date:  time.Now(),
			})
		}
	}

	return err
}

func (s *UserService) PhoneNumberExists(phone string) (bool, error) {
	var count int
	var err error = nil
	var rows *sql.Rows

	if phone != "" {
		rows, err = s.db.ExecuteSelect(findPhone, phone)

		if err == nil {
			for rows.Next() {
				rows.Scan(&count)
			}
		}
	}

	return count > 0, err
}

func (s *UserService) NickNameExists(name string) (bool, error) {
	var count int
	var err error = nil
	var rows *sql.Rows

	if name != "" {
		rows, err = s.db.ExecuteSelect(findNickName, name)
		if err == nil {
			for rows.Next() {
				rows.Scan(&count)
			}
		}
	}

	return count > 0, err
}

func (s *UserService) Exists(user *User) (bool, error) {
	var err error = nil
	var rows *sql.Rows

	if user.Uuid != "" {
		rows, err = s.db.ExecuteSelect(selectUserByUuid, user.Uuid)

		if err == nil {
			for rows.Next() {
				rows.Scan(
					&user.Id,
					&user.Person,
					&user.NickName,
					&user.Avatar,
					&user.Uuid,
					&user.CreatedAt,
					&user.Country,
					&user.Phone,
					&user.IsConfirmed,
				)
			}
		}
	}

	return user.Id != 0, err
}

func (s *UserService) One(id int64) (user User, err error) {
	var rows *sql.Rows
	rows, err = s.db.ExecuteSelect(selectUser, id)
	if err != nil {
		return User{}, err
	}

	for rows.Next() {
		rows.Scan(
			&user.Id,
			&user.Person,
			&user.NickName,
			&user.Avatar,
			&user.Uuid,
			&user.CreatedAt,
			&user.Country,
			&user.Phone,
			&user.IsConfirmed,
		)
	}

	return user, err
}

func (s *UserService) OneByPhone(phone string) (user User, err error) {
	var rows *sql.Rows
	rows, err = s.db.ExecuteSelect(
		selectUserByPhone,
		phone,
	)
	if err != nil {
		return User{}, err
	}

	for rows.Next() {
		rows.Scan(
			&user.Id,
			&user.Person,
			&user.NickName,
			&user.Avatar,
			&user.Uuid,
			&user.CreatedAt,
			&user.Country,
			&user.Phone,
			&user.IsConfirmed,
		)
	}

	return user, err
}

func (s *UserService) Get(limit, skip int64) (result []User, err error) {
	var rows *sql.Rows

	rows, err = s.db.ExecuteSelect(selectUsers)
	if err != nil {
		return nil, err
	}

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

	return result, err
}

func (s *UserService) ConfirmPhone(user User, code string) (confirmed bool, err error) {
	attempt := &contracts.AttemptConfirm{
		Uid:   user.Id,
		Code:  code,
		Phone: user.Phone,
	}

	codeExists, err := s.attemptConfirmService.CodeExists(attempt)

	if err == nil && codeExists {
		id, err := s.db.Execute(
			setConfirmed,
			user.Id,
			user.Phone,
		)

		confirmed = id > 0 && err == nil
	}

	return confirmed, err
}

func (s *UserService) GetLoggedUser(id string) (user User, err error) {
	var exists bool

	uid, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return User{}, err
	}

	if exists, err = s.Exists(&User{Id: int32(uid)}); err == nil && exists {
		s.loggedUsers[id] = user
	}

	return user, err
}

func (s *UserService) IsLogged(id string) bool {
	_, ok := s.loggedUsers[id]

	return ok
}

func (s *UserService) SetLoggedUser(user User) {
	s.loggedUsers[user.Uuid] = user
}
