CREATE TABLE "users" (
  "id" bigserial PRIMARY KEY,
  "username" varchar NOT NULL,
  "role" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE TABLE "posts" (
  "id" bigserial PRIMARY KEY,
  "title" varchar NOT NULL,
  "body" text NOT NULL,
  "user_id" bigint NOT NULL,
  "status" varchar NOT NULL,
  "created_at" timestamptz DEFAULT (now())
);

CREATE INDEX ON "posts" ("user_id");

CREATE INDEX ON "posts" ("title");

CREATE INDEX ON "posts" ("title", "user_id");

COMMENT ON COLUMN "posts"."body" IS 'Content of the post';

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
