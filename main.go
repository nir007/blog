package main

import (
	"./models"
	"encoding/json"
	"github.com/gorilla/mux"
	tpl "html/template"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"time"
	"fmt"
)

const cookieNameId = "uuid"
const contentTypeJson = "application/json"

var loggedUsers = make(map[string]models.User, 10)

func main() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public/dist/static"))
	rp := router.PathPrefix("/static/")
	rp.Handler(http.StripPrefix("/static/", fs))

	router.HandleFunc("/", indexAction)

	router.HandleFunc("/aj_add_user", ajAddUserAction)
	router.HandleFunc("/aj_get_tags", ajGetTags)
	router.HandleFunc("/aj_sign_in", ajSignInAction)
	router.HandleFunc("/aj_get_check_nickname", ajCheckNickNameAction)
	router.HandleFunc("/aj_is_logged", ajIsLoggedAction)

	router.HandleFunc("/get_articles", ajGetArticlesAction)
	router.HandleFunc("/aj_get_article", ajGetArticleAction)
	router.HandleFunc("/aj_add_article", ajAddArticleAction)
	router.HandleFunc("/aj_update_article", ajUpdateArticleAction)
	router.HandleFunc("/aj_get_person", ajGetPersonAction)
	router.HandleFunc("/aj_get_persons", ajGetPersonsAction)
	http.ListenAndServe(":81", router)
}

func indexAction(w http.ResponseWriter, r *http.Request) {
	t := tpl.Must(tpl.ParseFiles(
		"./public/dist/index.html",
	))
	errTpl := t.Execute(w, nil)

	if errTpl != nil {
		panic(errTpl)
	}
}

func getLoggedUser(uuid string) (user models.User, err error) {
	user.Uuid = uuid
	var exists bool

	if exists, err = user.Exists(); err == nil && exists {
		loggedUsers[user.Uuid] = user
	}

	return user, err
}

func ajIsLoggedAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	uid, err := r.Cookie("uuid")

	if err == nil {
		_, err2 := getLoggedUser(uid.Value)

		if err2 != nil {
			response.Status = 500
			response.Data = "Какая то хуйня"
		}
	}

	if uid != nil {
		if u, ok := loggedUsers[uid.Value]; ok {
			response.Status = 200
			response.Data = u
		}
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func ajGetTags(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	article := models.Article{}

	if tags, err := article.GetTags(); err != nil {
		response.Status = 500
		response.Data = err
	} else {
		response.Status = 200
		response.Data = tags
	}

	w.Header().Set("ContentType", contentTypeJson)
	w.Write(response.ToBytes())
}

func ajGetArticleAction(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	response := models.Response{
		Status: 404,
		Data: "Article not found",
	}

	id := r.FormValue("id")
	articleId, _ := strconv.ParseInt(id, 10, 64)

	article := models.Article{}
	err := article.One(articleId)

	if err != nil {
		response.Status = 500
		response.Data = err
	} else {
		uid, err := r.Cookie("uuid")
		if uid != nil && err == nil {
			user, err := getLoggedUser(uid.Value)
			if user.Id == article.AuthorId && err == nil {
				article.IsOwner = true
			}
		}

		complexData := map[string]interface{}{
			"article": article,
		}

		if article.Id > 0 {
			err = user.One(int64(article.AuthorId))
			if err != nil {
				response.Status = 500
				response.Data = err
			} else {
				response.Status = 200
				complexData["author"] = user
				response.Data = complexData
			}
		}
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func ajAddArticleAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	article := models.Article{}
	uid, _ := r.Cookie(cookieNameId)

	if uid != nil {
		if user, ok := loggedUsers[uid.Value]; ok {
			err := json.NewDecoder(r.Body).Decode(&article)

			if err != nil {
				response.Status = 500
				response.Data = err.Error()
			} else {
				newArticle := models.Article{
					AuthorId:  user.Id,
					Title:     article.Title,
					Text:      article.Text,
					Tags:      article.Tags,
					CreatedAt: time.Now(),
					Published: article.Published,
				}
				newArticle.Add()
				response.Status = 200
				response.Data = newArticle
			}
		}
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func ajUpdateArticleAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Status: 500, Data: "Something went wrong"}
	article := models.Article{}

	uid, err := r.Cookie(cookieNameId)

	fmt.Println(uid)

	if uid != nil && err == nil {
		err := json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			response.Status = 500
			response.Data = err.Error()
		} else if user, err := getLoggedUser(uid.Value);
			err == nil && user.Id == article.AuthorId {

			_, err := article.Update()
			if err != nil {
				response.Status = 500
				response.Data = err
			} else {
				response.Status = 200
				response.Data = article
			}
		}
	} else {
		response.Status = 500
		response.Data = "Access deny"
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func ajGetArticlesAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	authorId := r.FormValue("author_id")
	tag := r.FormValue("tag")
	limit := r.FormValue("limit")
	offset := r.FormValue("offset")

	aId, _ := strconv.ParseInt(authorId, 10, 64)
	perPage, _ := strconv.ParseInt(limit, 10, 64)
	skip, _ := strconv.ParseInt(offset, 10, 64)

	article := models.Article{}
	articles, err := article.Get(aId, perPage, skip, tag)

	if err != nil {
		response.Status = 500
		response.Data = err
	} else {
		uid, err := r.Cookie("uuid")

		if aId > 0 && uid != nil && err == nil {
			user, err := getLoggedUser(uid.Value)
			if err == nil {
				for k, v := range articles {
					if user.Id == v.AuthorId {
						articles[k].IsOwner = true
					}
				}
			}
		}

		response.Status = 200
		response.Data = articles
	}

	w.Header().Set("Content-type", contentTypeJson)
	w.Write(response.ToBytes())
}

func ajCheckNickNameAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}

	nickName := r.FormValue("nickname")
	user := models.User{NickName: nickName}

	if exists, err := user.NickNameExists(); err != nil && exists {
		response.Status = 200
		response.Data = true
	} else {
		response.Status = 500
		response.Data = false
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func ajAddUserAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}

	if r.Body == nil {
		http.Error(w, "Body is empty", 400)
		return
	}

	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	uid := uuid.Must(uuid.NewV4())

	newUser := models.User{
		Person:   user.Person,
		NickName: user.NickName,
		Avatar:   user.Avatar,
		Uuid:     uid.String(),
	}

	errAdd := newUser.Add()

	fmt.Println(errAdd)

	if newUser.Id != 0 {
		response.Status = 200
		response.Data = newUser

		cookie := http.Cookie{
			Name:    cookieNameId,
			Value:   newUser.Uuid,
			Expires: time.Now().Add(360 * 24 * time.Hour),
		}

		http.SetCookie(w, &cookie)
		loggedUsers[newUser.Uuid] = newUser
	} else {
		response.Status = 500
		response.Data = errAdd
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func ajSignInAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	uid := r.FormValue("uuid")
	user := models.User{Uuid: uid}

	if exists, err := user.Exists(); err == nil && exists {
		loggedUsers[user.Uuid] = user

		cookie := http.Cookie{
			Name:    cookieNameId,
			Value:   uid,
			Expires: time.Now().Add(360 * 24 * time.Hour),
		}
		http.SetCookie(w, &cookie)
		data, _ := json.Marshal(user)

		response.Status = 200
		response.Data = data

		w.Header().Set("Content-Type", contentTypeJson)
		w.Write(response.ToBytes())
		return
	}

	response.Status = 500
	response.Data = "User not found"

	w.Write(response.ToBytes())
}

func ajGetPersonAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Status: 500, Data: false}
	var user models.User
	id := r.FormValue("id")

	personId, _ := strconv.ParseInt(id, 10, 64)

	uid, err := r.Cookie("uuid")

	if personId == 0 {
		user = models.User{Uuid: uid.Value}

		if exists, err := user.Exists(); err == nil && exists {
			loggedUsers[user.Uuid] = user
		}
	} else {
		user = models.User{}
		user.One(personId)
	}

	if user.Id > 0 {
		if err == nil && uid.Value == user.Uuid {
			user.IsOwner = true
		} else {
			user.Uuid = ""
		}
		response.Status = 200
		response.Data = user
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func ajGetPersonsAction(w http.ResponseWriter, r *http.Request) {
	limit := r.FormValue("limit")
	offset := r.FormValue("offset")

	perPage, _ := strconv.ParseInt(limit, 10, 64)
	skip, _ := strconv.ParseInt(offset, 10, 64)

	user := models.User{}
	response := models.Response{}

	if users, err := user.Get(perPage, skip); err != nil {
		response.Status = 500
		response.Data = err
	} else {
		response.Status = 200
		response.Data = users
	}

	w.Header().Set("Content-type", contentTypeJson)
	w.Write(response.ToBytes())
}
