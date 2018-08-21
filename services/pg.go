package services

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"github.com/kylelemons/go-gypsy/yaml"
	"fmt"
)

var cs string
var db *sql.DB

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

	str := "user=%s dbname=%s host=%s port=%d password=%s sslmode=%s"
	cs = fmt.Sprintf(str, user, dbName, host, port, password, sslMode)

	db, err = sql.Open("postgres", cs)

	defer func() {
		if err := recover(); err != nil {
			db.Close()
			fmt.Println(err)
		}
	}()

	if err != nil {
		panic(fmt.Errorf("connection failed: %T", err.Error()))
	}
}

type Pg struct {}

func (p *Pg) Execute(query string, args ...interface{}) (id int32, err error) {
	defer func() {
		if dbErr := db.Ping(); dbErr != nil {
			err = dbErr
		}
	}()

	stmt, err := db.Prepare(query)

	if err == nil {
		row := stmt.QueryRow(args...)
		row.Scan(&id)
	}

	return id, err
}

func (p *Pg) ExecuteSelect(query string, args ...interface{}) (rows *sql.Rows, err error){
	if err = db.Ping(); err != nil {
		return rows, err
	}

	stmt, err := db.Prepare(query)

	if err != nil {
		log.Fatal("Prepare failed:", err.Error())
	}
	defer stmt.Close()

	rows, err = stmt.Query(args...)

	return rows, err
}