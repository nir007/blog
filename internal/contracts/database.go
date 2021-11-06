package contracts

import "database/sql"

type DatabaseFucker interface {
	Execute(query string, args ...interface{}) (id int32, err error)
	ExecuteSelect(query string, args ...interface{}) (rows *sql.Rows, err error)
}
