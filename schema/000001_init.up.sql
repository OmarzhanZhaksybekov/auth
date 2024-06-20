CREATE TABLE users (
    id serial NOT NULL,
    email varchar(255) NOT NULL,
    phone varchar(20),
    password varchar(255) NOT NULL,
    role varchar(50) DEFAULT 'user'::character varying
);
