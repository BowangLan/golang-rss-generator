// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: users.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createUser = `-- name: CreateUser :one
insert into users (id, created_at, updated_at, first_name, last_name, api_key)
values ($1, $2, $3, $4, $5, encode(sha256(random()::text::bytea), 'hex'))
returning id, created_at, updated_at, first_name, last_name, api_key
`

type CreateUserParams struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	FirstName string
	LastName  string
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser,
		arg.ID,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.FirstName,
		arg.LastName,
	)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FirstName,
		&i.LastName,
		&i.ApiKey,
	)
	return i, err
}

const getUserByApiKey = `-- name: GetUserByApiKey :one
select id, created_at, updated_at, first_name, last_name, api_key from users where api_key = $1
`

func (q *Queries) GetUserByApiKey(ctx context.Context, apiKey string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserByApiKey, apiKey)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FirstName,
		&i.LastName,
		&i.ApiKey,
	)
	return i, err
}

const getUserById = `-- name: GetUserById :one
select id, created_at, updated_at, first_name, last_name, api_key from users where id = $1
`

func (q *Queries) GetUserById(ctx context.Context, id uuid.UUID) (User, error) {
	row := q.db.QueryRowContext(ctx, getUserById, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.FirstName,
		&i.LastName,
		&i.ApiKey,
	)
	return i, err
}

const listUsers = `-- name: ListUsers :many
select id, created_at, updated_at, first_name, last_name, api_key from users
`

func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []User
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.FirstName,
			&i.LastName,
			&i.ApiKey,
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
