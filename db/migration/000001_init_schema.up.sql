CREATE TABLE "users" (
  "id"                    bigserial PRIMARY KEY,
  "email"                 varchar NOT NULL,
  "name"                  varchar NOT NULL,
  "username"              varchar NOT NULL,
  "password"              varchar NOT NULL,
  "password_changed_at"   timestamptz NOT NULL DEFAULT('0001-01-01 00:00:00Z'),  
  "phone_number"          bigint NOT NULL,
  "device_token"          text NOT NULL,
  "lang"                  varchar NOT NULL,
  "avatar"                varchar NOT NULL,
  "user_level"            varchar NOT NULL,
  "is_active"             boolean NOT NULL,
  "created_at"            timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "promos" (
  "id"                bigserial PRIMARY KEY,
  "promo_name"        varchar NOT NULL,
  "store_id"          bigint NOT NULL,
  "promo_code"        varchar NOT NULL,
  "promo_description" text NOT NULL, 
  "quantity"          bigint NOT NULL,
  "start_at"          bigint NOT NULL,
  "expired_at"        bigint NOT NULL,
  "is_active"         boolean NOT NULL,
  "created_by"        bigint NOT NULL,
  "created_at"        timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "stores" (
  "id"              bigserial PRIMARY KEY,
  "name"            varchar NOT NULL,
  "address"         varchar NOT NULL,
  "description"     text NOT NULL,
  "phone_number"    bigint NOT NULL,
  "operational_id"  bigint NOT NULL,
  "is_active"       boolean NOT NULL,
  "created_at"      timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "operational_time" (
  "id"                bigserial PRIMARY KEY,
  "opening_time"      varchar NOT NULL,
  "closing_time"      varchar NOT NULL,
  "operational_days"  text NOT NULL,
  "off_days"          varchar NOT NULL,
  "is_active"         boolean NOT NULL,
  "created_at"        timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id"              uuid PRIMARY KEY,
  "user_id"         bigint NOT NULL,
  "username"        varchar NOT NULL,
  "refresh_token"   varchar NOT NULL,
  "user_agent"      varchar NOT NULL,
  "client_ip"       varchar NOT NULL,
  "is_blocked"      boolean NOT NULL DEFAULT false,
  "expires_at"      timestamptz NOT NULL,
  "created_at"      timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "promos" ADD FOREIGN KEY ("store_id") REFERENCES "stores" ("id");

ALTER TABLE "promos" ADD FOREIGN KEY ("created_by") REFERENCES "users" ("id");

ALTER TABLE "stores" ADD FOREIGN KEY ("operational_id") REFERENCES "operational_time" ("id");