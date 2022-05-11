package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/drshapeless/slscoreboard/go-scoreboard/internal/data"
	"github.com/go-chi/chi"

	_ "github.com/mattn/go-sqlite3"
)

type config struct {
	port int
	path string
}

type application struct {
	config   config
	infoLog  *log.Logger
	errorLog *log.Logger
	models   data.Models
}

func main() {
	var cfg config
	flag.StringVar(&cfg.path, "path", os.Getenv("HOME")+"/scoreboard.db", "scoreboard database")
	flag.IntVar(&cfg.port, "port", 8000, "scoreboard port")

	flag.Parse()

	var app application
	app.config = cfg
	app.infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	app.errorLog = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	if !isExist(app.config.path) {
		app.errorLog.Fatalln("scoreboard database does not exist!")
		return
	} else {
		db, err := openDB(app.config.path)
		if err != nil {
			app.errorLog.Fatal(err)
		}
		app.models = data.NewModels(db)
	}

}

func openDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func isExist(path string) bool {
	_, err := os.Stat(path)
	return !errors.Is(err, os.ErrNotExist)
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", app.config.port),
		Handler:      app.routes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	app.infoLog.Printf("starting scoreboard server at port %d\n", app.config.port)
	err := srv.ListenAndServe()
	if err != nil {
		return err
	}
	app.infoLog.Printf("stopped scoreboard server\n")

	return nil
}

func (app *application) routes() http.Handler {
	r := chi.NewRouter()

	apiPrefix := "/api/v1"

	r.Route(apiPrefix, func(r chi.Router) {
		r.Post("/snooker", app.createSnookerHandler)
		r.Get("/snooker/{page}", app.listSnookerHandler)
	})

	return r
}

func (app *application) listSnookerHandler(w http.ResponseWriter, r *http.Request) {
	page, err := app.readPageParam(r)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	snookers, maxPage, err := app.models.Snookers.GetAll(int(page))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"snookers": snookers, "max_page": maxPage}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createSnookerHandler(w http.ResponseWriter, r *http.Request) {

}
