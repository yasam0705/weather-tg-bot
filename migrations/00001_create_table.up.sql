CREATE TABLE IF NOT EXISTS "client" (
    "guid"             UUID,
    "client_id"        BIGINT UNIQUE NOT NULL DEFAULT 0,
    "first_name"       VARCHAR NOT NULL DEFAULT '',
    "last_name"        VARCHAR NOT NULL DEFAULT '',
    "username"         VARCHAR NOT NULL DEFAULT '',
    "created_at"       TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    "updated_at"       TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT client_guid_pk PRIMARY KEY ("guid")
);

CREATE TABLE IF NOT EXISTS "message" (
    "guid"             UUID,
    "client_id"        BIGINT NOT NULL,
    "text"             VARCHAR NOT NULL DEFAULT '',
    "created_at"       TIMESTAMP(0) WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT message_guid_pk PRIMARY KEY ("guid")
);
