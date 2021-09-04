package app

import (
	"home-24/handler"
	"home-24/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Application interface {
	Serve(addr string)
}

type app struct {
	crawlerHandler handler.CrawlerHandler
}

func Init() Application {
	client := utils.NewHttpClient()
	return &app{
		crawlerHandler: handler.NewCrawlerhandler(client),
	}
}

func (instance *app) Serve(addr string) {
	router := chi.NewRouter()
	router.Use(middleware.RequestID, middleware.Logger, middleware.Recoverer, middleware.Heartbeat("/health"))
	router.Route("/", func(r chi.Router) {
		fs := http.FileServer(http.Dir("static"))
		router.Handle("/", fs)
		r.Post("/analyze", instance.crawlerHandler.Analyze)
	})
	http.ListenAndServe(addr, router)
}
