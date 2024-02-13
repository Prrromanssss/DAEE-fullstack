// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: agents.sql

package database

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const createAgent = `-- name: CreateAgent :one
INSERT INTO agents (id, created_at, number_of_parallel_calculations, last_ping, status)
VALUES ($1, $2, $3, $4, $5)
RETURNING id, number_of_parallel_calculations, last_ping, status, created_at
`

type CreateAgentParams struct {
	ID                           uuid.UUID
	CreatedAt                    time.Time
	NumberOfParallelCalculations int32
	LastPing                     time.Time
	Status                       AgentStatus
}

func (q *Queries) CreateAgent(ctx context.Context, arg CreateAgentParams) (Agent, error) {
	row := q.db.QueryRowContext(ctx, createAgent,
		arg.ID,
		arg.CreatedAt,
		arg.NumberOfParallelCalculations,
		arg.LastPing,
		arg.Status,
	)
	var i Agent
	err := row.Scan(
		&i.ID,
		&i.NumberOfParallelCalculations,
		&i.LastPing,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getAgentByID = `-- name: GetAgentByID :one
SELECT id, number_of_parallel_calculations, last_ping, status, created_at FROM agents
WHERE id = $1
`

func (q *Queries) GetAgentByID(ctx context.Context, id uuid.UUID) (Agent, error) {
	row := q.db.QueryRowContext(ctx, getAgentByID, id)
	var i Agent
	err := row.Scan(
		&i.ID,
		&i.NumberOfParallelCalculations,
		&i.LastPing,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getAgents = `-- name: GetAgents :many
SELECT id, number_of_parallel_calculations, last_ping, status, created_at FROM agents
`

func (q *Queries) GetAgents(ctx context.Context) ([]Agent, error) {
	rows, err := q.db.QueryContext(ctx, getAgents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Agent
	for rows.Next() {
		var i Agent
		if err := rows.Scan(
			&i.ID,
			&i.NumberOfParallelCalculations,
			&i.LastPing,
			&i.Status,
			&i.CreatedAt,
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

const updateAgentLastPing = `-- name: UpdateAgentLastPing :one
UPDATE agents
SET last_ping = $1
WHERE id = $2
RETURNING id, number_of_parallel_calculations, last_ping, status, created_at
`

type UpdateAgentLastPingParams struct {
	LastPing time.Time
	ID       uuid.UUID
}

func (q *Queries) UpdateAgentLastPing(ctx context.Context, arg UpdateAgentLastPingParams) (Agent, error) {
	row := q.db.QueryRowContext(ctx, updateAgentLastPing, arg.LastPing, arg.ID)
	var i Agent
	err := row.Scan(
		&i.ID,
		&i.NumberOfParallelCalculations,
		&i.LastPing,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const updateAgentStatus = `-- name: UpdateAgentStatus :one
UPDATE agents
SET status = $1
WHERE id = $2
RETURNING id, number_of_parallel_calculations, last_ping, status, created_at
`

type UpdateAgentStatusParams struct {
	Status AgentStatus
	ID     uuid.UUID
}

func (q *Queries) UpdateAgentStatus(ctx context.Context, arg UpdateAgentStatusParams) (Agent, error) {
	row := q.db.QueryRowContext(ctx, updateAgentStatus, arg.Status, arg.ID)
	var i Agent
	err := row.Scan(
		&i.ID,
		&i.NumberOfParallelCalculations,
		&i.LastPing,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}
