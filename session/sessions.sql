CREATE TABLE sessions (
  token TEXT PRIMARY KEY,
  data BYTEA NOT NULL,
  expiry TIMESTAMP NOT NULL
);
CREATE INDEX sessions_expiry_idx ON sessions (expiry);

