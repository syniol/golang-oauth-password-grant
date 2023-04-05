CREATE TABLE IF NOT EXISTS client_credential
(
    id      BIGSERIAL,
    data    jsonb NOT NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS client_credential_username_idx ON client_credential( (data->>'username') );
CREATE UNIQUE INDEX IF NOT EXISTS client_credential_client_id_idx ON client_credential( (data->>'clientId') );
