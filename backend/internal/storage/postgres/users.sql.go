// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: users.sql

package postgres

import (
	"context"
)

const getUser = `-- name: GetUser :one
SELECT user_id, email, password_hash
FROM users
WHERE user_id = $1
`

func (q *Queries) GetUser(ctx context.Context, userID int32) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, userID)
	var i User
	err := row.Scan(&i.UserID, &i.Email, &i.PasswordHash)
	return i, err
}

const saveUser = `-- name: SaveUser :one
INSERT INTO users
    (email, password_hash)
VALUES
    ($1, $2)
RETURNING user_id
`

type SaveUserParams struct {
	Email        string
	PasswordHash []byte
}

func (q *Queries) SaveUser(ctx context.Context, arg SaveUserParams) (int32, error) {
	row := q.db.QueryRowContext(ctx, saveUser, arg.Email, arg.PasswordHash)
	var user_id int32
	err := row.Scan(&user_id)
	return user_id, err
}
