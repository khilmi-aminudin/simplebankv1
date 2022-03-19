// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package db

import (
	"context"
)

const createUsers = `-- name: CreateUsers :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
) VALUES (
    $1, $2, $3, $4
) RETURNINg username, hashed_password, full_name, email, password_changed_at, created_at
`

type CreateUsersParams struct {
	Username       string `json:"username"`
	HashedPassword string `json:"hashed_password"`
	FullName       string `json:"full_name"`
	Email          string `json:"email"`
}

func (q *Queries) CreateUsers(ctx context.Context, arg CreateUsersParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUsers,
		arg.Username,
		arg.HashedPassword,
		arg.FullName,
		arg.Email,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUsers = `-- name: GetUsers :one
SELECT username, hashed_password, full_name, email, password_changed_at, created_at FROM users
WHERE username = $1 LIMIT 1
`

func (q *Queries) GetUsers(ctx context.Context, username string) (User, error) {
	row := q.db.QueryRowContext(ctx, getUsers, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.HashedPassword,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}