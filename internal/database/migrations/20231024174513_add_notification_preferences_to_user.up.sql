ALTER TABLE users
ADD COLUMN notification_preferences JSONB DEFAULT '{"email_enabled": false}';
