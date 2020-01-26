CREATE TABLE menu (
    id serial not null primary key,
    menuitem varchar not null,
    parent_menu integer REFERENCES menu (id)
);