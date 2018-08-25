package services

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/kylelemons/go-gypsy/yaml"
	"fmt"
	"strings"
)

var cs string
var db *sql.DB
var schema string

func init() {
	conf, err := yaml.ReadFile("./config/database.yaml")

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
}

type Pg struct {}

func (p *Pg) setSchema(queryBefore string) (queryAfter string) {
	return strings.Replace(queryBefore, "db_schema", schema, 1)
}

func (p *Pg) Execute(query string, args ...interface{}) (id int32, err error) {
	defer func() {
		if dbErr := db.Ping(); dbErr != nil {
			err = dbErr
		}
	}()

	stmt, err := db.Prepare(p.setSchema(query))

	if err == nil {
		row := stmt.QueryRow(args...)
		row.Scan(&id)
	}

	return id, err
}

func (p *Pg) ExecuteSelect(query string, args ...interface{}) (rows *sql.Rows, err error){
	if err = db.Ping(); err != nil {
		fmt.Println(err)
		return rows, err
	}

	stmt, err := db.Prepare(p.setSchema(query))

	if err != nil {
		return rows, err
	}

	defer stmt.Close()

	rows, err = stmt.Query(args...)

	return rows, err
}