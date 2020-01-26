CREATE TABLE news (
    news_id serial primary key not null,
    news_name varchar not null,
    news_date date not null default current_date,
    news_text text not null
);