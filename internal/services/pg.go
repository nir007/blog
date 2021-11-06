package services

import (
	"database/sql"
	_ "github.com/lib/pq"
	"strings"
)

type PGDatabaseFucker struct {
	db     *sql.DB
	schema string
}

func NewPGDatabaseFucker(db *sql.DB, schema string) *PGDatabaseFucker {
	return &PGDatabaseFucker{
		db:     db,
		schema: schema,
	}
}

func init() {
	/*conf, err := yaml.ReadFile("./config/database.yaml")

	if err != nil {
		println("Пизда рулю: ", err)
	}

	user, _ := conf.Get("user")
	dbName, _ := conf.Get("dbname")
	host, _ := conf.Get("host")
	port, _ := conf.GetInt("port")
	password, _ := conf.Get("password")
	sslMode, _ := conf.Get("sslmode")
	schema, _ = conf.Get("schema")

	str := "user=%s dbname=%s host=%s port=%d password=%s sslmode=%s"
	cs = fmt.Sprintf(str, user, dbName, host, port, password, sslMode)

	db, err = sql.Open("postgres", cs)
	*/
}

func (p *PGDatabaseFucker) setSchema(queryBefore string) (queryAfter string) {
	return strings.Replace(queryBefore, "db_schema", p.schema, -1)
}

func (p *PGDatabaseFucker) Execute(query string, args ...interface{}) (id int32, err error) {
	defer func() {
		if dbErr := p.db.Ping(); dbErr != nil {
			err = dbErr
		}
	}()

	stmt, err := p.db.Prepare(p.setSchema(query))

	if err == nil {
		row := stmt.QueryRow(args...)
		row.Scan(&id)
	}

	return id, err
}

func (p *PGDatabaseFucker) ExecuteSelect(query string, args ...interface{}) (rows *sql.Rows, err error) {
	if err = p.db.Ping(); err != nil {
		return rows, err
	}

	stmt, err := p.db.Prepare(p.setSchema(query))

	if err != nil {
		return rows, err
	}

	defer stmt.Close()

	rows, err = stmt.Query(args...)

	return rows, err
}
