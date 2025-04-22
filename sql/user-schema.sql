CREATE TYPE IF NOT EXISTS user_status AS ENUM ('pending', 'active', 'inactive', 'suspend');

CREATE TABLE IF NOT EXISTS users (
    id BIGSERIAL PRIMARY KEY,
    user_uuid UUID DEFAULT gen_random_uuid () NOT NULL,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    phone_number VARCHAR(255) UNIQUE,
    password_hash TEXT NOT NULL,
    password_config JSONB NOT NULL,
    status user_status NOT NULL DEFAULT 'pending',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW (),
    deleted_at TIMESTAMPTZ
);

DROP TRIGGER IF EXISTS update_users_updated_at ON users;

CREATE TRIGGER update_users_updated_at BEFORE
UPDATE ON users FOR EACH ROW EXECUTE FUNCTION update_updated_at_column ();

CREATE INDEX IF NOT EXISTS idx_users_status ON users (status);

CREATE INDEX IF NOT EXISTS idx_users_deleted_at_null ON users (deleted_at)
WHERE
    deleted_at IS NULL;

CREATE INDEX IF NOT EXISTS idx_users_username_not_deleted ON users (username)
WHERE
    deleted_at IS NULL;
