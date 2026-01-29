-- name: ListTasks :many
SELECT
    *
FROM 
    tasks;

-- name: FindTaskById :one
SELECT
    *
FROM
    tasks
WHERE
    id = $1;