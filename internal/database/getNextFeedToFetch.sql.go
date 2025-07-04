// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: getNextFeedToFetch.sql

package database

import (
	"context"

	"github.com/google/uuid"
)

const getNextFeedToFetch = `-- name: GetNextFeedToFetch :one
SELECT id, created_at, updated_at, name, url, user_id, last_fetched_at FROM feeds
WHERE user_id=$1
ORDER BY last_fetched_at NULLS FIRST
`

func (q *Queries) GetNextFeedToFetch(ctx context.Context, userID uuid.NullUUID) (Feed, error) {
	row := q.db.QueryRowContext(ctx, getNextFeedToFetch, userID)
	var i Feed
	err := row.Scan(
		&i.ID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Name,
		&i.Url,
		&i.UserID,
		&i.LastFetchedAt,
	)
	return i, err
}
