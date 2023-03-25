package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	//"net/http"
	// "github.com/ArtemRivs/shortlinker/internal/handlers/middleware"
)

const (
	DatabaseDsn = "firstbase"
)

func main() {

	log.Println("started")

	db, err := sql.Open("postgres", DatabaseDsn)
	if err != nil {
		log.Println("Connection error:", err)
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Println("Ping error:", err)
		panic(err)
	}
	// _, err = db.Exec("CREATE TABLE IF NOT EXISTS links (short_code VARCHAR NOT NULL, origin_url TEXT UNIQUE, user_id VARCHAR, is_deleted BOOLEAN DEFAULT FALSE);")

	// if err != nil {
	// 	panic(err)
	// }

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

	// log.Println("started")
	// log.Fatal(http.ListenAndServe(ServerAddr, r))
}
