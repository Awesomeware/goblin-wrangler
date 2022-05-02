package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	"awesomeware.org/kingpin/internal/models"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *models.SnippetModel
}

type config struct {
	addr string
}

func main() {
	var cfg config

	flag.StringVar(&cfg.addr, "addr", "4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	databaseUrl := "postgres://postgres:postgres@localhost:5432"

	//dbpool, err := pgxpool.Connect(context.Background(), databaseUrl)
	dbpool, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		errorLog.Printf("Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer dbpool.Close()

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &models.SnippetModel{DB: dbpool},
	}

	srv := &http.Server{
		Addr:     ":" + cfg.addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", cfg.addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
