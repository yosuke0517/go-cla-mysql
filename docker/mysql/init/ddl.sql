-- 'demo' というデータベースを作成
-- 'demo_user' というユーザー名のユーザーを作成
-- データベース 'demo_user' への権限を付与
CREATE DATABASE IF NOT EXISTS demo CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
GRANT ALL on demo.* TO `demo_user`@`%` identified BY 'demopass';

-- -----------------------------------------------------
-- Table `demo`.`todo`
-- -----------------------------------------------------
CREATE TABLE IF NOT EXISTS `demo`.`todo` (
   `id` INT NOT NULL COMMENT 'ID',
   `task` VARCHAR(64) NOT NULL COMMENT 'task',
   `limitDate` VARCHAR(64) NOT NULL COMMENT 'limitData',
   `status` BOOLEAN NOT NULL COMMENT 'status',
   PRIMARY KEY (`id`))
ENGINE = InnoDB
COMMENT = 'Todo';