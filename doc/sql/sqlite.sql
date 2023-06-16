create table `convert_config`
(
--         comment '主健',
    `id`         INTEGER PRIMARY KEY AUTOINCREMENT,

--         comment '输入kv参数配置',
    `in_args`    text      null     default null,
--         comment '输出kv参数配置',
    `out_args`   text      null     default null,

--         comment '创建时间',
    `created_at` timestamp not null default current_timestamp,
--         comment '更新时间',
    `updated_at` timestamp not null default current_timestamp,
    `deleted_at` timestamp null     default null
);
--     comment ='转码参数配置表';


-- default charset = utf8 comment ='转码任务表';

create table `convert_job`
(
--         comment '主健',
    `id`         INTEGER PRIMARY KEY AUTOINCREMENT,

--         comment '配置id',
    `convert_id` int(11)       not null default 0,
--         comment '输入路径',
    `in_put`     varchar(1024) not null default '',
--         comment '输出路径',
    `out_put`    varchar(1024) not null default '',

--         comment '任务状态 0等待 1进行中 2成功 3失败',
    `status`     int(2)        not null default 0,
--         comment '工作机器的Host',
    `host`       varchar(64)   null     default null,
--         comment '工作机器的Ip',
    `ip`         varchar(64)   null     default null,

--         comment '创建时间',
    `created_at` timestamp     not null default current_timestamp,
--         comment '更新时间',
    `updated_at` timestamp     not null default current_timestamp,
    `deleted_at` timestamp     null     default null

);
