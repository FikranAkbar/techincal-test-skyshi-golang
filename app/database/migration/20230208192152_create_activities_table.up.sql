create table activities
(
    activity_id bigint unsigned auto_increment
        primary key,
    title       varchar(300) not null,
    email       varchar(300) not null,
    created_at  datetime(3)  null,
    updated_at  datetime(3)  null,
    deleted_at  datetime(3)  null,
    constraint idx_email
        unique (email)
);

create index idx_activities_deleted_at
    on activities (deleted_at);
