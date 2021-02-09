CREATE TABLE advertising(
    advertising_id serial primary key NOT NULL ,
    name varchar(1000),
    photos text[],
    cost int
);
