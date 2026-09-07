package simple_connection

import (
	"context"

	"github.com/jackc/pgx/v5"
)



func CreateConnection(ctx context.Context) (*pgx.Conn, error) {
	return pgx.Connect(ctx, "postgres://postgres:pass@localhost:15432/postgres")
}