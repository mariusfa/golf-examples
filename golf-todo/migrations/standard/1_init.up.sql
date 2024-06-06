CREATE SCHEMA IF NOT EXISTS todoschema;

CREATE USER {{.User}} WITH PASSWORD '{{.Password}}';

ALTER USER {{.User}} SET search_path TO todoschema;

GRANT USAGE ON SCHEMA todoschema TO {{.User}};

CREATE TABLE IF NOT EXISTS todoschema.todos (
    id UUID PRIMARY KEY NOT NULL,
    title VARCHAR NOT NULL
);

GRANT SELECT, INSERT, UPDATE, DELETE ON todoschema.todos TO {{.User}};
