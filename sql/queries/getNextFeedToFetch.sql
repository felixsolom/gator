-- name: GetNextFeedToFetch :one
SELECT feeds.*
FROM feeds
INNER JOIN feed_follows ON feeds.id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
ORDER BY feeds.last_fetched_at NULLS FIRST
LIMIT 1;
