package main

import (
	"Mutants/go-app/config"
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	mErrors "Mutants/go-app/errors"
	"Mutants/go-app/store"
	"Mutants/go-app/transport"
	"Mutants/go-app/usecase"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/labstack/echo/middleware"
	_ "github.com/lib/pq"
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

	var cfg Config
	var dbInfo *sql.DB
	var errInfo error
	var urlDB string
	if os.Getenv("DB_HOST") != "" {
		err := config.LoadConfig(&cfg)
		if err != nil {
			logger.WithError(err).
				Fatal("cannot load configuration")
		}
		pgURL := url.URL{
			Scheme: "postgres",
			Host:   fmt.Sprintf("%s:%s", cfg.PGHost, cfg.PGPort),
			User:   url.UserPassword(cfg.PGUsername, cfg.PGPassword),
			Path:   cfg.DatabaseName,
		}
		options := url.Values{}
		options.Set("sslmode", cfg.SSLMode)
		pgURL.RawQuery = options.Encode()
		urlDB = pgURL.String()
		dbInfo, errInfo = sql.Open("postgres", urlDB)
		if err != nil {
			logger.Fatal("sql.Open: %v", err)
		}
		defer dbInfo.Close()
	} else {
		dbInfo, urlDB, errInfo = initSocketConnectionPool()
		if errInfo != nil {
			log.Fatalf("initSocketConnectionPool: unable to connect: %s", errInfo)
		}
	}

	mutanstStore, err := store.NewStore(dbInfo)
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

	/*migrator, err := migrate.New(
		"file://migrations/", urlDB,
	)
	if err != nil {
		logger.WithError(err).Fatalf("cannot open migrations folder")
	}*/

	/*go func() {
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
	}()*/

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

func initSocketConnectionPool() (*sql.DB, string, error) {
	// [START cloud_sql_postgres_databasesql_create_socket]
	var (
		dbUser                 = "postgres"                 // e.g. 'my-db-user'
		dbPwd                  = "12345678"                  // e.g. 'my-db-password'
		instanceConnectionName = "mutants-310305:us-central1:mutants-dev" // e.g. 'project:region:instance'
		dbName                 = "marvel"                // e.g. 'my-database'
	)

	socketDir, isSet := os.LookupEnv("DB_SOCKET_DIR")
	if !isSet {
		socketDir = "/cloudsql"
	}

	var dbURI string
	dbURI = fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)

	// dbPool is the pool of database connections.
	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, "", fmt.Errorf("sql.Open: %v", err)
	}

	// [START_EXCLUDE]
	configureConnectionPool(dbPool)
	// [END_EXCLUDE]

	return dbPool, dbURI, nil
	// [END cloud_sql_postgres_databasesql_create_socket]
}

func configureConnectionPool(dbPool *sql.DB) {
	// [START cloud_sql_postgres_databasesql_limit]

	// Set maximum number of connections in idle connection pool.
	dbPool.SetMaxIdleConns(5)

	// Set maximum number of open connections to the database.
	dbPool.SetMaxOpenConns(7)

	// [END cloud_sql_postgres_databasesql_limit]

	// [START cloud_sql_postgres_databasesql_lifetime]

	// Set Maximum time (in seconds) that a connection can remain open.
	dbPool.SetConnMaxLifetime(1800)

	// [END cloud_sql_postgres_databasesql_lifetime]
}

func mustGetenv(k string) string {
	v := os.Getenv(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}