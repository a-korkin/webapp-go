# Simple webapp on golang

## to start
1. need set env variables:
```code
DB_ADM_USR, DB_ADM_PWD, DB_NAME, DB_USR, DB_PWD
```
2. up docker with postgres instance:
```console
docker compose up -d
```
3. create table and use in database
```console
make prepare_db
```
4. run app
```console
make run
```

## example requests
```http
GET localhost:8080/person

GET localhost:8080/person/1

POST localhost:8080/person 
Content-Type: application/json

{
    "lname": "Travis",
    "fname": "Bickle",
    "age": 23,
}

DELETE localhost:8080/person/1
```
