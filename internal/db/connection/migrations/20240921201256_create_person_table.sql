-- +goose Up
-- +goose StatementBegin
create table person (
	id uuid unique primary key default gen_random_uuid(),
	id_household uuid references household(id),
	birth timestamp with time zone,
	age varchar,
	name varchar,
	country varchar,
	uf varchar,
	phone varchar,
	graduation varchar,
	rg varchar,
	rg_exp timestamp with time zone,
	cpf varchar,
	sus_card varchar,
	card_serie varchar,
	card_uf varchar,
	compnay varchar,
	work_function varchar,
	hiring_date timestamp with time zone,
	resignation_date timestamp with time zone,
	salary integer,
	extra_income integer,
	employment_card varchar,
	create_log uuid references log(id),
	update_log uuid references log(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table person;
-- +goose StatementEnd
