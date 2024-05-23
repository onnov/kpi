package main

import (
	"github.com/go-chi/chi/v5"
	m "github.com/go-chi/chi/v5/middleware"
	"kpi/internal/controller"
	"kpi/internal/dictionary"
	"kpi/internal/middleware"
	"kpi/internal/repository"
	"log"
	"net/http"
	"time"
)

func main() {
	// Функция сохраняет fact в БД из очереди(буфера)
	go func() {
		q := repository.GetQueue()
		for {
			err := q.SaveFact()
			if err != nil {
				log.Println("SaveFact: ", err)
				time.Sleep(time.Second)
			}
			time.Sleep(time.Millisecond)
		}
	}()

	// Роутер
	r := chi.NewRouter()
	r.Use(m.Logger)
	r.Use(m.Recoverer)

	// RESTy routes for "_api" resource
	r.Route("/_api", func(r chi.Router) {
		// use the Bearer Authentication middleware
		r.Use(middleware.SimpleBearerAuthorize)

		r.Post("/facts/save_fact", controller.FactsSaveFactHandler)           // POST /facts/save_fact
		r.Post("/indicators/get_facts", controller.IndicatorsGetFactsHandler) // POST /facts/save_fact
	})

	// Server start
	err := http.ListenAndServe(dictionary.Env.ListenedAddress, r)
	if err != nil {
		log.Fatal("http.ListenAndServe:", err)
	}
}
