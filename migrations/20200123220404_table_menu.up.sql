CREATE TABLE menu (
    menu_id SERIAL NOT NULL PRIMARY KEY,
    menu_item VARCHAR NOT NULL,
    menu_parent INTEGER NOT NULL DEFAULT 0,
    custom_link BOOLEAN NOT NULL DEFAULT false,
    custom_link_value VARCHAR NOT NULL DEFAULT ''
);

INSERT INTO menu (menu_item, menu_parent) 
VALUES 
('Об управлении', 0), ('Социальное обслуживание', 0), ('Социальная поддержка', 0), ('Опека и попечительство', 0),
('Прочее', 0);