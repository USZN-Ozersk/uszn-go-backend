CREATE TABLE pages (
    page_id SERIAL NOT NULL PRIMARY KEY,
    page_name VARCHAR NOT NULL,
    page_text TEXT NOT NULL,
    page_menu INTEGER UNIQUE REFERENCES menu (menu_id)
);