-- +goose Up
-- +goose StatementBegin
create table child (
    id uuid unique primary key default gen_random_uuid(),
	id_household uuid references household(id),
	id_mother uuid references person(id),
	code integer,
	name varchar,
	birth timestamp with time zone,
	age varchar,
	local varchar,
	race varchar,
	alive_birth_certificate varchar,
	birth_certificate varchar,
	rg varchar,
	cpf varchar,
	sus_card varchar,
	blood_type varchar,
	weight_at_birth varchar,
	height_at_birth varchar,
	first_apgar varchar,
	fifth_apgar varchar,
	neonatal_heel_prick timestamp with time zone,
	hear_test timestamp with time zone,
	hearing_triage timestamp with time zone,
	eye_test timestamp with time zone,
	od varchar,
	oe varchar,
	pregnacy_time varchar,
	login varchar,
	msd varchar,
	mmii varchar,
	create_log uuid references log(id),
	update_log uuid references log(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table child;
-- +goose StatementEnd
