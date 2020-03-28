create table resource_item
(
    id VARCHAR(20) PRIMARY KEY ,
    title TEXT NOT NULL,
    tag VARCHAR(20) NOT NULL,
    created_at TIMESTAMP,
    url TEXT NOT NULL,
    updated_at TIMESTAMP
);
create table tag
(
    tid VARCHAR(20) PRIMARY KEY ,
    name VARCHAR(20) NOT NULL,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    parent VARCHAR(20),
    FOREIGN KEY (parent)  REFERENCES tag (tid)    ON DELETE CASCADE

);