-- +goose Up
-- +goose StatementBegin
create table users (
	id uuid unique primary key default gen_random_uuid(),
	name varchar,
	email varchar,
	password varchar,
	role integer,
	created_at timestamp with time zone,
	updated_at timestamp with time zone,
	verified bool
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user;
-- +goose StatementEnd
