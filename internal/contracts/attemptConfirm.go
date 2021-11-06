package contracts

import "time"

type AttemptConfirm struct {
	Id    int
	Uid   int32
	Code  string
	Phone string
	Date  time.Time
}

type AttemptConformer interface {
	Add(a *AttemptConfirm) (id int32, err error)
	One(attempt *AttemptConfirm) (id int32, err error)
	CountLastHour(uid int32, phone string) (count int32, err error)
	CodeExists(attempt *AttemptConfirm) (confirmed bool, err error)
}
