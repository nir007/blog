package models

import (
	"../services"
)

const createRel = `INSERT INTO db_schema.series_article
	(series_id, article_id, order) VALUES($1, $2, $3)`

const updateOrders = `UPDATE db_schema.series_article
	SET `

const deleteRel = `DELETE FROM db_schema.series_article CASCADE
	WHERE id = $1`

type SeriesArticle struct {
	Id int        `json:"id"`
	SeriesId int  `json:"series_id"`
	ArticleId int `json:"article_id"`
	Order int     `json:"order"`
}

func (s * SeriesArticle) Create() (id int32, err error) {
	pg := new(services.Pg)
	return pg.Execute(
		createRel,
		s.SeriesId,
		s.ArticleId,
		s.Order,
	)
}

func (s *SeriesArticle) update()  {
	pg := new(services.Pg)
	pg.Execute(updateOrders, s.ArticleId, s.Order)
}

func (s *SeriesArticle) Delete() (int32, error) {
	pg := new(services.Pg)
	return pg.Execute(deleteRel, s.Id)
}
