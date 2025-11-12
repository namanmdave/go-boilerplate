-- name: GetUser :one
SELECT id, name, email, created_at, updated_at FROM users WHERE id = $1;

-- name: ListUsers :many
SELECT id, name, email, created_at, updated_at FROM users;

-- name: CreateUser :one
INSERT INTO users (name, email) VALUES ($1, $2)
RETURNING id, name, email, created_at, updated_at;

-- name: UpdateUser :exec
UPDATE users SET name = $1, email = $2, updated_at = CURRENT_TIMESTAMP WHERE id = $3;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;
