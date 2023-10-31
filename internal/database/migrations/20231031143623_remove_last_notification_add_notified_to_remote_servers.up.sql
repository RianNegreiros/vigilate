ALTER TABLE remote_servers
DROP COLUMN last_notification_time,
ADD COLUMN notified BOOLEAN NOT NULL DEFAULT FALSE;