CREATE DATABASE IF NOT EXISTS `users_api` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

use `users_api`;

CREATE TABLE IF NOT EXISTS `users` (
	`id` int(11) NOT NULL AUTO_INCREMENT,
	`name` varchar(255) NOT NULL,
	`email` varchar(255) NOT NULL,
	`is_active` tinyint(1) NOT NULL DEFAULT '1',
	`birthdate` date NOT NULL,
	`gender` enum('male','female','other') NOT NULL,
	`password` varchar(255) NOT NULL,
	`created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`),
	UNIQUE KEY `email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;