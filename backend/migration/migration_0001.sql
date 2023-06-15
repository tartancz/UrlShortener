CREATE TABLE redirects (
  "id" integer PRIMARY KEY,
  "url" varchar,
  "shorten_url" varchar,
  "created_at" timestamp DEFAULT (now())
);
