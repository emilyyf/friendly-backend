-- +goose Up
-- +goose StatementBegin
create table ultrassonography (
	id uuid unique primary key default gen_random_uuid(),
	date timestamp with time zone,
	weight varchar,
	height varchar,
	percentage varchar,
	bcf varchar,
	ila varchar,
	liq_am varchar,
	placenta varchar,
	degree varchar,
	id_child uuid references child(id),
	id_medical_history uuid references medical_history(id),
	create_log uuid references log(id),
	update_log uuid references log(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table ultrassonography;
-- +goose StatementEnd
