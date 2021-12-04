create table logs
(
    Entity    varchar(100) not null,
    Action    varchar(100) not null,
    UserID    int          not null,
    Timestamp timestamp    not null
);