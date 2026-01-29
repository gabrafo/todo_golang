-- +goose Up
CREATE TABLE IF NOT EXISTS task_categories (
  task_id UUID NOT NULL,
  category_id UUID NOT NULL,

  PRIMARY KEY (task_id, category_id),

  CONSTRAINT fk_task_categories_task
    FOREIGN KEY (task_id)
    REFERENCES tasks(id)
    ON DELETE CASCADE,

  CONSTRAINT fk_task_categories_category
    FOREIGN KEY (category_id)
    REFERENCES categories(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE IF EXISTS task_categories;
