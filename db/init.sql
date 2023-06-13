DROP TABLE IF EXISTS settings;
DROP TABLE IF EXISTS poll;

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";


CREATE TABLE IF NOT EXISTS settings (
    "key"                varchar  NOT NULL,
    "value"              varchar    ,
    CONSTRAINT pk_settings_key PRIMARY KEY ("key")
   );

CREATE TABLE IF NOT EXISTS polls (
    "id"   uuid DEFAULT uuid_generate_v4() NOT NULL,
    "name" varchar(150),
    "fields" jsonb NOT NULL,
    "created_by" uuid,
    "last_updated" timestamp
)
