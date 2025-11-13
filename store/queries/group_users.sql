-- name: AddUserToGroup :one
INSERT INTO user_groups (user_id, group_id) VALUES ($1, $2)
RETURNING user_id, group_id, added_at;

-- name: RemoveUserFromGroup :exec
UPDATE user_groups SET removed_at = CURRENT_TIMESTAMP WHERE user_id = $1 AND group_id = $2;

-- name: GetUserIDInGroup :one
SELECT user_id FROM user_groups WHERE user_id = $1 AND group_id = $2 AND removed_at IS NULL;