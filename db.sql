CREATE TABLE `user` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`first_name` VARCHAR(50) NULL DEFAULT NULL,
	`last_name` VARCHAR(50) NULL DEFAULT NULL,
	`email` VARCHAR(50) NULL DEFAULT NULL,
	`date_created` VARCHAR(50) NULL DEFAULT NULL,
	PRIMARY KEY (`id`),
	UNIQUE INDEX `email_UNIQUE` (`email`)
)
COLLATE='utf8mb4_0900_ai_ci'
ENGINE=InnoDB
;