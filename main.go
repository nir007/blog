package main

import (
	"blog/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./public/dist/static"))
	rp := router.PathPrefix("/static/")
	rp.Handler(http.StripPrefix("/static/", fs))

	router.HandleFunc("/", handlers.IndexAction)
	router.HandleFunc("/logout", handlers.Logout)
	router.HandleFunc("/aj_add_user", handlers.AddUserAction)
	router.HandleFunc("/aj_get_tags", handlers.GetTags)
	router.HandleFunc("/aj_sign_in", handlers.SignInAction)
	router.HandleFunc("/aj_get_check_nickname", handlers.CheckNickNameAction)
	router.HandleFunc("/aj_is_logged", handlers.IsLoggedAction)
	router.HandleFunc("/aj_confirm_phone", handlers.ConfirmPhoneAction)
	router.HandleFunc("/aj_get_check_phone", handlers.CheckPhoneNumberAction)
	router.HandleFunc("/create_series", handlers.CreateSeriesAction)
	router.HandleFunc("/get_one_series", handlers.GetOneSeriesAction)
	router.HandleFunc("/delete_series", handlers.DeleteSeriesAction)
	router.HandleFunc("/update_series", handlers.UpdateSeriesAction)
	router.HandleFunc("/get_user_series", handlers.GetUserSeries)
	router.HandleFunc("/get_articles", handlers.GetArticlesAction)
	router.HandleFunc("/get_published_articles", handlers.GetPublishedArticles)
	router.HandleFunc("/aj_get_article", handlers.GetArticleAction)
	router.HandleFunc("/remove_article", handlers.RemoveArticleAction)
	router.HandleFunc("/aj_add_article", handlers.AddArticleAction)
	router.HandleFunc("/aj_update_article", handlers.UpdateArticleAction)
	router.HandleFunc("/aj_get_person", handlers.GetPersonAction)
	router.HandleFunc("/aj_get_persons", handlers.GetPersonsAction)
	router.HandleFunc("/aj_get_code_to_login", handlers.GetCodeToLoginAction)

	log.Fatal(http.ListenAndServe(":80", router))
	//http.ListenAndServeTLS(":443", "rakan-tarakan.com.crt", "private.key", router)
}

func redirect(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req,
		"https://"+req.Host+req.URL.String(),
		http.StatusMovedPermanently)
}
