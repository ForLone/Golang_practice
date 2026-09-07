package simple_sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)



func CreateTable(ctx context.Context, conn *pgx.Conn) error{
	sqlQuery := `
		CREATE TABLE IF NOT EXISTS tasks (
			id SERIAL PRIMARY KEY,
			title VARCHAR(200) NOT NULL,
			description VARCHAR(1000) NOT NULL,
			completed BOOLEAN NOT NULL,
			created_at TIMESTAMP NOT NULL,
			completed_at TIMESTAMP,

			UNIQUE(title)
		);
	`
	if _, err :=conn.Exec(ctx, sqlQuery); err != nil{
		return err
	}
	return nil
}

func AlterTable(ctx context.Context, conn *pgx.Conn) error{
	sqlQuery := `
		ALTER TABLE tasks ADD email VARCHAR(200) UNIQUE;
	`
	if _, err := conn.Exec(ctx, sqlQuery); err != nil{
		return err
	}
	
	return nil
}