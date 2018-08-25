package models

import (
	"../services"
	"database/sql"
)

const create = `INSERT INTO db_schema.series
	(title, description, published) RETURNING id`

const read = `SELECT * FROM db_schema.series 
	WHERE author_id = $1`

const one = `SELECT * FROM db_schema.series 
	WHERE id = $1 AND author_id = $2`

const updateSeries = `UPDATE db_schema.series 
	SET title = $1, description = $2, published = $3`

const deleteSeries = `DELETE db_schema.series 
	WHERE id = $1 AND author_id = $2`

type Series struct {
	Id int             `json:"id"`
	AuthorId int       `json:"author_id"`
	Count int          `json:"count"`
	Title string       `json:"title"`
	Description string `json:"description"`
	Published string   `json:"published"`
}

func (s * Series) Create() (id int32, err error) {
	pg := services.Pg{}
	return pg.Execute(
		create,
		s.AuthorId,
		s.Title,
		s.Description,
		s.Published,
	)
}

func (s *Series) Read() (rows *sql.Rows, err error) {
	pg := services.Pg{}
	return pg.ExecuteSelect(read, s.AuthorId)
}

func (s *Series) One(id int) (rows *sql.Rows, err error) {
	pg := services.Pg{}
	return pg.ExecuteSelect(one, s.Id, s.AuthorId)
}

func (s *Series) Update() (id int32, err error) {
	pg := services.Pg{}
	return pg.Execute(
		updateSeries,
		s.Title,
		s.Description,
		s.Published,
	)
}

func (s *Series) Delete() (id int32, err error) {
	pg := services.Pg{}
	return pg.Execute(
		deleteSeries,
		s.Title,
		s.Description,
		s.Published,
	)
}
