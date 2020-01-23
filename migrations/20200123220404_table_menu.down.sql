CREATE TABLE menu (
    id bigserial not null primary key,
    menuitem varchar not null,
    parent_menu bigserial
);