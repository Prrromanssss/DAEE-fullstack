-- name: CreateOperation :exec
INSERT INTO operations (operation_type, execution_time)
VALUES ($1, $2)
RETURNING
    operation_id, operation_type, execution_time;

-- name: UpdateOperationTime :one
UPDATE operations
SET execution_time = $1
WHERE operation_type = $2
RETURNING *;

-- name: GetOperations :many
SELECT
    operation_id, operation_type, execution_time
FROM operations
ORDER BY operation_type DESC;

-- name: GetOperationTimeByType :one
SELECT execution_time FROM operations
WHERE operation_type = $1;

-- name: GetOperationByType :one
SELECT
    operation_id, operation_type, execution_time
FROM operations
WHERE operation_type = $1;
