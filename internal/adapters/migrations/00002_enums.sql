-- +goose Up
DO $$ BEGIN
  IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'task_status') THEN
    CREATE TYPE task_status AS ENUM ('todo','doing','done');
  END IF;
END$$;

-- +goose Down
DO $$ BEGIN
  IF EXISTS (SELECT 1 FROM pg_type WHERE typname = 'task_status') THEN
    DROP TYPE task_status;
  END IF;
END$$;
