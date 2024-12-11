# Simple webapp on golang

### to start
need set env variables:
```code
DB_ADM_USR, DB_ADM_PWD, DB_NAME, DB_USR, DB_PWD
```
1. up docker with postgres instance:
```console
docker compose up -d
```
2. create table and use in database
```console
make prepare_db
```
3. run app
```console
make run
```
