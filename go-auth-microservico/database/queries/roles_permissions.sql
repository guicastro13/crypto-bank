-- name: AddPermissionToRole :exec
INSERT INTO role_permissions (
  role_id, permission_id
) VALUES (
  $1, $2
);

-- name: RemovePermissionFromRole :exec
DELETE FROM role_permissions
WHERE role_id = $1 AND permission_id = $2;

-- name: GetPermissionsForRole :many
SELECT p.* FROM permissions p
JOIN role_permissions rp ON p.id = rp.permission_id
WHERE rp.role_id = $1;

-- name: GetRolesForPermission :many
SELECT r.* FROM roles r
JOIN role_permissions rp ON r.id = rp.role_id
WHERE rp.permission_id = $1;
