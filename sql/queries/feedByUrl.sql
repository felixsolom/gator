-- name: FeedByUrl :one 
SELECT * FROM feeds
WHERE url=$1; 