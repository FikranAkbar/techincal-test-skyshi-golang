create table todos
(
    todo_id           bigint unsigned auto_increment
        primary key,
    activity_group_id bigint unsigned null,
    title             varchar(300)    not null,
    priority          longtext        null,
    created_at        datetime(3)     null,
    updated_at        datetime(3)     null,
    deleted_at        datetime(3)     null,
    constraint fk_todos_activity
        foreign key (activity_group_id) references activities (activity_id)
);

create index idx_todos_deleted_at
    on todos (deleted_at);