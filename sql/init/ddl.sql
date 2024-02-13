CREATE TABLE users (
  id            SERIAL PRIMARY KEY,
  balance       BIGINT NOT NULL,
  balance_limit BIGINT NOT NULL
);

CREATE TABLE transactions (
    user_id     INTEGER NOT NULL,
    amount      BIGINT NOT NULL,
    description VARCHAR(10) NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ttype       CHAR(1) NOT NULL
);