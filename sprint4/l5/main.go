package main

import (
	"log"
	//"net/http"

	// "github.com/ArtemRivs/shortlinker/internal/handlers/middleware"
	"github.com/ArtemRivs/ypgoadv/sprint4/l5/internal/storage"
)

const (
	DatabaseDsn = ""
)

func main() {

	ls := storage.New(DatabaseDsn)
	defer ls.Close()
	// r := chi.NewRouter()
	// r.Use(middleware.GzipHandler)
	// r.Use(middleware.AuthHandler)

	// r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.GetFullLink(w, r, ls)
	// })
	// r.Post("/", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.GetShortLink(w, r, ls, cfg.BaseURL)
	// })
	// r.Post("/api/shorten", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.GetShortLinkJSON(w, r, ls, cfg.BaseURL)
	// })
	// r.Get("/api/user/urls", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.GetUserLinks(w, r, ls, cfg.BaseURL)
	// })
	// r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.DBPing(w, r, cfg.DatabaseDsn)
	// })
	// r.Post("/api/shorten/batch", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.GetShortLinksBatchJSON(w, r, ls, cfg.BaseURL)
	// })
	// r.Delete("/api/user/urls", func(w http.ResponseWriter, r *http.Request) {
	// 	handlers.DeleteUserLinksBatch(w, r, ls)
	// })

	log.Println("started")
	// log.Fatal(http.ListenAndServe(ServerAddr, r))
}
