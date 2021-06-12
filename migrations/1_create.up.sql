create table resource_item
(
    id SERIAL PRIMARY KEY ,
    title TEXT NOT NULL,
    tag VARCHAR(20) NOT NULL,
    created_at TIMESTAMP,
    url TEXT NOT NULL,
    updated_at TIMESTAMP,
    SITENAME TEXT,
    IMAGE TEXT,
    EXCERPT TEXT,
    AUTHOR TEXT

);
create table tag
(
    tid SERIAL PRIMARY KEY ,
    name VARCHAR(20) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    parent SERIAL,
    FOREIGN KEY (parent)  REFERENCES tag (tid)    ON DELETE CASCADE

);


create table user_info
(
    uid serial primary key ,
    name varchar(30),
    email varchar(50)
);
