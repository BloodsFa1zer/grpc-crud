CREATE TABLE IF NOT EXISTS Users (
    ID SERIAL primary key ,
    nick_name   text UNIQUE,
    first_name TEXT not null,
    last_name TEXT not null,
    password text not null,
    created_at text not null,
    updated_at text DEFAULT 'NULL',
    deleted_at text DEFAULT 'NULL'
);
