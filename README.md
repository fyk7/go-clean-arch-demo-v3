# go-clean-arch-demo-v3

## ADOP pattern inplementation

### migration commands
```sh
$ mysql.server stop # Stop mysql server which instlled to mac directly to use database in docker.
$ brew install golang-migrate
$ mkdir migrations & cd migrations
$ migrate create -ext sql -dir migrations -seq create_users_table 
$ export MYSQL_URL='mysql://gosqldemouser:gosqldemopass@tcp(127.0.0.1:3306)/gosqldemodb?charset=utf8mb4' # change user name and password correctly
$ migrate -database ${MYSQL_URL} -path migrations up
```