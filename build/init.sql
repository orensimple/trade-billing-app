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

ALTER TABLE billing.accounts
    ADD CONSTRAINT `blocked_amount_value` CHECK (accounts.blocked_amount <= accounts.balance);

ALTER TABLE billing.accounts
    ADD CONSTRAINT `blocked_amount_zero` CHECK (accounts.blocked_amount >= 0);

ALTER TABLE billing.accounts
    ADD CONSTRAINT `balance_zero` CHECK (accounts.balance >= 0);

INSERT INTO billing.accounts (id, currency_code, tariff, name, balance, blocked_amount, created_at, updated_at, deleted_at) VALUES ('d84db1a4-4f6c-4871-9ffa-1b3564c44111', 'RU', null, 'Main', 1000.00, 0.00, '2021-11-04 13:04:37', null, null);