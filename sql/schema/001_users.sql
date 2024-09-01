
-- +goose Up
create table users (
  id UUID PRIMARY KEY,
  created_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  updated_at TIMESTAMP NOT NULL DEFAULT current_timestamp,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL
);

-- +goose Down
drop table users;
