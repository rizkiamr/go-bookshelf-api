CREATE TABLE "authors" (
  "id" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "insertedAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "publishers" (
  "id" varchar PRIMARY KEY,
  "name" varchar NOT NULL,
  "insertedAt" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "books" (
  "id" bigserial PRIMARY KEY,
  "name" varchar NOT NULL,
  "year" integer NOT NULL,
  "author_id" varchar NOT NULL,
  "summary" text NOT NULL,
  "publisher_id" varchar NOT NULL,
  "pageCount" integer NOT NULL,
  "readPage" integer NOT NULL,
  "finished" bool NOT NULL,
  "reading" bool NOT NULL,
  "insertedAt" timestamptz NOT NULL DEFAULT (now()),
  "updatedAt" timestamptz
);

ALTER TABLE "books" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "books" ADD FOREIGN KEY ("publisher_id") REFERENCES "publishers" ("id");
