CREATE TABLE news (
    news_id SERIAL PRIMARY KEY NOT NULL,
    news_name VARCHAR NOT NULL,
    news_date DATE NOT NULL DEFAULT current_date,
    news_text TEXT NOT NULL
);