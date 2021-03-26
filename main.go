package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"time"

	"examenMutante/config"
	mErrors "examenMutante/errors"
	"examenMutante/store"
	"examenMutante/transport"
	"examenMutante/usecase"
	_ "github.com/lib/pq"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
)

type Config struct {
	PGHost            string
	PGPort            string
	PGUsername        string
	PGPassword        string
	DatabaseName      string
	SSLMode           string
	MaxConnectionTime int
}

func main() {
	ctx := context.Background()
	logger := logrus.WithContext(ctx)

	dbInfo := os.Getenv("DATABASE_URL")
	var cfg Config
	if dbInfo == "" {
		err := config.LoadConfig(&cfg)
		if err != nil {
			logger.WithError(err).
				Fatal("cannot load configuration")
		}
		pgURL := url.URL{
			Scheme: "postgres",
			Host: fmt.Sprintf("%s:%s", cfg.PGHost, cfg.PGPort),
			User: url.UserPassword(cfg.PGUsername, cfg.PGPassword),
			Path: cfg.DatabaseName,
		}
		options := url.Values{}
		options.Set("sslmode", cfg.SSLMode)
		pgURL.RawQuery = options.Encode()
		dbInfo = pgURL.String()
	}

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		logger.Fatal("Could not start postgres database: %v ", err)
	}
	defer db.Close()

	mutanstStore, err := store.NewStore(db, cfg.MaxConnectionTime)
	if err != nil {
		logrus.WithError(err).Fatalf("couldn't create database")
	}
	mutantsUseCase := usecase.NewMutants(mutanstStore)

	mutantsTransport := transport.NewMutant(mutantsUseCase)

	e := transport.NewRouter(mutantsTransport)
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.HTTPErrorHandler = mErrors.EchoErrorHandler()

	migrator, err := migrate.New(
		"file://migrations/", dbInfo,
		)
	if err != nil {
		logger.WithError(err).Fatalf("cannot open migrations folder")
	}

	go func() {
		var targetSchema uint = 1
		currentVersion, dirty, err := migrator.Version()
		if err != nil {
			if err != migrate.ErrNilVersion {
				logger.WithError(err).Fatalf("Failed to get the DB schema version")
			}
		}
		if dirty {
			logger.Fatalf("DB SchemaMigrations table is dirty. Manual intervention is required." +
				"After revision please set SchemaMigrations.Dirty to false")
		}
		steps := int(targetSchema) - int(currentVersion)
		if steps > 0 {
			logger.Infof("Migrating DB schema up to version '%v' from version '%v'", targetSchema, currentVersion)
			if err := migrator.Steps(steps); err != nil {
				logger.WithError(err).Fatalf("Failed to migrate up to DB schema version '%v'", targetSchema)
			}
		} else if steps <= 0 {
			logger.Infof("The current DB schema version is '%v'. "+
				"Assuming compatibility with target version '%v'.", currentVersion, targetSchema)
		}
		logger.Infof("Service is ready to handle requests.")
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	go func() {
		if err := e.Start(":" + port); err != nil {
			e.Logger.Info("shutting down server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
