-- 'demo' というデータベースを作成
-- 'demo_user' というユーザー名のユーザーを作成
-- データベース 'demo_user' への権限を付与
CREATE DATABASE IF NOT EXISTS demo CHARACTER SET utf8mb4 COLLATE utf8mb4_bin;
GRANT ALL on demo.* TO `demo_user`@`%` identified BY 'demopass';