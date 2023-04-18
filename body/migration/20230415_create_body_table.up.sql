create table body (
	id serial PRIMARY KEY,
	user_id int not null default 0,
	weight NUMERIC(5,2) not null default 0,
	muscle NUMERIC(5,2) not null default 0,
	fat_rate NUMERIC(5,2) not null default 0,
	created_at timestamp not null default CURRENT_TIMESTAMP,
	updated_at timestamp not null default CURRENT_TIMESTAMP
);
