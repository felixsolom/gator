-- name: GetPostsForUser :many
SELECT posts.*
FROM posts 
INNER JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id=$1
ORDER BY posts.updated_at DESC 
LIMIT $2; 

