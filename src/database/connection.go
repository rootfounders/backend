package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Connection struct {
	Host      string
	Port      int
	User      string
	Password  string
	Database  string
	Chain     uint
	conn      *pgx.Conn
	batchSize int
}

func (c *Connection) Connect(ctx context.Context) (err error) {
	if err := c.validate(); err != nil {
		return err
	}

	c.conn, err = pgx.Connect(ctx, c.dsn())
	if err != nil {
		return fmt.Errorf("connection.Connect: %w", err)
	}
	return
}

func (c *Connection) Close(ctx context.Context) error {
	return c.conn.Close(ctx)
}

func (c *Connection) Conn() *pgx.Conn {
	return c.conn
}

func (c *Connection) String() string {
	var pass string
	if len(c.Password) > 0 {
		pass = "xxx"
	}
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", c.Host, c.User, pass, c.Database, c.Port)
}

func (c *Connection) dsn() string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", c.Host, c.User, c.Password, c.Database, c.Port)
}

func (c *Connection) validate() error {
	if c.Chain < 1 {
		return fmt.Errorf("invalid chain id %d", c.Chain)
	}

	// Ports below 1024 are reserved, so it would be strange
	// if database ran there
	if c.Port < 1024 {
		return fmt.Errorf("invalid database port %d", c.Port)
	}

	// We need host, even if it's localhost
	if c.Host == "" {
		return fmt.Errorf("database host missing")
	}

	// We need database name
	if c.Database == "" {
		return fmt.Errorf("database name missing")
	}

	return nil
}
