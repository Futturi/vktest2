CREATE TABLE  users
(
    id BIGSERIAL PRIMARY KEY,
    username text UNIQUE,
    password_hash text
);

CREATE TABLE announcements
(
    id BIGSERIAL PRIMARY KEY,
    name text,
    body text,
    image text,
    price int,
    data int
);

CREATE TABLE users_announcements
(
    user_id int references users(id) on delete cascade,
    announcement_id int references announcements(id) on delete cascade
);
