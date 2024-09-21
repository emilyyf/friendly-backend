-- +goose Up
-- +goose StatementBegin
create table household (
    id uuid unique primary key default gen_random_uuid(),
	code integer,
	date timestamp with time zone,
	address varchar,
	adress_number varchar,
	adress_complement varchar,
	cep varchar,
	city varchar,
	neighborhood varchar,
	residense_type varchar,
	rent_value integer,
	building_materials varchar,
	in_house_bathroom bool,
	residents integer,
	rooms integer,
	beds integer,
	car bool,
	television bool,
	refrigerator bool,
	microwave bool,
	washing_machine bool,
	road_type varchar,
	referring_person varchar,
	create_log uuid references log(id),
	update_log uuid references log(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table household;
-- +goose StatementEnd
