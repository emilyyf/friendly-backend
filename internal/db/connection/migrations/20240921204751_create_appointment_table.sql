-- +goose Up
-- +goose StatementBegin
create table appointment (
    id uuid unique primary key default gen_random_uuid(),
	date timestamp with time zone,
	ig varchar,
	weight varchar,
	pa varchar,
	au varchar,
	bcf varchar,
	id_child uuid references child(id),
	create_log uuid references log(id),
	update_log uuid references log(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table appointment;
-- +goose StatementEnd
