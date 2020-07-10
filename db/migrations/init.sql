CREATE TABLE IF NOT EXISTS "chats" (
    "id"             INTEGER NOT NULL PRIMARY KEY,
    "type"           TEXT,
    "title"          TEXT,
    "username"       TEXT,
    "first_name"     TEXT,
    "last_name"      TEXT,
    "description"    TEXT,
    "pinned_message" TEXT
);

CREATE TABLE IF NOT EXISTS "me" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "first_name" TEXT,
    "last_name" TEXT,
    "username" TEXT,
    "language_code" TEXT,
    "can_join_groups" INTEGER,
    "can_read_all_group_messages" INTEGER,
    "supports_inline_queries" INTEGER
);

CREATE TABLE IF NOT EXISTS "messages" (
	"id"                  INTEGER NOT NULL PRIMARY KEY,
    "from_id"             INTEGER NOT NULL,
    "text"                TEXT,
    "date"                INTEGER,
    "reply_to_message_id" INTEGER
);

CREATE TABLE IF NOT EXISTS "users" (
	"id"            INTEGER NOT NULL PRIMARY KEY,
    "is_bot"        INTEGER DEFAULT 0,
    "first_name"	TEXT,
	"last_name"     TEXT,
    "username"      TEXT,
    "language_code" TEXT
);
