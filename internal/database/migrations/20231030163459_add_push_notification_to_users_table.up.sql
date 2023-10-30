ALTER TABLE users
ALTER COLUMN notification_preferences SET DEFAULT '{"email_enabled": false, "push_enabled": false}';
