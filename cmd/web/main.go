package main

import (
	// "errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/bjorndonald/bookings/internal/config"
	"github.com/bjorndonald/bookings/internal/handlers"
	"github.com/bjorndonald/bookings/internal/render"
)

const portNumber = ":8080"

// func Divide(w http.ResponseWriter, r *http.Request) {
// 	var x, y float32
// 	x = 100.0
// 	y = 0.0
// 	f, err := divideValues(x, y)
// 	if err != nil {
// 		fmt.Fprintln(w, "Cannot divide by 0")
// 		return
// 	}
// 	_, _ = fmt.Fprintf(w, fmt.Sprintf("%f divided by %f is %f", x, y, f))
// }

// func addValues(x, y int) int {
// 	return x + y
// }

// func divideValues(x, y float32) (float32, error) {
// 	if y == 0 {
// 		err := errors.New("cannot divide by zero")
// 		return 0, err
// 	}

//		return x / y, nil
//	}
var app config.AppConfig
var session *scs.SessionManager

func main() {
	// get the template cache from the app config

	// change to true when in production
	app.InProduction = false

	session = scs.New()
	session.Lifetime = 24 * time.Hour
	session.Cookie.Persist = true
	session.Cookie.SameSite = http.SameSiteLaxMode
	session.Cookie.Secure = app.InProduction

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("cannopt create Template cache")
	}

	app.TemplateCache = tc
	app.UseCache = false
	app.Session = session
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	render.NewTemplates(&app)

	fmt.Printf(fmt.Sprintf("Starting application on port %s", portNumber))
	srv := &http.Server{
		Addr:    portNumber,
		Handler: routes(&app),
	}

	err = srv.ListenAndServe()
	log.Fatal(err)
}
