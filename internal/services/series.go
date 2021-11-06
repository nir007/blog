package services

import (
	"database/sql"
	"github.com/nir007/blog/internal/contracts"
)

type SeriesService struct {
	db contracts.DatabaseFucker
}

func NewSeriesService(db contracts.DatabaseFucker) *SeriesService {
	return &SeriesService{
		db: db,
	}
}

const createSeries  = `INSERT INTO db_schema.series
	(author_id, title, description, published) 
	VALUES($1, $2, $3, $4) RETURNING id`

const getSeries = `SELECT id, author_id, count, title, description, published
	FROM db_schema.series 
	WHERE author_id = $1 ORDER BY published DESC LIMIT $2 OFFSET $3`

const one = `SELECT id, author_id, title, description, published, count
	FROM db_schema.series WHERE id = $1 LIMIT 1`

const updateSeries = `UPDATE db_schema.series 
	SET title = $2, description = $3, published = $4, count = $5 
	WHERE id = $1 RETURNING id`

const deleteSeries = `DELETE FROM db_schema.series CASCADE
	WHERE id = $1 AND author_id = $2 RETURNING id`

const articles = `SELECT a.id, a.author_id, a.title, a.text, a.created_at, a.published, rel.order_num 
	FROM db_schema.series_article AS rel
	LEFT JOIN db_schema.article AS a ON rel.article_id = a.id 
	WHERE series_id = $1
	ORDER BY rel.order_num`

type Series struct {
	Id          int32     `json:"id"`
	AuthorId    int32     `json:"author_id"`
	Count       int64     `json:"count"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Published   rune      `json:"published"`
	Articles    []Article `json:"articles"`
}

func (s *SeriesService) Create(create *Series) (id int32, err error) {
	return s.db.Execute(
		createSeries,
		create.AuthorId,
		create.Title,
		create.Description,
		create.Published,
	)
}

func (s *SeriesService) Read(authorId int32, limit, offset int64) (result []Series, err error) {
	rows, err := s.db.ExecuteSelect(getSeries, authorId, limit, offset)
	if err != nil {
		return nil, err
	}

	var id int32
	var count sql.NullInt64
	var title sql.NullString
	var description sql.NullString
	var published rune

	for rows.Next() {
		err = rows.Scan(
			&id,
			&authorId,
			&count,
			&title,
			&description,
			&published,
		)
		if err != nil {
			return nil, err
		}

		if !count.Valid {
			count.Int64 = 0
		}

		if !title.Valid {
			title.String = "Untitled"
		}

		if !description.Valid {
			description.String = "Some series"
		}

		item := Series{
			Id:          id,
			AuthorId:    authorId,
			Count:       count.Int64,
			Title:       title.String,
			Description: description.String,
			Published:   published,
		}
		result = append(result, item)
	}

	return result, nil
}

func (s *SeriesService) One(id int64) (res *Series, err error) {
	res  = &Series{}
	rows, err := s.db.ExecuteSelect(one, id)
	if err != nil {
		return res, err
	}

	for rows.Next() {
		rows.Scan(
			&res.Id,
			&res.AuthorId,
			&res.Title,
			&res.Description,
			&res.Published,
			&res.Count,
		)
		break
	}

	if res.Id > 0 {
		rows, err = s.db.ExecuteSelect(articles, res.Id)

		if err == nil {
			for rows.Next() {
				article := Article{}
				if err = rows.Scan(
					&article.Id,
					&article.AuthorId,
					&article.Title,
					&article.Text,
					&article.CreatedAt,
					&article.Published,
					&article.Order,
				); err != nil {
					return nil, err
				}
				res.Articles = append(res.Articles, article)
			}
		}
	}

	return res, err
}

func (s *SeriesService) Update(update *Series) (id int32, err error) {
	return s.db.Execute(
		updateSeries,
		update.Id,
		update.Title,
		update.Description,
		update.Published,
		update.Count,
	)
}

func (s *SeriesService) Delete(sId, authorId int64) (id int32, err error) {
	return s.db.Execute(deleteSeries, sId, authorId)
}
