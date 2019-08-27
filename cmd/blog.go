package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/shinichi2510/go-blog-service/internal/http/handler"
	"github.com/shinichi2510/go-blog-service/internal/repo"
	"github.com/shinichi2510/go-blog-service/internal/service"
	"github.com/spf13/viper"
)

func main() {
	viper.SetEnvPrefix("BLOG")
	viper.AutomaticEnv()
	dbURL := viper.GetString("DB_URL")

	log.Info().Str("DB_URL", dbURL).Msg("DB_URL")

	logLevelStr := viper.GetString("LOG_LEVEL")
	logLevel, err := zerolog.ParseLevel(logLevelStr)
	if err != nil {
		log.Fatal().Err(err).Msg("Invalid LOG_LEVEL")
	}
	zerolog.LevelFieldName = "severity"
	zerolog.SetGlobalLevel(logLevel)
	log.Logger = log.With().Caller().Logger()

	db, err := sqlx.Connect("mysql", dbURL)

	if err != nil {
		log.Fatal().Str("dbURL", dbURL).Err(err).Msg("Cannot connect to DB:")
	}

	defer func() {
		if err = db.Close(); err != nil {
			log.Error().Err(err).Msg("Cannot close DB connection")
		}
	}()

	articleRepo := repo.NewArticle(db)
	createArticleService := service.NewCreateArticleService(articleRepo)
	getArticleService := service.NewGetArticleService(articleRepo)
	deleteArticleService := service.NewDeleteArticleService(articleRepo)
	updateArticleService := service.NewUpdateArticleService(articleRepo)
	//listArticleService := service.NewUpdateArticleService(articleRepo)

	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Token", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler)

	r.Route("/articles", func(r chi.Router) {
		//r.Method("GET", "/", handler.NewListArticleHandler(listArticleService))
		r.Method("GET", "/{articleID}", handler.NewShowArticleHandler(getArticleService))
		r.Method("POST", "/", handler.NewCreateArticleHandler(createArticleService))
		r.Method("PUT", "/{articleID}", handler.NewUpdateArticleHandler(updateArticleService))
		r.Method("DELETE", "/{articleID}", handler.NewDeleteArticleHandler(deleteArticleService))
	})

	err = http.ListenAndServe(":8081", r)
	log.Error().Err(err).Msg("Stopped serving HTTP")
}
