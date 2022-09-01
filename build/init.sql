create table if not exists historical_data
(
    id       serial primary key,
    currency varchar ( 50 ) not null ,
    rate     double precision,
    date     timestamp with time zone default CURRENT_DATE not null,
    unique (currency, date)
);