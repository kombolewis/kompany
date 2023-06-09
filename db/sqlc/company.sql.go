// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: company.sql

package db

import (
	"context"
	"database/sql"
)

const createCompany = `-- name: CreateCompany :one
INSERT INTO company (
  name, 
  description,
  amount,
  registered,
  type
) VALUES (
  $1, $2,$3,$4,$5
)
RETURNING id, name, description, amount, registered, type, created_at
`

type CreateCompanyParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Amount      int32          `json:"amount"`
	Registered  bool           `json:"registered"`
	Type        string         `json:"type"`
}

func (q *Queries) CreateCompany(ctx context.Context, arg CreateCompanyParams) (Company, error) {
	row := q.db.QueryRowContext(ctx, createCompany,
		arg.Name,
		arg.Description,
		arg.Amount,
		arg.Registered,
		arg.Type,
	)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Amount,
		&i.Registered,
		&i.Type,
		&i.CreatedAt,
	)
	return i, err
}

const deleteCompany = `-- name: DeleteCompany :exec
DELETE FROM company
WHERE id = $1
`

func (q *Queries) DeleteCompany(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deleteCompany, id)
	return err
}

const getCompany = `-- name: GetCompany :one
SELECT id, name, description, amount, registered, type, created_at FROM company
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetCompany(ctx context.Context, id int64) (Company, error) {
	row := q.db.QueryRowContext(ctx, getCompany, id)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Amount,
		&i.Registered,
		&i.Type,
		&i.CreatedAt,
	)
	return i, err
}

const listCompanies = `-- name: ListCompanies :many
SELECT id, name, description, amount, registered, type, created_at FROM company
ORDER BY name
LIMIT $1
OFFSET $2
`

type ListCompaniesParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListCompanies(ctx context.Context, arg ListCompaniesParams) ([]Company, error) {
	rows, err := q.db.QueryContext(ctx, listCompanies, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Company{}
	for rows.Next() {
		var i Company
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Amount,
			&i.Registered,
			&i.Type,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateCompany = `-- name: UpdateCompany :one
UPDATE company
  set name = $2,
  description = $3,
  amount = $4,
  registered = $5,
  type = $6
WHERE id = $1
RETURNING id, name, description, amount, registered, type, created_at
`

type UpdateCompanyParams struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	Amount      int32          `json:"amount"`
	Registered  bool           `json:"registered"`
	Type        string         `json:"type"`
}

func (q *Queries) UpdateCompany(ctx context.Context, arg UpdateCompanyParams) (Company, error) {
	row := q.db.QueryRowContext(ctx, updateCompany,
		arg.ID,
		arg.Name,
		arg.Description,
		arg.Amount,
		arg.Registered,
		arg.Type,
	)
	var i Company
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.Amount,
		&i.Registered,
		&i.Type,
		&i.CreatedAt,
	)
	return i, err
}
