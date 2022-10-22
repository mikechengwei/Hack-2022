drop table `hackathon_datasource`;
CREATE TABLE `hackathon_datasource`
(
    `name`             varchar(255) NOT NULL DEFAULT '' COMMENT '名称',
    `host`             varchar(255) NOT NULL DEFAULT '' COMMENT '地址',
    `port`             varchar(255) NOT NULL DEFAULT '' COMMENT '端口',
    `username`         varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
    `password`         varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
    `statusPort`       varchar(255) NOT NULL DEFAULT '' COMMENT '状态端口',
    `pdAddress`        varchar(255) NOT NULL DEFAULT '' COMMENT 'pd地址',
    `service_name`     varchar(255) NOT NULL DEFAULT '' COMMENT 'oracle service name',
    `type`             int(4) NOT NULL DEFAULT 0 COMMENT '数据库源类型',
    `id`               bigint(21) NOT NULL AUTO_INCREMENT COMMENT '应用自增id',
    `created_at`       datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `created_by`       varchar(100) NOT NULL DEFAULT '' COMMENT '创建人',
    `last_modified_by` varchar(100) NOT NULL DEFAULT '' COMMENT '修改人',
    `maintainer`       varchar(100) NOT NULL DEFAULT '' COMMENT '维护人',
    `updated_at`       datetime null COMMENT '修改时间',
    `deleted_at`       datetime COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY                `name` (`name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT '数据源表';

drop table `hackathon_task`;

CREATE TABLE `hackathon_task`
(
    `name`               varchar(255) NOT NULL DEFAULT '' COMMENT '名称',
    `source_client`      varchar(255) NOT NULL DEFAULT '' COMMENT '客户端',
    `source_datasource`  bigint(21) NOT NULL DEFAULT 0 COMMENT '数据源id',
    `source_database`    varchar(255) NOT NULL DEFAULT '' COMMENT '源数据库',
    `source_tables`      varchar(255) NOT NULL DEFAULT '' COMMENT '源表列表',
    `source_sql`         varchar(255) NOT NULL DEFAULT '' COMMENT '源sql',
    `source_split_mode`  varchar(255) NOT NULL DEFAULT '' COMMENT '任务切分模式',
    `target_datasource`  bigint(21) NOT NULL DEFAULT 0 COMMENT '数据源',
    `target_database`    varchar(255) NOT NULL DEFAULT '' COMMENT '数据库',
    `target_import_mode` varchar(255) NOT NULL DEFAULT '' COMMENT '导入模式',
    `concurrent`         int(4) NOT NULL DEFAULT 0 COMMENT '并发度',
    `sync_schema`        int(4) NOT NULL DEFAULT 0 COMMENT '是否同步schema',
    `status`             int(4) NOT NULL DEFAULT 0 COMMENT '任务状态',
    `id`                 bigint(21) NOT NULL AUTO_INCREMENT COMMENT '应用自增id',
    `created_at`         datetime     NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `created_by`         varchar(100) NOT NULL DEFAULT '' COMMENT '创建人',
    `last_modified_by`   varchar(100) NOT NULL DEFAULT '' COMMENT '修改人',
    `maintainer`         varchar(100) NOT NULL DEFAULT '' COMMENT '维护人',
    `updated_at`         datetime null COMMENT '修改时间',
    `finish_at`          datetime null COMMENT '完成时间',
    `deleted_at`         datetime COMMENT '删除时间',
    PRIMARY KEY (`id`),
    KEY                  `name` (`name`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT '任务表';


create table marvin
(
    ID            bigint(21) NOT NULL AUTO_INCREMENT COMMENT '应用自增id',
    INC_DATETIME  datetime      NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    RANDOM_ID     bigint(21) not null DEFAULT 0 COMMENT '随机id',
    RANDOM_STRING varchar(1000) NOT NULL DEFAULT '' COMMENT '随机字符串',
    PRIMARY KEY (`ID`)
);