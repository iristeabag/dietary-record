create table eat (
	id serial PRIMARY KEY,
	food_id int not null,
	name varchar not null,
	amount NUMERIC(5,2) not null default 0,
	unit varchar(5) not null,
	carb NUMERIC(5,2) not null default 0,
	portein NUMERIC(5,2) not null default 0,
	fat NUMERIC(5,2) not null default 0,
	cal NUMERIC(5,2) not null default 0,
	created_at timestamp not null default CURRENT_TIMESTAMP,
	updated_at timestamp not null default CURRENT_TIMESTAMP
);
CREATE TRIGGER set_timestamp BEFORE UPDATE ON eat FOR EACH ROW EXECUTE PROCEDURE trigger_set_timestamp();
