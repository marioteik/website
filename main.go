package main

import (
	"fmt"
	"github.com/marioteik/helpers"
	"github.com/marioteik/website/config"
	"net/http"
	"time"
)

const version = "0.0.0"
const author = "Mario Teik"

type application struct {
	env           *config.Env
	logger        *config.Logger
	version       string
	author        string
	templateCache *helpers.TemplateHelper
	//db       	*gorm.DB
}

func main() {
	logger := config.InitLoggers()

	env, err := config.GetEnvironment()
	if err != nil {
		logger.Fatal(err)
	}

	componentPath := "./components"
	templateHelper := helpers.NewTemplateHelper("./pages", "./components/templates", &componentPath)

	if env.InProduction {
		err = templateHelper.CreateTemplateCache()
		if err != nil {
			logger.Fatal(err)
		}
	} else {
		templateHelper.UseCache = false
	}

	app := &application{
		env:           env,
		logger:        logger,
		version:       version,
		author:        author,
		templateCache: templateHelper,
		/*db:       config.DB*/
	}

	err = app.serve()
	if err != nil {
		app.logger.Fatal(err)
	}
}

func (app *application) serve() error {
	srv := &http.Server{
		Addr:              fmt.Sprintf(":%s", app.env.Port),
		Handler:           app.routes(),
		IdleTimeout:       30 * time.Second,
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      5 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
	}

	app.logger.Info(fmt.Sprintf("Starting application server on %s", srv.Addr))

	return srv.ListenAndServe()
}
