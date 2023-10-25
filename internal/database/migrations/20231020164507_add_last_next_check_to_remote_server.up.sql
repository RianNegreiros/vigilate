ALTER TABLE remote_servers
ADD COLUMN last_check_time TIMESTAMP,
ADD COLUMN next_check_time TIMESTAMP;