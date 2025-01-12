-- name: GetPermissionByID :one
SELECT * FROM permissions
WHERE id = $1 LIMIT 1;

-- name: GetPermissionByName :one
SELECT * FROM permissions
WHERE name = $1 LIMIT 1;

-- name: CreatePermission :one
INSERT INTO permissions (
  name, description
) VALUES (
  $1, $2
)
RETURNING *;

-- name: UpdatePermission :exec
UPDATE permissions
  SET name = $2,
      description = $3
WHERE id = $1;

-- name: DeletePermission :exec
DELETE FROM permissions
WHERE id = $1;

-- name: ListPermissions :many
SELECT * FROM permissions
ORDER BY name;
