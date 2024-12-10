create table if not exists person
(
    id serial primary key,
    lname varchar(255),
    fname varchar(255),
    age integer
);

create user :DB_USR with password :'DB_PWD';
grant all privileges on database :DB_NAME to :DB_USR;
grant all privileges on schema public to :DB_USR;
grant usage on schema public to :DB_USR;
grant all on all tables in schema public to :DB_USR;
grant usage, select on all sequences in schema public to :DB_USR;
