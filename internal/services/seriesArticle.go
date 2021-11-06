package services

import (
	"fmt"
	"github.com/nir007/blog/internal/contracts"
)

type SeriesArticleService struct {
	db contracts.DatabaseFucker
}

func NewSeriesArticleService(db contracts.DatabaseFucker) *SeriesArticleService {
	return &SeriesArticleService{
		db: db,
	}
}

const deleteRelations = `DELETE FROM db_schema.series_article
	WHERE series_id = $1`

const insertRelations = `INSERT INTO db_schema.series_article
	(series_id, article_id, order_num) VALUES `

type SeriesArticle struct {
	Id        int `json:"id"`
	SeriesId  int `json:"series_id"`
	ArticleId int `json:"article_id"`
	Order     int `json:"order"`
}

//update or insert relations series and articles
func (s *SeriesArticleService) Rebuild(sId int32, articles []Article) (err error) {
	_, err = s.db.Execute(deleteRelations, sId)

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
			_, err = s.db.Execute(insertRelations + values[0:len(values)-1])
		}
	}

	return err
}
