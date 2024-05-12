CREATE TABLE users 
(
    id serial not null unique,
    name varchar(255) not null,
    username varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE bills
(
    id serial not null unique,
    product varchar(255) not null,
    price float,
    created_at date
);

CREATE TABLE users_bill
(
    id serial not null unique,
    user_id int references users(id) on delete cascade not null,
    bill_id int references bills(id) on delete cascade not null
);