CREATE TABLE "authors" (
  "id" bigInt PRIMARY KEY,
  "name" varchar,
  "insertedAt" timestamptz
);

CREATE TABLE "publishers" (
  "id" bigInt PRIMARY KEY,
  "name" varchar,
  "insertedAt" timestamptz
);

CREATE TABLE "books" (
  "id" bigInt PRIMARY KEY,
  "name" varchar,
  "year" integer,
  "author_id" bigInt,
  "summary" text,
  "publisher_id" bigInt,
  "pageCount" integer,
  "readPage" integer,
  "finished" bool,
  "reading" bool,
  "insertedAt" timestamptz,
  "updatedAt" timestamptz
);

ALTER TABLE "books" ADD FOREIGN KEY ("author_id") REFERENCES "authors" ("id");

ALTER TABLE "books" ADD FOREIGN KEY ("publisher_id") REFERENCES "publishers" ("id");
