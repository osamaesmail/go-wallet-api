BEGIN;

CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TYPE currency AS ENUM ('USD', 'BTC', 'ETC');

CREATE TABLE IF NOT EXISTS accounts
(
    id         UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    user_id    UUID NOT NULL,
    balance    DECIMAL NOT NULL,
    currency   TEXT NOT NULL,
    created_at TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP WITHOUT TIME ZONE NOT NULL,
    deleted_at TIMESTAMP WITHOUT TIME ZONE
);

ALTER TABLE accounts ADD CONSTRAINT unique_account_currency UNIQUE ("user_id", "currency");

CREATE TABLE IF NOT EXISTS transactions
(
    id           UUID DEFAULT gen_random_uuid() PRIMARY KEY,
    amount       DECIMAL NOT NULL,
    from_account UUID NOT NULL REFERENCES accounts (id),
    to_account   UUID NOT NULL REFERENCES accounts (id),
    created_at   TIMESTAMP WITHOUT TIME ZONE NOT NULL DEFAULT NOW()
);

commit;