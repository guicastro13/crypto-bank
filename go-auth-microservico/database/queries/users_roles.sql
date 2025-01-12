-- name: AddRoleToUser :exec
INSERT INTO user_roles (
  user_id, role_id
) VALUES (
  $1, $2
);

-- name: RemoveRoleFromUser :exec
DELETE FROM user_roles
WHERE user_id = $1 AND role_id = $2;

-- name: GetRolesForUser :many
SELECT r.* FROM roles r
JOIN user_roles ur ON r.id = ur.role_id
WHERE ur.user_id = $1;

-- name: GetUsersForRole :many
SELECT u.* FROM users u
JOIN user_roles ur ON u.id = ur.user_id
WHERE ur.role_id = $1;
