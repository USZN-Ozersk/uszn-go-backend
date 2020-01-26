CREATE TABLE menu (
    id serial not null primary key,
    menuitem varchar not null,
    parent_id integer foreign key unique REFERENCES menu (id)
);