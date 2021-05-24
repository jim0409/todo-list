
# create an user `jim` with `markdb`
docker exec -i mysql mysql -u root -psecret < ./models/mysqldb/sql/create_mysql_user_tb.sql