-- +goose Up
-- +goose StatementBegin
create table person_vaccine (
	id uuid unique primary key default gen_random_uuid(),
	person_id uuid references person(id),
	vaccine_id uuid references vaccines(id),
	dosage_id uuid references vaccine_dosages(id),
	date timestamp with time zone,
	id_medical_histoy uuid references medical_history(id),
	create_log uuid references log(id),
	update_log uuid references log(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table person_vaccine;
-- +goose StatementEnd
