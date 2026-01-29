-- +goose Up
CREATE INDEX IF NOT EXISTS idx_tasks_user_id
  ON tasks(user_id);

CREATE INDEX IF NOT EXISTS idx_task_categories_task_id
  ON task_categories(task_id);

CREATE INDEX IF NOT EXISTS idx_task_categories_category_id
  ON task_categories(category_id);

-- +goose Down
DROP INDEX IF EXISTS idx_task_categories_category_id;
DROP INDEX IF EXISTS idx_task_categories_task_id;
DROP INDEX IF EXISTS idx_tasks_user_id;
