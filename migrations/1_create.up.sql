create table resource_item
(
    id SERIAL PRIMARY KEY ,
    title TEXT NOT NULL,
    tag VARCHAR(20) NOT NULL,
    created_at TIMESTAMP,
    url TEXT NOT NULL,
    updated_at TIMESTAMP
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
