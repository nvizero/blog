
CREATE TABLE "users" (
  "id" int PRIMARY KEY,
  "username" varchar,
  "password" varchar NOT NULL,
  "created_at" timestamp NOT NULL
);

CREATE TABLE "cates" (
  "id" int PRIMARY KEY,
  "name" varchar,
  "created_at" timestamp NOT NULL
);

CREATE TABLE "posts" (
  "id" int PRIMARY KEY,
  "cate_id" int,
  "user_id" int,
  "title" varchar,
  "content" text,
  "disable" int NOT NULL,
  "created_at" timestamp NOT NULL
);

ALTER TABLE "posts" ADD FOREIGN KEY ("cate_id") REFERENCES "cates" ("id");

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
