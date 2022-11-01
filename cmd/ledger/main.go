package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/KekemonBS/ledgerTest/env"
	"github.com/KekemonBS/ledgerTest/handlers"
	"github.com/KekemonBS/ledgerTest/router"
	"github.com/KekemonBS/ledgerTest/storage/postgresql"

	"github.com/golang-migrate/migrate/v4"
	pgxm "github.com/golang-migrate/migrate/v4/database/pgx"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
)

func main() {
	cfg, err := env.NewConfig()
	if err != nil {
		log.Fatal(err)
	}
	logger := log.New(os.Stdout, "log:", log.Lshortfile)

	ctx, cancel := context.WithCancel(context.Background())
	//-------------------------------------------------------------------
	cfgpgx, err := pgx.ParseConfig(cfg.PostgresURI)
	if err != nil {
		logger.Fatal(err)
	}
	db, err := pgx.ConnectConfig(ctx, cfgpgx)
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close(ctx)

	logger.Println("Postgres URI : ", cfg.PostgresURI)
	err = db.Ping(ctx)
	if err != nil {
		logger.Fatal(err)
	}

	tmpconn := stdlib.OpenDB(*cfgpgx)
	driver, err := pgxm.WithInstance(tmpconn, &pgxm.Config{})
	if err != nil {
		logger.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://./storage/migrations",
		"postgres", driver)
	if err != nil {
		logger.Fatal(err)
	}
	_ = m.Up()
	tmpconn.Close()
	//-------------------------------------------------------------------
	dbImpl := postgresql.New(db)
	if err != nil {
		logger.Fatal(err)
	}

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		logger.Printf("got signal: %v", <-ch)
		cancel()
	}()

	//Init handlers
	h := handlers.New(ctx, logger, dbImpl)
	//Init router
	r := router.New(h)
	//Listen and serve
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	go func() {
		err = s.ListenAndServe()
		if err != nil {
			logger.Fatal(err)
		}
	}()
	for {
		select {
		case <-ctx.Done():
			s.Close()
			return
		}
	}

}
