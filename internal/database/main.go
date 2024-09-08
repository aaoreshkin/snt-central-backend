package database

import (
	"context"
	"net"
	"os"
	"runtime"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

type (
	Connection struct {
		*pgxpool.Pool
	}
)

var (
	dbpoolConfig *pgxpool.Config
	dbpool       *pgxpool.Pool

	err error
)

// New creates a new database connection pool.
// Args:
//
//	ctx (context.Context): The context in which the connection pool is created.
//
// Returns:
//
//	(*pgxpool.Pool, error): The newly created connection pool and an error if any.
func New(ctx context.Context) (*Connection, error) {
	if dbpoolConfig, err = pgxpool.ParseConfig(os.Getenv("DATABASE_URL")); err != nil {
		return nil, err
	}

	// MaxConns is the maximum number of database connections that can be opened in the connection pool at one time.
	// If all connections are in use, new requests will wait for a connection to be released or will be rejected if the limit is reached.
	// The default value is 10 times the number of CPU cores.
	//
	// Source: https://pkg.go.dev/github.com/jackc/pgx/v5/pgxpool#Config
	dbpoolConfig.MaxConns = int32(runtime.NumCPU() * 5)

	// MinConns is the minimum number of database connections that should always be open and ready to be used in the connection pool.
	// This value ensures that even under low load, your application will maintain the specified number of always-active connections
	// to ensure a fast response to incoming requests.
	//
	// Source: https://pkg.go.dev/github.com/jackc/pgx/v5/pgxpool#Config
	dbpoolConfig.MinConns = 5

	// HealthCheckPeriod is the time period over which the connection pool will perform a health check for each open connection.
	// This helps to ensure that faulty or hung connections are detected and closed, and replaced with new connections if necessary,
	// to keep the connection pool up and running.
	//
	// Source: https://pkg.go.dev/github.com/jackc/pgx/v5/pgxpool#Config
	dbpoolConfig.HealthCheckPeriod = 1 * time.Minute

	// MaxConnLifetime is the maximum lifetime of a connection in the pool.
	// When this time expires, the connection is closed and removed from the pool,
	// even if it is still active. This prevents potential problems associated with
	// long-lived connections, such as memory leaks or outdated settings.
	//
	// Source: https://pkg.go.dev/github.com/jackc/pgx/v5/pgxpool#Config
	dbpoolConfig.MaxConnLifetime = 12 * time.Hour

	// MaxConnIdleTime is the maximum amount of time a connection is idle after which
	// the connection will be closed and removed from the pool. This helps free up
	// resources that could have been spent on maintaining a connection that has been
	// idle for a long time.
	//
	// Source: https://pkg.go.dev/github.com/jackc/pgx/v5/pgxpool#Config
	dbpoolConfig.MaxConnIdleTime = 15 * time.Minute

	// ConnectTimeout is the maximum time to wait for a connection to be established to the database.
	// If the connection cannot be established within the specified time, the connection process will be aborted with an error.
	//
	// Source: https://pkg.go.dev/github.com/jackc/pgx/v5/pgxpool#ConnConfig
	dbpoolConfig.ConnConfig.ConnectTimeout = 10 * time.Second

	dbpoolConfig.ConnConfig.DialFunc = (&net.Dialer{
		// KeepAlive specifies the time interval at which keep-alive messages will be sent to maintain the connection.
		// This helps to ensure that the connection is not closed by a network device or server due to inactivity.
		//
		// Source: https://pkg.go.dev/net#Dialer.KeepAlive
		KeepAlive: dbpoolConfig.HealthCheckPeriod,

		// Timeout is the maximum time to wait for a connection to be established before the operation is terminated by timeout.
		//
		// Source: https://pkg.go.dev/net#Dialer.Timeout
		Timeout: dbpoolConfig.ConnConfig.ConnectTimeout,
	}).DialContext

	// Create a new connection pool with the specified configuration.
	//
	// Args:
	//   ctx (context.Context): The context in which the connection pool is created.
	//
	// Returns:
	//   (*pgxpool.Pool, error): The newly created connection pool and an error if any.
	if dbpool, err = pgxpool.NewWithConfig(ctx, dbpoolConfig); err != nil {
		return nil, err
	}

	return &Connection{
		dbpool,
	}, nil
}
