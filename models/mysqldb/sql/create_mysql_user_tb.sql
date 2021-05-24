create user 'jim' IDENTIFIED by 'password';

create database `markdb`;
grant all privileges on markdb.* to 'jim';
