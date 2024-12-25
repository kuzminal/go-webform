package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"github.com/kuzminal/go-webform/internal/routes"
	"net/http"

	_ "modernc.org/sqlite"

	"github.com/kuzminal/go-webform/pkg/commentdb"
)

func main() {
	db, err := sql.Open("sqlite", "./webform.db")
	if err != nil {
		panic(err)
	}
	commentdb.InitDB(db)

	r := mux.NewRouter()
	routes.Build(r, db)

	server := http.Server{
		Addr:    ":8181",
		Handler: r,
	}
	server.ListenAndServe()
}
