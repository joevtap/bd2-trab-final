#!/bin/bash

set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" <<-EOSQL
    CREATE DATABASE $DB_NAME;
    CREATE USER $DB_USER WITH PASSWORD '$DB_PASSWORD';

    CREATE ROLE application;

    GRANT application TO $DB_USER;

    GRANT ALL ON DATABASE $DB_NAME TO application;
EOSQL

