-- Create the enum types
DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'transaction_type') THEN
        CREATE TYPE transaction_type AS ENUM ('income', 'expense');
    END IF;
END$$;

DO $$
BEGIN
    IF NOT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'category_type') THEN
        CREATE TYPE category_type AS ENUM ('income', 'expense');
    END IF;
END$$;

-- Create the "users" table
CREATE TABLE users (
  "id" bigserial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "full_name" varchar NOT NULL,
  "email" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "role" varchar NOT NULL DEFAULT 'user',
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  "updated_at" timestamptz NOT NULL DEFAULT(now()),
  "deleted_at" timestamptz DEFAULT null,
  "user_uuid" uuid NOT NULL UNIQUE -- Add UNIQUE constraint
);

-- Create the "transaction" table
CREATE TABLE transaction (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "amount" bigint NOT NULL,
  "type" transaction_type NOT NULL,
  "description" varchar NOT NULL,
  "category_id" uuid NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  "updated_at" timestamptz NOT NULL DEFAULT(now()),
  "deleted_at" timestamptz DEFAULT null
);

-- Create the "categories" table
CREATE TABLE categories (
  "id" uuid PRIMARY KEY,
  "name" varchar NOT NULL,
  "type" category_type NOT NULL
);

-- Create the "reports" table
CREATE TABLE reports (
  "id" uuid PRIMARY KEY,
  "user_id" uuid NOT NULL,
  "period" varchar NOT NULL,
  "total_income" decimal,
  "total_expense" decimal,
  "created_at" timestamptz NOT NULL DEFAULT(now()),
  "updated_at" timestamptz NOT NULL DEFAULT(now()),
  "deleted_at" timestamptz DEFAULT null
);

-- Create the "sessions" table
CREATE TABLE sessions (
    "id" uuid PRIMARY KEY,
    "user_id" uuid NOT NULL,
    "refresh_token" varchar NOT NULL,
    "expires_at" timestamptz NOT NULL DEFAULT(now()),
    "created_at" timestamptz NOT NULL DEFAULT(now())
);

-- Add foreign keys
ALTER TABLE "transaction" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_uuid");
ALTER TABLE "transaction" ADD FOREIGN KEY ("category_id") REFERENCES "categories" ("id");
ALTER TABLE "reports" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_uuid");
ALTER TABLE "sessions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("user_uuid");