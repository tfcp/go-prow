create DATABASE IF NOT EXISTS `prow` charset utf8mb4;
use `prow`;
CREATE TABLE `prow_owner` (
  `id` int(10) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '' COMMENT 'owner用户名称',
  `path` varchar(32) NOT NULL DEFAULT '' COMMENT '地址',
  `email` varchar(32) NOT NULL DEFAULT '' COMMENT 'email',
  `phone` varchar(32) NOT NULL DEFAULT '' COMMENT '电话',
  `type` varchar(32) NOT NULL DEFAULT '' COMMENT 'service: 底层微服务 interface: 业务聚合层',
  `create_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
