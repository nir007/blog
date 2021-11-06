package services

import "github.com/nir007/blog/internal/contracts"

type EventService struct {
	smsClient contracts.Sms
	db contracts.DatabaseFucker
}

func NewEventService(db contracts.DatabaseFucker, smsClient contracts.Sms) *EventService{
	return &EventService{
		db: db,
		smsClient: smsClient,
	}
}

const addGroup = `INSERT INTO db_schema.groups(id, author_id, name, description, created_at, deleted_at)
	VALUES($1, $2, $3, $4, now()) RETURNING id
`

const addEvent = `INSERT INTO db_schema.events(id, group_id, author_id, name, description, event_date, created_at, deleted_at)
	VALUES($1, $2, $3, $4, $5, $6, now()) RETURNING id
`