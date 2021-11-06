package main

import (
	"github.com/nir007/blog/internal/handlers"
	"github.com/nir007/blog/internal/services"
	"github.com/spf13/viper"
)

func main() {
	logger := getLogger()

	if err := initConfig(); err != nil {
		logger.Error(err)
		return
	}

	db := getDatabase()

	//services
	postgresService := services.NewPGDatabaseFucker(
		db,
		viper.GetString("db.schema"),
	)

	iqSmsSrv := services.NewIQSms(
		viper.GetString("sendPulse.login"),
		viper.GetString("sendPulse.password"),
		viper.GetString("sendPulse.baseUrl"),
		viper.GetString("sendPulse.sendUrl"),
		viper.GetString("sendPulse.sender"),
	)

	attemptConfirmSrv := services.NewAttemptConfirmService(postgresService)
	userSrv := services.NewUserService(postgresService, iqSmsSrv, attemptConfirmSrv)

	articleSrv := services.NewArticleService(postgresService)
	attemptLoginSrv := services.NewAttemptLoginService(iqSmsSrv, postgresService)
	seriesSrv := services.NewSeriesService(postgresService)
	seriesArticleSrv := services.NewSeriesArticleService(postgresService)

	//handlers
	articleHandler := handlers.NewArticleHandler(articleSrv)
	authHandler := handlers.NewAuthHandler(userSrv, attemptLoginSrv)
	seriesHandler := handlers.NewSeriesHandler(seriesSrv, userSrv, seriesArticleSrv)
	tagHandler := handlers.NewTagHandler(articleSrv)
	userHandler := handlers.NewUserHandler(userSrv)

	initRoutesAndStartServ(
		logger,
		articleHandler,
		authHandler,
		seriesHandler,
		tagHandler,
		userHandler,
	)
}