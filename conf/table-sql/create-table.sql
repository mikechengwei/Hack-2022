drop table `hackathon_datasource`;
CREATE TABLE `hackathon_datasource`
(
    `name`             varchar(255) NOT NULL DEFAULT '' COMMENT '名称',
    `host`             varchar(255) NOT NULL DEFAULT '' COMMENT '地址',
    `port`             varchar(255) NOT NULL DEFAULT '' COMMENT '端口',
    `username`         varchar(255) NOT NULL DEFAULT '' COMMENT '用户名',
    `password`         varchar(255) NOT NULL DEFAULT '' COMMENT '密码',
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