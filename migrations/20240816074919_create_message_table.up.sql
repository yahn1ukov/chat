CREATE TABLE IF NOT EXISTS messages (
    "id" UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    "user_id" UUID NOT NULL REFERENCES users("id") ON DELETE CASCADE,
    "text" VARCHAR NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT now()
);
