// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: listUsers.sql

package database

import (
	"context"
)

const listUsers = `-- name: ListUsers :many
SELECT name FROM users
`

func (q *Queries) ListUsers(ctx context.Context) ([]string, error) {
	rows, err := q.db.QueryContext(ctx, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []string
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		items = append(items, name)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
