CREATE TABLE pages (
    page_id SERIAL NOT NULL PRIMARY KEY,
    page_name VARCHAR NOT NULL,
    page_text TEXT NOT NULL,
    page_menu INT UNIQUE,
    FOREIGN KEY (page_menu) REFERENCES menu (menu_id) ON DELETE NO ACTION
);