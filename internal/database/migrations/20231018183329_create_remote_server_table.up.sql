CREATE TABLE IF NOT EXISTS remote_servers (
  id SERIAL PRIMARY KEY,
  user_id INT,
  name VARCHAR(255),
  address VARCHAR(255),
  port INT,
  is_active BOOLEAN
);