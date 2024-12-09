#!/bin/bash
source .env
PGPASSWORD=$DB_ADM_PWD psql -h localhost -v DB_NAME=$DB_NAME -v DB_USR=$DB_USR -v DB_PWD=$DB_PWD -U $DB_ADM_USR -d $DB_NAME < init.sql 
