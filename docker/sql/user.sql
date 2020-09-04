CREATE DATABASE IF NOT EXISTS aio_db DEFAULT CHARACTER SET = utf8mb4;

CREATE USER 'aio_db_user'@'%' IDENTIFIED BY 'RjwuGv12XSDaVsW7';

GRANT ALL ON aio_db.* TO 'aio_db_user'@'%';

flush privileges;

