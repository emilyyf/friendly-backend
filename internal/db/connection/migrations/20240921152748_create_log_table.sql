-- +goose Up
-- +goose StatementBegin
create table log (
    id uuid unique primary key default gen_random_uuid(),
	user_id uuid,
	table_name varchar,
	date timestamp with time zone,
	description varchar,
	action varchar,
	row_id uuid
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table log;
-- +goose StatementEnd
