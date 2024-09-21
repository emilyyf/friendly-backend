-- +goose Up
-- +goose StatementBegin
create table vaccines (
    id uuid unique primary key default gen_random_uuid(),
	name varchar,
	create_log uuid references log(id),
	update_log uuid references log(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table vaccines;
-- +goose StatementEnd
