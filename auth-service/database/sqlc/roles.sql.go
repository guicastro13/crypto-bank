// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: roles.sql

package sqlc

import (
	"context"
)

const createRole = `-- name: CreateRole :one
INSERT INTO roles (
  name
) VALUES (
  $1
)
RETURNING id, name
`

func (q *Queries) CreateRole(ctx context.Context, name string) (Role, error) {
	row := q.db.QueryRowContext(ctx, createRole, name)
	var i Role
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const deleteRole = `-- name: DeleteRole :exec
DELETE FROM roles
WHERE id = $1
`

func (q *Queries) DeleteRole(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteRole, id)
	return err
}

const getRoleByID = `-- name: GetRoleByID :one
SELECT id, name FROM roles
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetRoleByID(ctx context.Context, id int32) (Role, error) {
	row := q.db.QueryRowContext(ctx, getRoleByID, id)
	var i Role
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const getRoleByName = `-- name: GetRoleByName :one
SELECT id, name FROM roles
WHERE name = $1 LIMIT 1
`

func (q *Queries) GetRoleByName(ctx context.Context, name string) (Role, error) {
	row := q.db.QueryRowContext(ctx, getRoleByName, name)
	var i Role
	err := row.Scan(&i.ID, &i.Name)
	return i, err
}

const listRoles = `-- name: ListRoles :many
SELECT id, name FROM roles
ORDER BY name
`

func (q *Queries) ListRoles(ctx context.Context) ([]Role, error) {
	rows, err := q.db.QueryContext(ctx, listRoles)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Role
	for rows.Next() {
		var i Role
		if err := rows.Scan(&i.ID, &i.Name); err != nil {
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

const updateRole = `-- name: UpdateRole :exec
UPDATE roles
  SET name = $2
WHERE id = $1
`

type UpdateRoleParams struct {
	ID   int32
	Name string
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) error {
	_, err := q.db.ExecContext(ctx, updateRole, arg.ID, arg.Name)
	return err
}