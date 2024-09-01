-- name: CreateUser :one
insert into users (id, created_at, updated_at, first_name, last_name, api_key)
values ($1, $2, $3, $4, $5, encode(sha256(random()::text::bytea), 'hex'))
returning *;

-- name: ListUsers :many
select * from users;

-- name: GetUserById :one
select * from users where id = $1;

-- name: GetUserByApiKey :one
select * from users where api_key = $1;