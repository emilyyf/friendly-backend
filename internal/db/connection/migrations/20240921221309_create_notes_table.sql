-- +goose Up
-- +goose StatementBegin
create table notes (
    id uuid unique primary key default gen_random_uuid(),
	person_id uuid references person(id),
	child_id uuid references child(id),
	date timestamp with time zone,
	description varchar,
	create_log uuid references log(id),
	update_log uuid references log(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table notes;
-- +goose StatementEnd
