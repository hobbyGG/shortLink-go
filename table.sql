-- Active: 1744292313152@@127.0.0.1@13306@slDB
use slDB;

create table `sequence` (
    `id` BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    `stub` VARCHAR(1) not NULL DEFAULT 'a',
    `timestamp` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_unique_stub` (`stub`)
)ENGINE=MyISAM;

create table `short_url_map`(
    `id` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    `create_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `create_by` VARCHAR(32) NOT NULL DEFAULT '',
    `is_del` TINYINT UNSIGNED NOT NULL DEFAULT '0',

    `lurl` VARCHAR(160) DEFAULT NULL,
    `lurl_md5` VARCHAR(64) DEFAULT NULL,
    `surl` VARCHAR(11) DEFAULT NULL,
    PRIMARY KEY (`id`),
    index (`is_del`),
    UNIQUE (`lurl`),
    UNIQUE (`surl`)
)ENGINE=MyISAM;