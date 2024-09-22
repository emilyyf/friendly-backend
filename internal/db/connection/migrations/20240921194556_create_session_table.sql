-- +goose Up
-- +goose StatementBegin
create table session (
	id uuid unique primary key default gen_random_uuid(),
	user_id uuid references users(id),
	token varchar,
	refresh varchar(75),
	type varchar(20),
	is_revoked bool
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table session;
-- +goose StatementEnd
