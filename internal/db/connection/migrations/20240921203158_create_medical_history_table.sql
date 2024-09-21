-- +goose Up
-- +goose StatementBegin
create table medical_history (
	id uuid unique primary key default gen_random_uuid(),
	id_person uuid references person(id),
	smoker bool,
	alcohool bool,
	aborts integer,
	vaginal_delivery integer,
	caesarian integer,
	pregnancy varchar,
	blood_type varchar,
	blood_glucose varchar,
	syphilis varchar,
	hiv varchar,
	create_log uuid references log(id),
	update_log uuid references log(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table medical_history;
-- +goose StatementEnd
