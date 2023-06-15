CREATE TABLE redirects (
  "id" SERIAL PRIMARY KEY,
  "url" varchar,
  "shorten_url" varchar,
  "created_at" timestamp DEFAULT (now())
);

CREATE UNIQUE INDEX redirects_shorten_url on redirects (shorten_url);