package main

import (
	"fmt"
	"github.com/marioteik/helpers"
	"github.com/marioteik/website/config"
	"html/template"
	"log"
	"net/http"
	"time"
)

const version = "0.0.0"
const author = "Mario Teik"

type application struct {
	env           *config.Env
	info          *log.Logger
	error         *log.Logger
	version       string
	author        string
	templateCache map[string]*template.Template
	//db       	*gorm.DB
}

func main() {
	logger, logError := config.InitLoggers()

	env, err := config.GetEnvironment()
	if err != nil {
		logError.Fatal(err)
	}

	tc, err := helpers.CreateTemplateCache()
	if err != nil {
		logError.Fatal(err)
	}

	app := &application{
		env:           env,
		info:          logger,
		error:         logError,
		version:       version,
		author:        author,
		templateCache: tc,
		/*db:       config.DB*/
	}

	err = app.serve()
	if err != nil {
		app.error.Println("")
	}
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", app.env.Port),
		//Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	app.info.Printf("Starting application server on %s", srv.Addr)

	return srv.ListenAndServe()
}
