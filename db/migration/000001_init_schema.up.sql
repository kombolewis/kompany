CREATE TABLE "company" (
  "id" bigserial PRIMARY KEY,
  "name" char(15) UNIQUE NOT NULL,
  "description" char(3000),
  "amount" int NOT NULL,
  "registered" boolean NOT NULL DEFAULT FALSE,
  "type" varchar(25) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "company" ("name");
