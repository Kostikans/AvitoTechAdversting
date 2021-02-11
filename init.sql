CREATE TABLE advertising(
    advertising_id serial primary key NOT NULL ,
    name varchar(200),
    description varchar(1000),
    mainPhoto text,
    photos text[],
    cost int,
    created timestamp default now()
);

CREATE TABLE advertising_count(
    partition_id serial primary key NOT NULL ,
    count int
);

INSERT INTO advertising_count VALUES (default,0);
