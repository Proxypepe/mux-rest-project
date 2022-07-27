package app

import (
	"database/sql"
	"fmt"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"log"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize(user, password, dbname string) {
	conString := fmt.Sprintf("postgresql://%s:%s@localhost:5432/%s?sslmode=disable", user, password, dbname)
	var err error
	a.DB, err = sql.Open("postgres", conString)
	if err != nil {
		log.Fatal(err)
	}
	err = a.DB.Ping()
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()
}

func (a *App) Run(addr string) {

}
