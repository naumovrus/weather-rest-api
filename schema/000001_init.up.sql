CREATE TABLE users 
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE cities
(
    id serial not null unique,
    name varchar(255) not null unique 
);

CREATE TABLE users_city
(
    id serial not null unique,
    user_id int references users (id) on delete cascade not null,
    city_id int references cities (id) on delete cascade not null 
);

