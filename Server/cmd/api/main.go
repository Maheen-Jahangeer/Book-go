package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type Appstatus struct {
	Status      string `json:"status"`
	Environment string `json:"environment"`
}

type application struct {
	config config
	logger *log.Logger
}

func main() {
	var conf config

	flag.IntVar(&conf.port, "port", 4000, "Server port listen on")
	flag.StringVar(&conf.env, "env", "development", "Application environment (devlepment | production)")
	flag.StringVar(&conf.db.dsn, "dsn", "postgres://mahi:qburst@localhost/books?sslmode=disable", "postgress connection string")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := openDB(conf)
	if err != nil {
		logger.Print("error :", err)
	}
	defer db.Close()

	app := &application{
		config: conf,
		logger: logger,
	}

	serv := &http.Server{
		Addr:         fmt.Sprintf(":%d", conf.port),
		Handler:      app.router(),
		ReadTimeout:  10 * time.Second,
		IdleTimeout:  time.Minute,
		WriteTimeout: 10 * time.Second,
	}

	logger.Println("Server is running on port", conf.port)

	err = serv.ListenAndServe()

	if err != nil {
		fmt.Println("Error", err)
	}
}

func openDB(cfg config) (*sql.DB, error) {
	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		return nil, err
	}

	ctxt, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctxt)
	if err != nil {
		return nil, err
	}

	return db, nil
}
