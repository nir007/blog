package models

import (
	"fmt"
	"../services"
)

const deleteRelations = `DELETE FROM db_schema.series_article
	WHERE series_id = $1`

const insertRelations = `INSERT INTO db_schema.series_article
	(series_id, article_id, order_num) VALUES `

type SeriesArticle struct {
	Id int        `json:"id"`
	SeriesId int  `json:"series_id"`
	ArticleId int `json:"article_id"`
	Order int     `json:"order"`
}

//update or insert relations series and articles
func (s *SeriesArticle) Rebuild(sId int32, articles []Article) (err error) {
	pg := new(services.Pg)
	_, err = pg.Execute(deleteRelations, sId)

	if err == nil {
		values := ""
		for _, a := range articles {
			values += fmt.Sprintf(
				"(%d, %d, %d),",
				sId,
				a.Id,
				a.Order,
			)
		}
		if len(values) > 0 {
			_, err = pg.Execute(insertRelations + values[0: len(values) - 1])
		}
	}

	return err
}

