CREATE TABLE IF NOT EXISTS users (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "username" VARCHAR UNIQUE NOT NULL,
    "color" VARCHAR NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);
