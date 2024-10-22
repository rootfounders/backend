package main

import (
	"context"
	"flag"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/rootfounders/rootfounders/config"
	"github.com/rootfounders/rootfounders/database"
)

var dbConn *database.Connection

func main() {
	configPath := flag.String("config", "./rootfounders.toml", "path to config file")
	flag.Parse()
	cnf, err := config.Load(*configPath)
	if err != nil {
		log.Fatalln(err)
	}
	if err := InitIpfs(cnf); err != nil {
		log.Fatalln("ipfs init error:", err)
	}
	ctx := context.Background()
	initDatabase(ctx, cnf)
	defer dbConn.Conn().Close(ctx)

	url := cnf.Website.ListenUrl
	if url == "" {
		log.Fatalln("listen url required")
	}
	log.Println("Listening on", url)

	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	InitRoutes(r)

	fs := http.FileServer(http.Dir("static"))
	r.Handle("/static/*", http.StripPrefix("/static/", fs))
	if err := http.ListenAndServe(url, r); err != nil {
		log.Fatalln(err)
	}
}

func initDatabase(ctx context.Context, cnf *config.Config) {
	dbConn = &database.Connection{
		Chain:    1, // TODO: make sure it's OK -- we only query db from here
		Host:     cnf.Database.Host,
		Port:     cnf.Database.Port,
		User:     cnf.Database.User,
		Password: cnf.Database.Password,
		Database: cnf.Database.Database,
	}
	if err := dbConn.Connect(ctx); err != nil {
		log.Fatalln("database connection error:", err)
	}
}
