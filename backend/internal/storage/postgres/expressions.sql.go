// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: expressions.sql

package postgres

import (
	"context"
	"database/sql"
	"time"
)

const createExpression = `-- name: CreateExpression :one
INSERT INTO expressions
    (created_at, updated_at, data, parse_data, status, user_id)
VALUES
    ($1, $2, $3, $4, $5, $6)
RETURNING
    expression_id, user_id, agent_id,
    created_at, updated_at, data, parse_data,
    status, result, is_ready
`

type CreateExpressionParams struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	Data      string
	ParseData string
	Status    ExpressionStatus
	UserID    int32
}

func (q *Queries) CreateExpression(ctx context.Context, arg CreateExpressionParams) (Expression, error) {
	row := q.db.QueryRowContext(ctx, createExpression,
		arg.CreatedAt,
		arg.UpdatedAt,
		arg.Data,
		arg.ParseData,
		arg.Status,
		arg.UserID,
	)
	var i Expression
	err := row.Scan(
		&i.ExpressionID,
		&i.UserID,
		&i.AgentID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Data,
		&i.ParseData,
		&i.Status,
		&i.Result,
		&i.IsReady,
	)
	return i, err
}

const getComputingExpressions = `-- name: GetComputingExpressions :many
SELECT
    expression_id, user_id, agent_id,
    created_at, updated_at, data, parse_data,
    status, result, is_ready
FROM expressions
WHERE status IN ('ready_for_computation', 'computing', 'terminated')
ORDER BY created_at DESC
`

func (q *Queries) GetComputingExpressions(ctx context.Context) ([]Expression, error) {
	rows, err := q.db.QueryContext(ctx, getComputingExpressions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expression
	for rows.Next() {
		var i Expression
		if err := rows.Scan(
			&i.ExpressionID,
			&i.UserID,
			&i.AgentID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Data,
			&i.ParseData,
			&i.Status,
			&i.Result,
			&i.IsReady,
		); err != nil {
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

const getExpressionByID = `-- name: GetExpressionByID :one
SELECT
    expression_id, user_id, agent_id,
    created_at, updated_at, data, parse_data,
    status, result, is_ready
FROM expressions
WHERE expression_id = $1
`

func (q *Queries) GetExpressionByID(ctx context.Context, expressionID int32) (Expression, error) {
	row := q.db.QueryRowContext(ctx, getExpressionByID, expressionID)
	var i Expression
	err := row.Scan(
		&i.ExpressionID,
		&i.UserID,
		&i.AgentID,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.Data,
		&i.ParseData,
		&i.Status,
		&i.Result,
		&i.IsReady,
	)
	return i, err
}

const getExpressions = `-- name: GetExpressions :many
SELECT
    expression_id, user_id, agent_id,
    created_at, updated_at, data, parse_data,
    status, result, is_ready
FROM expressions
ORDER BY created_at DESC
`

func (q *Queries) GetExpressions(ctx context.Context) ([]Expression, error) {
	rows, err := q.db.QueryContext(ctx, getExpressions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expression
	for rows.Next() {
		var i Expression
		if err := rows.Scan(
			&i.ExpressionID,
			&i.UserID,
			&i.AgentID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Data,
			&i.ParseData,
			&i.Status,
			&i.Result,
			&i.IsReady,
		); err != nil {
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

const getTerminatedExpressions = `-- name: GetTerminatedExpressions :many
SELECT
    expression_id, user_id, agent_id,
    created_at, updated_at, data, parse_data,
    status, result, is_ready
FROM expressions
WHERE status = 'terminated'
ORDER BY created_at DESC
`

func (q *Queries) GetTerminatedExpressions(ctx context.Context) ([]Expression, error) {
	rows, err := q.db.QueryContext(ctx, getTerminatedExpressions)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Expression
	for rows.Next() {
		var i Expression
		if err := rows.Scan(
			&i.ExpressionID,
			&i.UserID,
			&i.AgentID,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Data,
			&i.ParseData,
			&i.Status,
			&i.Result,
			&i.IsReady,
		); err != nil {
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

const makeExpressionReady = `-- name: MakeExpressionReady :exec
UPDATE expressions
SET parse_data = $1, result = $2, updated_at = $3, is_ready = True, status = 'result'
WHERE expression_id = $4
`

type MakeExpressionReadyParams struct {
	ParseData    string
	Result       int32
	UpdatedAt    time.Time
	ExpressionID int32
}

func (q *Queries) MakeExpressionReady(ctx context.Context, arg MakeExpressionReadyParams) error {
	_, err := q.db.ExecContext(ctx, makeExpressionReady,
		arg.ParseData,
		arg.Result,
		arg.UpdatedAt,
		arg.ExpressionID,
	)
	return err
}

const makeExpressionsTerminated = `-- name: MakeExpressionsTerminated :exec
UPDATE expressions
SET status = 'terminated'
WHERE agent_id = $1
`

func (q *Queries) MakeExpressionsTerminated(ctx context.Context, agentID sql.NullInt32) error {
	_, err := q.db.ExecContext(ctx, makeExpressionsTerminated, agentID)
	return err
}

const updateExpressionData = `-- name: UpdateExpressionData :exec
UPDATE expressions
SET data = $1
WHERE expression_id = $2
`

type UpdateExpressionDataParams struct {
	Data         string
	ExpressionID int32
}

func (q *Queries) UpdateExpressionData(ctx context.Context, arg UpdateExpressionDataParams) error {
	_, err := q.db.ExecContext(ctx, updateExpressionData, arg.Data, arg.ExpressionID)
	return err
}

const updateExpressionParseData = `-- name: UpdateExpressionParseData :exec
UPDATE expressions
SET parse_data = $1
WHERE expression_id = $2
`

type UpdateExpressionParseDataParams struct {
	ParseData    string
	ExpressionID int32
}

func (q *Queries) UpdateExpressionParseData(ctx context.Context, arg UpdateExpressionParseDataParams) error {
	_, err := q.db.ExecContext(ctx, updateExpressionParseData, arg.ParseData, arg.ExpressionID)
	return err
}

const updateExpressionStatus = `-- name: UpdateExpressionStatus :exec
UPDATE expressions
SET status = $1
WHERE expression_id = $2
`

type UpdateExpressionStatusParams struct {
	Status       ExpressionStatus
	ExpressionID int32
}

func (q *Queries) UpdateExpressionStatus(ctx context.Context, arg UpdateExpressionStatusParams) error {
	_, err := q.db.ExecContext(ctx, updateExpressionStatus, arg.Status, arg.ExpressionID)
	return err
}
