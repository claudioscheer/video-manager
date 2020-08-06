DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` varchar(36) NOT NULL,
  `name` varchar(100) NOT NULL,
  `description` varchar(500) DEFAULT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` varchar(36) NOT NULL,
  `name` varchar(100) NOT NULL,
  `email` varchar(100) NOT NULL,
  `password` varchar(1000) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role` (
  `user` varchar(36) NOT NULL,
  `role` varchar(36) NOT NULL,
  PRIMARY KEY (`user`,`role`),
  KEY `user_role_FK_1` (`role`),
  CONSTRAINT `user_role_FK` FOREIGN KEY (`user`) REFERENCES `user` (`id`),
  CONSTRAINT `user_role_FK_1` FOREIGN KEY (`role`) REFERENCES `role` (`id`)
);