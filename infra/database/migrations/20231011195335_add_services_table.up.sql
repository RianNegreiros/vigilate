CREATE TABLE IF NOT EXISTS services (
  id bigserial PRIMARY KEY,
  name varchar(255) NOT NULL,
  description text NOT NULL,
  url varchar(255) NOT NULL,
  status varchar(255) NOT NULL,
  created_at timestamptz DEFAULT now(),
  updated_at timestamptz DEFAULT now()
)