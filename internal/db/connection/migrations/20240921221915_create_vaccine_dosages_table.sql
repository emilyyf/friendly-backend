-- +goose Up
-- +goose StatementBegin
create table vaccine_dosages (
	id uuid unique primary key default gen_random_uuid(),
	vaccine_id uuid references vaccines(id),
	description varchar,
	create_log uuid references log(id),
	update_log uuid references log(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table vaccine_dosages;
-- +goose StatementEnd
