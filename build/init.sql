CREATE DATABASE billing;

USE billing;

create table accounts
(
    id             varchar(36)                              not null,
    currency_code  varchar(36)                              not null,
    tariff_id      varchar(36)                              not null,
    name           varchar(256)                             not null,
    balance        decimal(10, 2) default 0.00              not null comment 'сurrent balance',
    blocked_amount decimal(10, 2) default 0.00              not null comment 'blocked amount balance',
    created_at     timestamp      default CURRENT_TIMESTAMP not null,
    updated_at     timestamp                                null,
    deleted_at     timestamp                                null,
    constraint accounts_id_uindex
        unique (id)
)
    comment 'User billing table';

alter table accounts
    add primary key (id);

