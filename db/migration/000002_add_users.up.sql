CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "name" varchar(50) NOT NULL,
  "email" varchar(100) UNIQUE NOT NULL,
  "password" varchar(200) NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);