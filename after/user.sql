/* As root user */
CREATE USER 'gin_blogs'@'localhost' IDENTIFIED BY 'tmp_pwd';
CREATE DATABASE gin_blogs;
GRANT ALL PRIVILEGES ON gin_blogs . * TO 'gin_blogs'@'localhost';