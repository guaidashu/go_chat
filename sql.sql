create table `user`(
    `id` bigint(20) auto_increment not null COMMENT '唯一主键用户id',
    `token` varchar(255) COMMENT '用户SSO登录的token',
    `avatar` varchar(255) COMMENT '用户头像',
    `sex` varchar(2) COMMENT '性别',
    `nickname` varchar(255) COMMENT '昵称',
    `salt` varchar(255),
    `online` int(11),
    `createat` datetime COMMENT '创建时间',
    `stat` int(11) COMMENT '1/可用，0冻结',
    `mobile` varchar(255),
    `passwd` varchar(255) COMMENT '密码',
    `memo` varchar(255) COMMENT '简单一点描述你自己',
    PRIMARY KEY (`id`)
)ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;