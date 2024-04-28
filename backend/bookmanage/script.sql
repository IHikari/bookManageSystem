create table categories
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3) null,
    updated_at datetime(3) null,
    deleted_at datetime(3) null,
    name       varchar(50) not null,
    alias      varchar(50) not null,
    constraint uni_categories_alias
        unique (alias),
    constraint uni_categories_name
        unique (name)
);

create table books
(
    id         bigint unsigned auto_increment
        primary key,
    book_name  varchar(50)     not null,
    image      varchar(255)    null,
    author     varchar(50)     not null,
    cate_id    bigint unsigned null,
    publisher  varchar(50)     null,
    number     bigint unsigned not null,
    remarks    varchar(255)    null,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    constraint books_categories_id_fk
        foreign key (cate_id) references categories (id)
);

create index idx_books_deleted_at
    on books (deleted_at);

create index idx_categories_deleted_at
    on categories (deleted_at);

create table reserves
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    school_id  bigint unsigned null,
    book_id    bigint unsigned null,
    book_num   bigint unsigned null
);

create index idx_reserves_deleted_at
    on reserves (deleted_at);

create table users
(
    id         bigint unsigned auto_increment
        primary key,
    name       longtext             not null,
    nickname   varchar(50)          null,
    password   longtext             not null,
    permission tinyint(1) default 0 null,
    user_pic   varchar(255)         null,
    email      varchar(50)          null,
    created_at datetime(3)          null,
    updated_at datetime(3)          null,
    deleted_at datetime(3)          null
);

create table records
(
    id         bigint unsigned auto_increment
        primary key,
    created_at datetime(3)     null,
    updated_at datetime(3)     null,
    deleted_at datetime(3)     null,
    school_id  bigint unsigned null,
    book_id    bigint unsigned null,
    book_num   bigint unsigned null,
    is_borrow  tinyint(1)      null,
    constraint records_users_id_fk
        foreign key (school_id) references users (id)
);

create index idx_records_deleted_at
    on records (deleted_at);

create index idx_users_deleted_at
    on users (deleted_at);


