CREATE TABLE IF NOT EXISTS posts (
  id bigserial PRIMARY KEY,
  title text NOT NULL,
  user_id bigint NOT NULL,
  content text NOT NULL,
  created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  updated_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
  tags VARCHAR(100)[],
  version INT DEFAULT 0,
  CONSTRAINT fk_user FOREIGN KEY (user_id) REFERENCES users (id)
);