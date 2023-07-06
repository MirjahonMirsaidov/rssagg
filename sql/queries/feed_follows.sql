-- name: CreateFeedFollows :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5)
Returning *;

-- name: GetMyFollows :many
SELECT feed_follows.id, feeds.name, feeds.url FROM feed_follows
    JOIN feeds ON feeds.id = feed_follows.feed_id
                             WHERE feed_follows.user_id = $1;

-- name: Unfollow :exec
DELETE FROM feed_follows WHERE id = $1 AND user_id = $2;