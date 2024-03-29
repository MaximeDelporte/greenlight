package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

/*
Hold the dependencies for our HTTP handlers, helpers and middleware.
*/
type application struct {
	config config
	logger *log.Logger
}

func main() {
	var cfg config

	/*
		Read the value of the port and anv command-line flags into the config struct.
		We default to using the port 4000 and the environment "development" if no corresponding flags are provided.
	*/
	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "development", "Environment (development|staging|production)")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	app := &application{
		config: cfg,
		logger: logger,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.port),
		Handler:      handler(app.routes()),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	logger.Printf("starting %s server on %s", cfg.env, srv.Addr)
	err := srv.ListenAndServe()
	logger.Fatal(err)
}

func handler(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w = &Interceptor{writer: w, request: r}
		h.ServeHTTP(w, r)
	}

	return http.HandlerFunc(fn)
}
