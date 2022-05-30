package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/spf13/viper"
)

func mustGetenv(k string) string {
	//v := os.Getenv(k)
	v := viper.GetString(k)
	fmt.Printf("env: %s=%s\n", k, v)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func initDb() (*sql.DB, error) {
	if viper.GetString("DB_SOCKET_DIR") != "" {
		return initSocketConnectionPool()
	} else {
		return initTCPConnectionPool()
	}
}

// initSocketConnectionPool initializes a Unix socket connection pool
func initSocketConnectionPool() (*sql.DB, error) {
	var (
		dbUser                 = mustGetenv("DB_USER")
		dbPwd                  = mustGetenv("DB_PASS")
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
		dbName                 = mustGetenv("DB_NAME")
		socketDir              = mustGetenv("DB_SOCKET_DIR")
	)

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)

	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	configureConnectionPool(dbPool)

	return dbPool, nil
}

// initTCPConnectionPool initializes a TCP connection pool
func initTCPConnectionPool() (*sql.DB, error) {
	var (
		dbUser    = mustGetenv("DB_USER")
		dbPwd     = mustGetenv("DB_PASS")
		dbTCPHost = mustGetenv("DB_HOST")
		dbPort    = mustGetenv("DB_PORT")
		dbName    = mustGetenv("DB_NAME")
	)

	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", dbTCPHost, dbUser, dbPwd, dbPort, dbName)

	dbPool, err := sql.Open("pgx", dbURI)
	if err != nil {
		return nil, fmt.Errorf("sql.Open: %v", err)
	}

	configureConnectionPool(dbPool)

	return dbPool, nil
}

func configureConnectionPool(dbPool *sql.DB) {
	dbPool.SetMaxIdleConns(5)
	dbPool.SetMaxOpenConns(7)
	dbPool.SetConnMaxLifetime(1800 * time.Second)
}
