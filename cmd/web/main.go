package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Saksham0239/bookings/pkg/config"
	"github.com/Saksham0239/bookings/pkg/handlers"
	"github.com/Saksham0239/bookings/pkg/render"

	"github.com/alexedwards/scs/v2"
)

const PORT = ":8000"

var app config.AppConfig

var session *scs.SessionManager

func main() {

	//set to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	app.Session = session

	tc, err := render.CreateTempateCache()

	if err != nil {
		log.Fatal("Cannot create template cache")
	}
	//loading the template cache
	app.TemplateCache = tc
	app.UseCache = false

	render.NewTemplates(&app)

	//adding repo to handlers
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	fmt.Printf("Server started on PORT %s \n", PORT)

	srv := &http.Server{
		Addr:    PORT,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
