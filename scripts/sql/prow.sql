CREATE TABLE `prow_notice` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `callback_url` varchar(128) DEFAULT NULL,
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;


CREATE TABLE `prow_owner` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `owner_name` varchar(32) DEFAULT NULL COMMENT 'gitlab user name',
  `project_name` varchar(32) DEFAULT NULL,
  `path_name` varchar(64) DEFAULT NULL,
  `phone` varchar(24) DEFAULT NULL COMMENT '电话',
  `email` varchar(255) DEFAULT NULL,
  `description` mediumtext,
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

CREATE TABLE `prow_path` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `path_name` varchar(64) DEFAULT NULL,
  `project_name` varchar(32) DEFAULT NULL,
  `description` mediumtext,
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `prow_project` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `project_name` varchar(32) DEFAULT NULL,
  `description` mediumtext,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `prow_robot` (
  `id` int(32) NOT NULL AUTO_INCREMENT,
  `robot_name` varchar(32) DEFAULT NULL,
  `description` mediumtext,
  `gitlab_url` varchar(128) DEFAULT NULL,
  `gitlab_token` varchar(128) DEFAULT NULL,
  `create_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `update_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;