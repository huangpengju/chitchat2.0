CREATE DATABASE IF NOT EXISTS chitchat DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

CREATE TABLE `cc_user`(
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户ID',
    `username` varchar(50) unique NOT NULL COMMENT '用户名',
    `password` varchar(100) NOT NULL COMMENT '密码',
    `email` varchar(100) DEFAULT NULL COMMENT '邮箱',
    `phone` varchar(20) DEFAULT NULL COMMENT '电话',
    `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `username_UNIQUE` (`username`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';  

CREATE TABLE `cc_user_info`(
    `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户信息ID',
    `user_id` int(11) NOT NULL COMMENT '用户ID',
    `name` varchar(50) DEFAULT NULL COMMENT '姓名',
    `age` tinyint(3) DEFAULT NULL COMMENT '年龄',
    `gender` tinyint(3) DEFAULT '0' COMMENT '性别：0-未知，1-男，2-女',
    `avatar` varchar(100) DEFAULT NULL COMMENT '头像',
    `address` varchar(100) DEFAULT NULL COMMENT '地址',
    `intro` varchar(500) DEFAULT NULL COMMENT '简介',
    PRIMARY KEY (`id`),
    KEY `FK_user` (`user_id`),
    CONSTRAINT `FK_user` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE 
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 COMMENT='用户信息表';