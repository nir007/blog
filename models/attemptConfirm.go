package models

import "time"

type AttemptConfirm struct {
	id    int
	uid   int
	code  string
	phone string
	date  time.Time
}

func (a *AttemptConfirm) Add() {

}

func (a *AttemptConfirm) CountLastHour() {

}

