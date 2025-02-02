// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: permissions.sql

package sqlc

import (
	"context"
	"database/sql"
)

const createPermission = `-- name: CreatePermission :one
INSERT INTO permissions (
  name, description
) VALUES (
  $1, $2
)
RETURNING id, name, description
`

type CreatePermissionParams struct {
	Name        string
	Description sql.NullString
}

func (q *Queries) CreatePermission(ctx context.Context, arg CreatePermissionParams) (Permission, error) {
	row := q.db.QueryRowContext(ctx, createPermission, arg.Name, arg.Description)
	var i Permission
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const deletePermission = `-- name: DeletePermission :exec
DELETE FROM permissions
WHERE id = $1
`

func (q *Queries) DeletePermission(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deletePermission, id)
	return err
}

const getPermissionByID = `-- name: GetPermissionByID :one
SELECT id, name, description FROM permissions
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetPermissionByID(ctx context.Context, id int32) (Permission, error) {
	row := q.db.QueryRowContext(ctx, getPermissionByID, id)
	var i Permission
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const getPermissionByName = `-- name: GetPermissionByName :one
SELECT id, name, description FROM permissions
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetPermissionByName(ctx context.Context, name string) (Permission, error) {
	row := q.db.QueryRowContext(ctx, getPermissionByName, name)
	var i Permission
	err := row.Scan(&i.ID, &i.Name, &i.Description)
	return i, err
}

const listPermissions = `-- name: ListPermissions :many
SELECT id, name, description FROM permissions
ORDER BY name
`

func (q *Queries) ListPermissions(ctx context.Context) ([]Permission, error) {
	rows, err := q.db.QueryContext(ctx, listPermissions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Permission
	for rows.Next() {
		var i Permission
		if err := rows.Scan(&i.ID, &i.Name, &i.Description); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updatePermission = `-- name: UpdatePermission :exec
UPDATE permissions
  SET name = $2,
      description = $3
WHERE id = $1
`

type UpdatePermissionParams struct {
	ID          int32
	Name        string
	Description sql.NullString
}

func (q *Queries) UpdatePermission(ctx context.Context, arg UpdatePermissionParams) error {
	_, err := q.db.ExecContext(ctx, updatePermission, arg.ID, arg.Name, arg.Description)
	return err
}
