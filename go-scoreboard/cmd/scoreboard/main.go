package main

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/drshapeless/slscoreboard/go-scoreboard/internal/data"
	"github.com/go-chi/chi"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Pass = "yot"
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

	app.serve()
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

	apiPrefix := "/v1"

	r.NotFound(app.notFoundResponse)
	r.MethodNotAllowed(app.methodNotAllowedResponse)

	r.Route(apiPrefix, func(r chi.Router) {
		r.Post("/snooker/", app.createSnookerHandler)
		r.Get("/snooker/{page}", app.listSnookerHandler)

		r.Post("/dee/", app.createDeeHandler)
		r.Get("/dee/{page}", app.listDeeHandler)

		r.Post("/landlord/", app.createLandlordHandler)
		r.Get("/landlord/{page}", app.listLandlordHandler)
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
	var input struct {
		Pass   string `json:"pass"`
		Winner string `json:"winner"`
		Loser  string `json:"loser"`
		Diff   int    `json:"diff"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Pass != Pass {
		app.invalidCredentialsResponse(w, r)
		return
	}

	snooker := &data.Snooker{
		Winner: input.Winner,
		Loser:  input.Loser,
		Diff:   input.Diff,
	}

	err = app.models.Snookers.Insert(snooker)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"snooker": snooker}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listDeeHandler(w http.ResponseWriter, r *http.Request) {
	page, err := app.readPageParam(r)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	dees, maxPage, err := app.models.Snookers.GetAll(int(page))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"dees": dees, "max_page": maxPage}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createDeeHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Pass       string `json:"pass"`
		Winner     string `json:"winner"`
		Loser1     string `json:"loser1"`
		Loser1Card int    `json:"loser1_card"`
		Loser2     string `json:"loser2"`
		Loser2Card int    `json:"loser2_card"`
		Loser3     string `json:"loser3"`
		Loser3Card int    `json:"loser3_card`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Pass != Pass {
		app.invalidCredentialsResponse(w, r)
		return
	}

	dee := &data.Dee{
		Winner:     input.Winner,
		Loser1:     input.Loser1,
		Loser1Card: input.Loser1Card,
		Loser2:     input.Loser2,
		Loser2Card: input.Loser2Card,
		Loser3:     input.Loser3,
		Loser3Card: input.Loser3Card,
	}

	err = app.models.Dees.Insert(dee)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"dee": dee}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) listLandlordHandler(w http.ResponseWriter, r *http.Request) {
	page, err := app.readPageParam(r)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	landlords, maxPage, err := app.models.Landlords.GetAll(int(page))
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"landlords": landlords, "max_page": maxPage}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) createLandlordHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Pass     string `json:"pass"`
		Landlord string `json:"landlord"`
		Farmer1  string `json:"farmer1"`
		Farmer2  string `json:"farmer2"`
		Win      bool   `json:"win"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	if input.Pass != Pass {
		app.invalidCredentialsResponse(w, r)
		return
	}

	var winInt = 0
	if input.Win {
		winInt = 1
	}

	landlord := &data.Landlord{
		Landlord: input.Landlord,
		Farmer1:  input.Farmer1,
		Farmer2:  input.Farmer2,
		Win:      winInt,
	}

	err = app.models.Landlords.Insert(landlord)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"landlord": landlord}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
