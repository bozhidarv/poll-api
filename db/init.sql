DROP TABLE IF EXISTS settings;
DROP TABLE IF EXISTS polls;
DROP TABLE IF EXISTS users;

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
    "last_updated" timestamp,

    CONSTRAINT pk_polls_key PRIMARY KEY ("id")
);

CREATE TABLE IF NOT EXISTS users (
    "id"   uuid DEFAULT uuid_generate_v4() NOT NULL,
    "username" varchar(150),
    "email" varchar(150) NOT NULL UNIQUE,
    "password" varchar(200) NOT NULL,
    "last_updated" timestamp,

    CONSTRAINT pk_users_key PRIMARY KEY ("id")
);
