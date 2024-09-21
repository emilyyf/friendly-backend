-- +goose Up
-- +goose StatementBegin
create table exams (
	id uuid unique primary key default gen_random_uuid(),
	description varchar,
	date timestamp with time zone,
	result varchar,
	igm varchar,
	igg varchar,
	medical_history uuid references medical_history(id),
	create_log uuid references log(id),
	update_log uuid references log(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table exams;
-- +goose StatementEnd
