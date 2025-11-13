-- name: CreateGroupMessage :one
INSERT INTO group_messages (group_id, sender_id, type, message) VALUES ($1, $2, $3, $4)
RETURNING id, group_id, sender_id, type, message, created_at;

-- name: GetGroupMessagesByGroupID :many
SELECT id, group_id, sender_id, type, message, created_at FROM group_messages WHERE group_id = $1 ORDER BY created_at DESC LIMIT $2 OFFSET $3;