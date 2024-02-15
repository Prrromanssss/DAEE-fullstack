-- name: UpdateOperationTime :one
UPDATE operations
SET execution_time = $1
WHERE operation_type = $2
RETURNING *;

-- name: GetOperations :many
SELECT * FROM operations;

-- name: GetOperationTimeByType :one
SELECT execution_time FROM operations
WHERE operation_type = $1;
