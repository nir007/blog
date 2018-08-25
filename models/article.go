package models

import (
	"../services"
	"time"
	"encoding/json"
	"fmt"
	"database/sql"
	"strings"
	"regexp"
)

const insertArticle = `INSERT INTO db_schema.article(author_id, title, text, tags, created_at, published)
	VALUES($1, $2, $3, $4, NOW(), $5) RETURNING id`

const selectArticles = `SELECT id, author_id, title, text, tags, created_at, published 
	FROM db_schema.article WHERE published = 1::bit ORDER BY created_at DESC LIMIT $1 OFFSET $2`

const selectArticlesByTag = `SELECT id, author_id, title, text, tags, created_at, published 
	FROM db_schema.article WHERE tags?$1 AND published = 1::bit ORDER BY created_at DESC LIMIT $2 OFFSET $3`

const selectArticle = `SELECT id, author_id, title, text, tags, created_at, published 
	FROM db_schema.article WHERE id = $1`

const selectTags = `SELECT DISTINCT tags 
	FROM db_schema.article WHERE published = 1::bit`

const selectForAuthor = `SELECT id, author_id, title, text, tags, created_at, published 
	FROM db_schema.article WHERE author_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3`

const updateArticle = `UPDATE db_schema.article SET title = $1, text = $2, tags = $3, published = $4
	WHERE id = $5 AND author_id = $6 RETURNING id`

type Article struct {
	Id        int32     		`json:"id,string,omitempty"`
	AuthorId  int32     		`json:"author_id,string,omitempty"`
	Title     string    		`json:"title"`
	Text      string			`json:"text"`
	CreatedAt time.Time 		`json:"created"`
	Tags      map[string]string	`json:"tags"`
	Published rune				`json:"published"`
	IsOwner bool                `json:"is_owner"`
}

func (a *Article) Add() {
	pg := services.Pg{}

	tags := map[string]string {}
	var reg = regexp.MustCompile(`[#|$|%|^|&|*|(|)|@|!|?|>|<|/]`)

	for k, _ := range a.Tags {
		key := strings.Replace(k, " ", "", 1000)
		key = string(reg.ReplaceAllString(key, ""))
		tags[fmt.Sprintf("%s", key)] = k
	}

	tagsJson, err := json.Marshal(tags)

	if err != nil {
		tagsJson = nil
	}

	a.Id, err = pg.Execute(
		insertArticle,
		a.AuthorId,
		a.Title,
		a.Text,
		tagsJson,
		a.Published,
	)
}

func (a *Article) Update() (id int32, err error ) {
	pg := services.Pg{}
	tags := map[string]string {}
	var reg = regexp.MustCompile(`[#|$|%|^|&|*|(|)|@|!|?|>|<|/]`)

	for k, _ := range a.Tags {
		key := strings.Replace(k, " ", "", 1000)
		key = string(reg.ReplaceAllString(key, ""))
		tags[fmt.Sprintf("%s", key)] = k
	}

	var tagsJson []byte
	var errParse error

	tagsJson, errParse = json.Marshal(tags)

	if errParse != nil {
		tagsJson = nil
	}

	id, err = pg.Execute(
		updateArticle,
		a.Title,
		a.Text,
		tagsJson,
		a.Published,
		a.Id,
		a.AuthorId,
	)

	return id, err
}

func (a *Article) Get(authorId, perPage, skip int64, tag string) (result []Article, err error) {
	pg := services.Pg{}
	var rows *sql.Rows

	if perPage == 0 {
		perPage = 10
	}

	if authorId > 0 {
		rows, err = pg.ExecuteSelect(selectForAuthor, authorId, perPage, skip)
	} else if tag != "" {
		rows, err = pg.ExecuteSelect(selectArticlesByTag, tag, perPage, skip)
	} else {
		rows, err = pg.ExecuteSelect(selectArticles, perPage, skip)
	}

	if err == nil {
		for rows.Next() {
			article := Article{}
			var id int
			var authorId int
			var title string
			var text string
			var tags sql.NullString
			var createdAt time.Time
			var published rune

			err = rows.Scan(
				&id,
				&authorId,
				&title,
				&text,
				&tags,
				&createdAt,
				&published,
			)

			tagsMap := map[string]string {}

			if tags.Valid {
				json.Unmarshal([]byte(tags.String), &tagsMap)
				article.Tags = tagsMap
			}

			article.Id = int32(id)
			article.AuthorId = int32(authorId)
			article.Title = title
			article.Text = text
			article.Tags = tagsMap
			article.CreatedAt = createdAt
			article.Published = published

			result = append(result, article)
		}
	}

	return result, err
}

func (a *Article) One(id int64) (err error) {
	pg := services.Pg{}
	var rows *sql.Rows
	var nullTags sql.NullString

	rows, err = pg.ExecuteSelect(selectArticle, id)

	for rows.Next() {
		rows.Scan(
			&a.Id,
			&a.AuthorId,
			&a.Title,
			&a.Text,
			&nullTags,
			&a.CreatedAt,
			&a.Published,
		)

		tagsMap := map[string]string {}

		if nullTags.Valid {
			json.Unmarshal([]byte(nullTags.String), &tagsMap)
			a.Tags = tagsMap
		}
	}

	return err
}

func (a *Article) GetTags() (result map[string] string, err error) {
	pg := services.Pg{}
	rows, err := pg.ExecuteSelect(selectTags)

	if err == nil {
		for rows.Next() {
			var jsonVal string
			rows.Scan(&jsonVal)
			json.Unmarshal([]byte(jsonVal), &result)
		}
	}

	return result, err
}