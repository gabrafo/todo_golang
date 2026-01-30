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

-- name: CreateUser :one
INSERT INTO users (name, email, password_hash)
VALUES ($1, $2, $3)
RETURNING *;

-- name: CreateCategory :one
INSERT INTO categories (name)
VALUES ($1)
RETURNING *;

-- name: CreateTask :one
INSERT INTO tasks (user_id, name, status)
VALUES ($1, $2, $3)
RETURNING *;

-- name: AddTaskCategory :exec
INSERT INTO task_categories (task_id, category_id)
VALUES ($1, $2);