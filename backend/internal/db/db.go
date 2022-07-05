package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
)

func mustGetenv(k string) string {
	v := viper.GetString(k)
	if v == "" {
		log.Fatalf("Warning: %s environment variable not set.\n", k)
	}
	return v
}

func New() (*pgxpool.Pool, error) {
	if viper.GetString("DB_SOCKET_DIR") != "" {
		return initSocketConnectionPool()
	} else {
		return initTCPConnectionPool()
	}
}

// initSocketConnectionPool initializes a Unix socket connection pool
func initSocketConnectionPool() (*pgxpool.Pool, error) {
	var (
		dbUser                 = mustGetenv("DB_USER")
		dbPwd                  = mustGetenv("DB_PASS")
		instanceConnectionName = mustGetenv("INSTANCE_CONNECTION_NAME") // e.g. 'project:region:instance'
		dbName                 = mustGetenv("DB_NAME")
		socketDir              = mustGetenv("DB_SOCKET_DIR")
	)

	dbURI := fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)

	dbPool, err := pgxpool.Connect(context.Background(), dbURI)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.Connect: %v", err)
	}

	return dbPool, nil
}

// initTCPConnectionPool initializes a TCP connection pool
func initTCPConnectionPool() (*pgxpool.Pool, error) {
	var (
		dbUser    = mustGetenv("DB_USER")
		dbPwd     = mustGetenv("DB_PASS")
		dbTCPHost = mustGetenv("DB_HOST")
		dbPort    = mustGetenv("DB_PORT")
		dbName    = mustGetenv("DB_NAME")
	)

	dbURI := fmt.Sprintf("host=%s user=%s password=%s port=%s database=%s", dbTCPHost, dbUser, dbPwd, dbPort, dbName)

	dbPool, err := pgxpool.Connect(context.Background(), dbURI)
	if err != nil {
		return nil, fmt.Errorf("pgxpool.Connect: %v", err)
	}

	return dbPool, nil
}
