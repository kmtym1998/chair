package client

import (
	"context"
	"database/sql"
	"time"

	"github.com/cockroachdb/errors"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Opts struct {
	DataSourceName string
	MaxOpenConns   int
	MaxIdleConns   int
	MaxLifetime    time.Duration
}

type client struct {
	sqlDB *sql.DB
}

type Client interface {
	Close() error
	DB() *sql.DB
	Tx(ctx context.Context, level sql.IsolationLevel, do func(tx *sql.Tx) error) error
}

func New(o Opts) (*client, error) {
	db, err := sql.Open("pgx", o.DataSourceName)
	if err != nil {
		return nil, errors.Wrap(err, "error in opening db")
	}

	db.SetMaxOpenConns(o.MaxOpenConns)
	db.SetMaxIdleConns(o.MaxIdleConns)
	db.SetConnMaxLifetime(o.MaxLifetime)

	if err := db.Ping(); err != nil {
		return nil, errors.Wrap(err, "error in pinging db")
	}

	return &client{
		sqlDB: db,
	}, nil
}

func (c *client) Close() error {
	if c.sqlDB == nil {
		return nil
	}

	if err := c.sqlDB.Close(); err != nil {
		return errors.Wrap(err, "error in closing db")
	}

	return nil
}

func (c *client) DB() *sql.DB {
	return c.sqlDB
}
