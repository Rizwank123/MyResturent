# #!/usr/bin/env bash
# DB_DATABASE_NAME=$(grep DB_DATABASE_NAME test.env | cut -d '=' -f2)
# psql -d "${DB_DATABASE_NAME}" -p "${DB_PORT}" -c "DROP SCHEMA public CASCADE;"
# psql -d "${DB_DATABASE_NAME}" -p "${DB_PORT}" -c "CREATE SCHEMA public;"



#!/bin/bash

set -e

# Load env vars
export $(grep -v '^#' test.env | xargs)

# Optional: check required vars
: "${DB_DATABASE_NAME?Missing DB_DATABASE_NAME}"
: "${DB_USERNAME?Missing DB_USERNAME}"
: "${DB_PORT?Missing DB_PORT}"

# Drop and recreate schema
echo "Resetting schema for database: $DB_DATABASE_NAME"

PGPASSWORD=$DB_PASSWORD psql -U "$DB_USERNAME" -d "$DB_DATABASE_NAME" -p "$DB_PORT" -c "DROP SCHEMA public CASCADE;"
PGPASSWORD=$DB_PASSWORD psql -U "$DB_USERNAME" -d "$DB_DATABASE_NAME" -p "$DB_PORT" -c "CREATE SCHEMA public;"

echo "Database reset successfully."