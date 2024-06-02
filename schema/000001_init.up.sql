CREATE TABLE users (
    id bigserial not null unique,
    name varchar(255) not null,
    username varchar(255) not null,
    password varchar(255) not null
);

CREATE TABLE todo_lists (
    id bigserial not null unique,
    title varchar(255) not null,
    description varchar(255) not null
);

CREATE TABLE todo_items (
    id bigserial not null unique,
    title varchar(255) not null,
    description varchar(255) not null,
    done boolean not null default false
);

CREATE TABLE users_lists (
    id bigserial not null unique,
    user_id bigint not null references users(id) on delete cascade,
    list_id bigint not null references todo_lists(id) on delete cascade
);

CREATE TABLE lists_items (
    id bigserial not null unique,
    item_id bigint not null references todo_items(id) on delete cascade,
    list_id bigint not null references todo_lists(id) on delete cascade
);
