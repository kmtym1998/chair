package testutil

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/ory/dockertest/v3"
	"github.com/ory/dockertest/v3/docker"
)

type PostgreSQLContainer struct {
	resource *dockertest.Resource
	pool     *dockertest.Pool
	dbOpts   TestDBOptions
}

type TestDBOptions struct {
	User     string
	Password string
	DBName   string
}

func NewPostgreSQLContainer(t *testing.T, ver string, o TestDBOptions) (*PostgreSQLContainer, error) {
	t.Helper()

	c := &PostgreSQLContainer{
		dbOpts: o,
	}

	pool, err := dockertest.NewPool("")
	if err != nil {
		return nil, fmt.Errorf("could not construct pool: %w", err)
	}

	if err := pool.Client.Ping(); err != nil {
		return nil, fmt.Errorf("could not connect to Docker: %w", err)
	}

	c.pool = pool

	resource, err := pool.RunWithOptions(&dockertest.RunOptions{
		Repository: "postgres",
		Tag:        ver,
		Env: []string{
			"POSTGRES_PASSWORD=" + o.Password,
			"POSTGRES_USER=" + o.User,
			"POSTGRES_DB=" + o.DBName,
			"listen_addresses='*'",
		},
	}, func(config *docker.HostConfig) {
		config.AutoRemove = true                                // resource の Purge 後にコンテナを削除する
		config.RestartPolicy = docker.RestartPolicy{Name: "no"} // コンテナの再起動を行わない
	})
	if err != nil {
		return nil, fmt.Errorf("could not start resource: %w", err)
	}

	c.resource = resource

	return c, nil
}

func (c *PostgreSQLContainer) ConnectDB() (*sql.DB, error) {
	if c.pool == nil || c.resource == nil {
		return nil, fmt.Errorf("PostgreSQLcontainer is not initialized")
	}

	databaseURI := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		c.dbOpts.User,
		c.dbOpts.Password,
		c.resource.GetHostPort("5432/tcp"),
		c.dbOpts.DBName,
	)

	// NOTE: The internal application inside the PostgreSQLcontainer may not be ready to accept connections, even if the PostgreSQLcontainer is up and running.
	var db *sql.DB
	c.pool.MaxWait = 20 * time.Second
	if err := c.pool.Retry(func() error {
		innerDB, err := sql.Open("pgx", databaseURI)
		if err != nil {
			return err
		}

		if err := innerDB.Ping(); err != nil {
			return fmt.Errorf("could not ping database: %w", err)
		}

		db = innerDB

		return nil
	}); err != nil {
		return nil, fmt.Errorf("could not connect to docker: %w", err)
	}

	return db, nil
}

func (c *PostgreSQLContainer) Purge() {
	if c.pool == nil || c.resource == nil {
		return
	}

	if err := c.pool.Purge(c.resource); err != nil {
		log.Panicf("could not purge resource: %s", err)
	}
}
