CREATE TABLE users_go(
    id serial PRIMARY KEY,
    uuid varchar(64) not null unique,
    nickname VARCHAR(255) not null,
    email VARCHAR(255) not null,
    password VARCHAR(255) not null,
    created_at TIMESTAMP not null
);

CREATE TABLE sessions_go (
  id         serial primary key,
  uuid       varchar(64) not null unique,
  email      VARCHAR(255),
  user_id    integer REFERENCES users(id),
  created_at timestamp not null
);

