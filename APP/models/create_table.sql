CREATE DATABASE IF NOT EXISTS app;

USE app;

CREATE TABLE `user` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `username` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `password` varchar(64) COLLATE utf8mb4_general_ci NOT NULL,
    `phone` varchar(64) COLLATE utf8mb4_general_ci NOT NULL UNIQUE,
    `gender` tinyint(4) NOT NULL DEFAULT '0',
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE
    CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE `code`(
    `id` INT AUTO_INCREMENT,
    `code` INT NOT NULL,
    `phone` varchar(255) NOT NULL,
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

CREATE TABLE `bill` (
    `id` bigint(20) NOT NULL AUTO_INCREMENT,
    `user_id` bigint(20) NOT NULL,
    `bill_id` bigint(20) NOT NULL,
    `income_or_outcome` varchar(64) NOT NULL,
    `money` bigint(20) NOT NULL,
    `bill_type` varchar(64) NOT NULL,
    `bill_info` varchar(256),
    `bill_year` varchar(64) NOT NULL,
    `bill_month` varchar(64) NOT NULL,
    `bill_isoweek` varchar(64) NOT NULL,
    `bill_date` varchar(64) NOT NULL,
    `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;