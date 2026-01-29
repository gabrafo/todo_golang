-- +goose Up
CREATE TYPE task_status AS ENUM ('todo', 'doing', 'done');

-- +goose Down
DROP TYPE task_status;