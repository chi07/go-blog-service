CREATE TABLE `articles`
(
    `id`               INT(255)     NOT NULL AUTO_INCREMENT,
    `title`            VARCHAR(255) NOT NULL,
    `description`      VARCHAR(255) NOT NULL,
    `content`          TEXT         NOT NULL,
    `meta_keywords`    VARCHAR(255) NULL,
    `meta_description` VARCHAR(255) NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (`id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8;