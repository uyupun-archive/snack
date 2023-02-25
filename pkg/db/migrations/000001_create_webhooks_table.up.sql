CREATE TABLE IF NOT EXISTS webhooks(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    app_name TEXT UNIQUE NOT NULL,
    webhook_url TEXT UNIQUE NOT NULL,
    created_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime')),
    updated_at TEXT NOT NULL DEFAULT (DATETIME('now', 'localtime'))
);

CREATE TRIGGER trigger_webhooks_updated_at AFTER UPDATE ON webhooks
BEGIN
    UPDATE webhooks SET updated_at = DATETIME('now', 'localtime') WHERE rowid == NEW.rowid;
END;
