CREATE TABLE menu (
    menu_id SERIAL NOT NULL PRIMARY KEY,
    menu_item VARCHAR NOT NULL,
    menu_parent INTEGER NOT NULL DEFAULT 0
);

INSERT INTO menu (menu_item, menu_parent) 
VALUES 
('Об управлении', 0), ('Социальное обслуживание', 0), ('Социальная поддержка', 0), ('Опека и попечительство', 0),
('Руководство', 1), ('Структура', 1), ('Льготы ЖКУ', 3), ('Субсидии ЖКУ', 3),
('Подразделения', 1), ('Опека над несовершеннолетними', 4), ('Комплексный центр', 2), ('Детские пособия', 2);