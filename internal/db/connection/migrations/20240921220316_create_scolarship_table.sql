-- +goose Up
-- +goose StatementBegin
create table scholarship (
    id uuid unique primary key default gen_random_uuid(),
	child_id uuid references child(id),
	grade varchar,
	school varchar,
	year integer,
	period integer,
	create_log uuid references log(id),
	update_log uuid references log(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table scholarship;
-- +goose StatementEnd
