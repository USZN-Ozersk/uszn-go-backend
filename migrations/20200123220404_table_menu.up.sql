CREATE TABLE menu (
    menu_id serial not null primary key,
    menu_item varchar not null,
    menu_parent integer REFERENCES menu (menu_id) ON DELETE cascade
);