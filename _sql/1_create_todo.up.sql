CREATE TABLE IF NOT EXISTS `demo`.`todos` (
    `id` INT NOT NULL COMMENT 'ID',
    `task` VARCHAR(64) NOT NULL COMMENT 'task',
    `limitDate` VARCHAR(64) NOT NULL COMMENT 'limitData',
    `status` BOOLEAN NOT NULL COMMENT 'status',
    `deleted` BOOLEAN NOT NULL DEFAULT false COMMENT 'deleted',
    PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'Todos';