package main

import (
	"./helpers"
	"./models"
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/satori/go.uuid"
	tpl "html/template"
	"net/http"
	"strconv"
	"time"
)

const cookieNameId = "uuid"
const contentTypeJson = "application/json"

var helper = new(helpers.Helper)
var loggedUsers = make(map[string]models.User, 10)

func main() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public/dist/static"))
	rp := router.PathPrefix("/static/")
	rp.Handler(http.StripPrefix("/static/", fs))

	router.HandleFunc("/", indexAction)
	router.HandleFunc("/logout", logout)
	router.HandleFunc("/aj_add_user", addUserAction)
	router.HandleFunc("/aj_get_tags", getTags)
	router.HandleFunc("/aj_sign_in", signInAction)
	router.HandleFunc("/aj_get_check_nickname", checkNickNameAction)
	router.HandleFunc("/aj_is_logged", isLoggedAction)
	router.HandleFunc("/aj_confirm_phone", confirmPhoneAction)
	router.HandleFunc("/aj_get_check_phone", checkPhoneNumberAction)
	router.HandleFunc("/create_series", createSeriesAction)
	router.HandleFunc("/get_one_series", getOneSeriesAction)
	router.HandleFunc("/delete_series", deleteSeriesAction)
	router.HandleFunc("/update_series", updateSeriesAction)
	router.HandleFunc("/get_user_series", getUserSeries) 
	router.HandleFunc("/get_articles", getArticlesAction)
	router.HandleFunc("/get_published_articles", getPublishedArticles)
	router.HandleFunc("/aj_get_article", getArticleAction)
	router.HandleFunc("/remove_article", removeArticleAction)
	router.HandleFunc("/aj_add_article", addArticleAction)
	router.HandleFunc("/aj_update_article", updateArticleAction)
	router.HandleFunc("/aj_get_person", getPersonAction)
	router.HandleFunc("/aj_get_persons", getPersonsAction)
	router.HandleFunc("/aj_get_code_to_login", getCodeToLoginAction)

	go http.ListenAndServe(":80", http.HandlerFunc(redirect))
	http.ListenAndServeTLS(":443", "rakan-tarakan.com.crt", "private.key", router)
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req,
		"https://" + req.Host + req.URL.String(),
		http.StatusMovedPermanently)
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

func createSeriesAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Status: 500, Data: "some error"}
	uid, err := r.Cookie(cookieNameId)

	if err == nil {
		if user, err := getLoggedUser(uid.Value); err == nil && user.Id > 0 {

			series := new(models.Series)
			errDecode := json.NewDecoder(r.Body).Decode(&series)

			if errDecode != nil {
				response.Status = 500
				response.Data = errDecode
			} else {
				series.AuthorId = user.Id
				if id, err := series.Create(); err != nil {
					response.Status = 500
					response.Data = err.Error()
				} else {
					response.Status = 200
					response.Data = id
				}
			}
		}
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func getOneSeriesAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Status: 500, Data: "Some error"}

	seriesId := r.FormValue("series_id")
	sId, _ := strconv.ParseInt(seriesId, 10, 32)

	series := new(models.Series)
	err := series.One(sId)

	if err != nil {
		response.Status = 500
		response.Data = err
	} else if series.Id == 0 {
		response.Status = 404
		response.Data = "Series not found"
	} else {
		response.Status = 200
		response.Data = series
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func deleteSeriesAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Status: 500, Data: "Some error"}
	seriesId := r.FormValue("id")
	sId, _ := strconv.ParseInt(seriesId, 10, 32)

	uid, err := r.Cookie(cookieNameId)

	if err == nil {
		if user, err := getLoggedUser(uid.Value); err == nil && user.Id > 0 {
			series := new(models.Series)
			id, err := series.Delete(sId, int64(user.Id))

			if err != nil {
				response.Status = 500
				response.Data = err
			} else {
				response.Status = 200
				response.Data = id
			}
		}
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func updateSeriesAction(w http.ResponseWriter, r *http.Request)  {
	response := models.Response{Status: 500, Data: "Some error"}
	series := new(models.Series)

	err := json.NewDecoder(r.Body).Decode(&series)

	if err != nil {
		response.Status = 500
		response.Data = err
	} else {
		seriesId, err := series.Update()
		if err != nil {
			response.Status = 500
			response.Data = err
		} else if seriesId == 0 {
			response.Status = 404
			response.Data = "Series not found"
		} else {
			seriesArticles := new(models.SeriesArticle)
			err := seriesArticles.Rebuild(series.Id, series.Articles)
			if err == nil {
				response.Status = 200
				response.Data = seriesId
			}
		}
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func getUserSeries(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Status: 500, Data: "some error"}
	authorId := r.FormValue("author_id")
	limit := r.FormValue("limit")
	offset := r.FormValue("offset")

	aId, _ := strconv.ParseInt(authorId, 10, 32)
	perPage, _ := strconv.ParseInt(limit, 10, 32)
	skip, _ := strconv.ParseInt(offset, 10, 32)

	series := models.Series {AuthorId: int32(aId)}
	rows, err := series.Read(perPage, skip)

	if err != nil {
		response.Status = 500
		response.Data = err
	} else {
		response.Status = 200
		response.Data = rows
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func logout(w http.ResponseWriter, r *http.Request) {
	uid, err := r.Cookie(cookieNameId)

	if err == nil {
		cookie := http.Cookie{
			Name:    cookieNameId,
			Value:   uid.Value,
			Expires: time.Now().Add(-100 * time.Hour),
		}
		http.SetCookie(w, &cookie)
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

func isLoggedAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	uid, err := r.Cookie(cookieNameId)

	if err == nil {
		_, err2 := getLoggedUser(uid.Value)

		if err2 != nil {
			response.Status = 500
			response.Data = "Wtf"
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

func confirmPhoneAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Status: 404, Data: "not to confirm"}
	confirmCode := r.FormValue("code")
	uid, err := r.Cookie(cookieNameId)

	if err == nil && len(confirmCode) > 0 {
		user, _ := getLoggedUser(uid.Value)
		confirmed, err := user.ConfirmPhone(confirmCode)
		if err != nil {
			response.Status = 500
			response.Data = err
		} else if !confirmed {
			response.Status = 403
			response.Data = "bad code"
		} else {
			response.Status = 200
			response.Data = confirmed
		}
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func getTags(w http.ResponseWriter, r *http.Request) {
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

func getArticleAction(w http.ResponseWriter, r *http.Request) {
	user := models.User{}
	response := models.Response{
		Status: 404,
		Data:   "Article not found",
	}

	id := r.FormValue("id")
	articleId, _ := strconv.ParseInt(id, 10, 64)

	article := models.Article{}
	err := article.One(articleId)

	if err != nil {
		response.Status = 500
		response.Data = err
	} else {
		uid, err := r.Cookie(cookieNameId)
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

func removeArticleAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Status: 401, Data: "Unauthorized"}
	id := r.FormValue("id")
	uid, _ := r.Cookie(cookieNameId)

	if uid != nil && id != "" {
		articleId, _ := strconv.ParseInt(id, 10, 32)
		article := models.Article{Id: int32(articleId)}
		if user, _ := getLoggedUser(uid.Value); user.Id > 0 {
			article.AuthorId = user.Id
			id, err := article.Remove()

			if err != nil {
				response.Status = 500
				response.Data = err
			} else if id == 0{
				response.Status = 404
				response.Data = "impossible to delete this article"
			} else {
				response.Status = 200
				response.Data = id
			}
		}
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func addArticleAction(w http.ResponseWriter, r *http.Request) {
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

func updateArticleAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Status: 500, Data: "Something went wrong"}
	article := models.Article{}

	uid, err := r.Cookie(cookieNameId)

	if uid != nil && err == nil {
		err := json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			response.Status = 500
			response.Data = err.Error()
		} else if user, err := getLoggedUser(uid.Value); err == nil && user.Id == article.AuthorId {
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

func getPublishedArticles(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}

	uid, err := r.Cookie(cookieNameId)
	user, _ := getLoggedUser(uid.Value)

	if err != nil {
		response.Status = 403
		response.Data = "Require auth"
	} else {
		if user.Id > 0 {
			limit := r.FormValue("limit")
			offset := r.FormValue("offset")
			perPage, _ := strconv.ParseInt(limit, 10, 64)
			skip, _ := strconv.ParseInt(offset, 10, 64)

			article := models.Article{}
			articles, err := article.GetPublished(int64(user.Id), perPage, skip)

			if err != nil {
				response.Status = 500
				response.Data = err
			} else {
				response.Status = 200
				response.Data = articles
			}
		}
	}

	w.Header().Set("Content-type", contentTypeJson)
	w.Write(response.ToBytes())
}

func getArticlesAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	authorId := r.FormValue("author_id")
	tag := r.FormValue("tag")
	limit := r.FormValue("limit")
	offset := r.FormValue("offset")
	showPublished := r.FormValue("show_published")

	aId, _ := strconv.ParseInt(authorId, 10, 64)
	perPage, _ := strconv.ParseInt(limit, 10, 64)
	skip, _ := strconv.ParseInt(offset, 10, 64)
	sPublished, _ := strconv.ParseInt(showPublished, 10, 64)

	article := models.Article{}
	articles, err := article.Get(sPublished, aId, perPage, skip, tag)

	if err != nil {
		response.Status = 500
		response.Data = err
	} else {
		uid, err := r.Cookie(cookieNameId)

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

func checkPhoneNumberAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	phone := helpers.ThoroughlyClearString(r.FormValue("phone"))
	user := models.User{Phone: phone}

	if exists, err := user.PhoneNumberExists(); err == nil {
		response.Status = 200
		response.Data = exists
	} else {
		response.Status = 500
		response.Data = err
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func checkNickNameAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	nickName := r.FormValue("nickname")
	user := models.User{NickName: nickName}

	if exists, err := user.NickNameExists(); err == nil {
		response.Status = 200
		response.Data = exists
	} else {
		response.Status = 500
		response.Data = err
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func addUserAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}

	user := models.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		response.Status = 500
		response.Data = err.Error()
	} else {
		uid := uuid.Must(uuid.NewV4())

		newUser := models.User{
			Person:   user.Person,
			NickName: user.NickName,
			Avatar:   user.Avatar,
			Uuid:     uid.String(),
			Country:  user.Country,
			Phone:    user.Phone,
		}

		if errAdd := newUser.Add(); errAdd != nil {
			response.Status = 500
			response.Data = newUser
		} else if newUser.Id == 0 {
			response.Status = 403
			response.Data = "authentication data is already used"
		} else {
			cookie := http.Cookie{
				Name:    cookieNameId,
				Value:   newUser.Uuid,
				Expires: time.Now().Add(360 * 24 * time.Hour),
			}

			http.SetCookie(w, &cookie)
			loggedUsers[newUser.Uuid] = newUser

			response.Status = 200
			response.Data = newUser
		}
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func signInAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{}
	code := helpers.ThoroughlyClearString(r.FormValue("code"))
	phone := helpers.ThoroughlyClearString(r.FormValue("phone"))

	attemptLogin := models.AttemptLogin{Code: code, Phone: phone}

	if id, err := attemptLogin.Last(); err == nil && id > 0 {
		user := new(models.User)
		user.OneByPhone(phone)

		if user.Id > 0 {
			loggedUsers[user.Uuid] = *user

			cookie := http.Cookie{
				Name:    cookieNameId,
				Value:   user.Uuid,
				Expires: time.Now().Add(360 * 24 * time.Hour),
			}
			http.SetCookie(w, &cookie)

			data, _ := json.Marshal(user)

			response.Status = 200
			response.Data = data
		} else {
			response.Status = 404
			response.Data = "user not found"
		}
	} else {
		response.Status = 403
		response.Data = "bad code"
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func getPersonAction(w http.ResponseWriter, r *http.Request) {
	response := models.Response{Status: 404, Data: false}
	id := r.FormValue("id")
	personId, err := strconv.ParseInt(id, 10, 64)

	var user models.User
	uuidCookie, errCookie := r.Cookie(cookieNameId)

	if err == nil && personId > 0 {
		user = models.User{}
		if err := user.One(personId); err != nil {
			response.Status = 500
			response.Data = err
		} else if user.Id == 0 {
			response.Status = 404
			response.Data = "User not found"
		} else {
			if errCookie == nil && uuidCookie.Value == user.Uuid {
				user.IsOwner = true
			} else {
				user.Uuid = ""
				user.Phone = ""
			}
			response.Status = 200
			response.Data = user
		}
	} else if errCookie == nil {
		user = models.User{Uuid: uuidCookie.Value}
		user.IsOwner = true

		if exists, err := user.Exists(); err == nil && exists {
			loggedUsers[user.Uuid] = user
			response.Status = 200
			response.Data = user
		}
	}

	w.Header().Set("Content-Type", contentTypeJson)
	w.Write(response.ToBytes())
}

func getPersonsAction(w http.ResponseWriter, r *http.Request) {
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

func getCodeToLoginAction(w http.ResponseWriter, r *http.Request) {
	ip := helper.GetIp()
	response := new(models.Response)
	phone := helpers.ThoroughlyClearString(r.FormValue("phone"))

	user := models.User{Phone:phone}

	if len(phone) < 10 {
		response.Status = 400
		response.Data = "invalid phone"
	} else if exists, err := user.PhoneNumberExists(); err != nil || !exists {
		response.Status = 404
		response.Data = "this phone does not exist"
	} else {
		attemptCode := models.AttemptLogin{
			Phone: phone,
			Code: helper.GetConformationCode(),
			Ip: ip,
		}

		err := attemptCode.SendCode()
		if err != nil {
			response.Status = 500
			response.Data = err.Error()
		} else {
			response.Status = 200
			response.Data = "code is sended"
		}
	}

	w.Header().Set("Content-type", contentTypeJson)
	w.Write(response.ToBytes())
}

