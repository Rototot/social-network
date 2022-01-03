create table users
(
    id         bigint auto_increment,
    email      VARCHAR(512)                                                                  not null,
    password   varchar(512)                                                                  not null,
    first_name varchar(1024)                                                                 null,
    last_name  varchar(1024)                                                                 null,
    gender     ENUM ('male', 'female') default 'male'                                        not null,
    age        tinyint                 default 0                                             not null,
    interests  JSON                                                                          null ,
    city       varchar(512)                                                                  not null,
    created_at timestamp               default CURRENT_TIMESTAMP                             not null,
    updated_at timestamp               DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP not null,
    constraint users_pk
        primary key (id)
);

create unique index users_email_uindex
    on users (email);
