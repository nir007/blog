package main

import (
	"github.com/gorilla/mux"
	"github.com/nir007/blog/internal/handlers"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"net/http"
)

func initRoutesAndStartServ(
	logger *zap.SugaredLogger,
	articleHandler *handlers.ArticleHandler,
	authHandler *handlers.AuthHandler,
	seriesHandler *handlers.SeriesHandler,
	tagHandler *handlers.TagHandler,
	userHandler *handlers.UserHandler,
) {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public/dist/static"))
	rp := router.PathPrefix("/static/")
	rp.Handler(http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.Index)

	//auth
	router.HandleFunc("/logout", authHandler.Logout)
	router.HandleFunc("/aj_sign_in", authHandler.SignIn)
	router.HandleFunc("/aj_get_code_to_login", authHandler.GetCodeToLogin)
	router.HandleFunc("/aj_get_check_nickname", authHandler.CheckNickName)
	router.HandleFunc("/aj_is_logged", authHandler.IsLogged)
	router.HandleFunc("/aj_confirm_phone", authHandler.ConfirmPhone)
	router.HandleFunc("/aj_get_check_phone", authHandler.CheckPhoneNumber)

	// tags
	router.HandleFunc("/aj_get_tags", tagHandler.GetTags)

	// series
	router.HandleFunc("/create_series", seriesHandler.CreateSeries)
	router.HandleFunc("/get_one_series", seriesHandler.GetOneSeries)
	router.HandleFunc("/delete_series", seriesHandler.DeleteSeries)
	router.HandleFunc("/update_series", seriesHandler.UpdateSeries)
	router.HandleFunc("/get_user_series", seriesHandler.GetUserSeries)

	//articles
	router.HandleFunc("/get_articles", articleHandler.GetArticles)
	router.HandleFunc("/get_published_articles", articleHandler.GetPublishedArticles)
	router.HandleFunc("/aj_get_article", articleHandler.GetArticle)
	router.HandleFunc("/remove_article", articleHandler.RemoveArticle)
	router.HandleFunc("/aj_add_article", articleHandler.AddArticle)
	router.HandleFunc("/aj_update_article", articleHandler.UpdateArticle)

	//users
	router.HandleFunc("/aj_add_user", userHandler.AddUser)
	router.HandleFunc("/aj_get_person", userHandler.GetPerson)
	router.HandleFunc("/aj_get_persons", userHandler.GetPersons)

	logger.Info("http server started on", viper.GetString("port"))
	log.Fatal(http.ListenAndServe(viper.GetString("port"), router))
	//http.ListenAndServeTLS(":443", "rakan-tarakan.com.crt", "private.key", router)
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req,
		"https://"+req.Host+req.URL.String(),
		http.StatusMovedPermanently)
}
