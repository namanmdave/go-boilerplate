-- name: CreateGroup :one
INSERT INTO groups (group_id, name, description) VALUES ($1, $2, $3)
RETURNING id, group_id, name, description, created_at, updated_at;

-- name: GetGroup :one
SELECT id, group_id, name, description, created_at, updated_at FROM groups WHERE group_id = $1;