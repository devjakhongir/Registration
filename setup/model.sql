

drop table if exists registration cascade;
-- drop table if exists token cascade;

create table registration (
    registration_id serial not null primary key,
    username varchar(64) not null,
    email varchar(120) not null,
    password varchar(100) not null
);

-- create table token (
--     token_id serial not null primary key,
--     auth varchar(140) not null,
--     created_at timestamp with time zone default current_timestamp,
--     registration_id int not null references registration(registration_id)
-- );