-- CREATE SEQUENCE client_id_seq
--     AS integer
--     INCREMENT 1
--     MINVALUE 1
--     MAXVALUE 250000
--     START 1
--     CACHE 1
--     NO CYCLE;

-- CREATE TABLE IF NOT EXISTS client_credential (
--     id integer NOT NULL DEFAULT nextval('client_id_seq'),
--     client_id  VARCHAR(128),
--     public_key VARCHAR
--     id SERIAL PRIMARY KEY,
--     client_id VARCHAR NOT NULL,
--     public_key VARCHAR,
--     private_key VARCHAR
-- )

CREATE TABLE IF NOT EXISTS client_credential
(
    id          SERIAL PRIMARY KEY,
    client_id   VARCHAR NOT NULL,
    username    VARCHAR NOT NULL,
    public_key  VARCHAR NOT NULL,
    private_key VARCHAR NOT NULL
)
