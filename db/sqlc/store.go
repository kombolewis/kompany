package db

import "database/sql"

type Store interface {
	Querier
}

type SQLstore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) *SQLstore {
	return &SQLstore{
		db:      db,
		Queries: New(db),
	}
}
