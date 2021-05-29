
# create an user `jim` with `markdb`
docker exec -i mysql mysql -u root -psecret < ./scripts/create_mysql_user_tb.sql