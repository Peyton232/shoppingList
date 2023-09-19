#!/bin/bash

# Read database and user credentials from environment variables
DB_USER="$DB_USER"
DB_NAME="$DB_NAME"
DB_PASSWORD="$DB_PASSWORD"

if [ -z "$DB_USER" ] || [ -z "$DB_NAME" ] || [ -z "$DB_PASSWORD" ]; then
  echo "Error: Database credentials not provided via environment variables."
  exit 1
fi

# Create the database and user
echo "Creating database and user..."
psql -U postgres -c "CREATE DATABASE $DB_NAME;"
psql -U postgres -c "CREATE USER $DB_USER WITH ENCRYPTED PASSWORD '$DB_PASSWORD';"
psql -U postgres -c "ALTER ROLE $DB_USER SET client_encoding TO 'utf8';"
psql -U postgres -c "ALTER ROLE $DB_USER SET default_transaction_isolation TO 'read committed';"
psql -U postgres -c "ALTER ROLE $DB_USER SET timezone TO 'UTC';"
psql -U postgres -c "GRANT ALL PRIVILEGES ON DATABASE $DB_NAME TO $DB_USER;"

# Run PostgreSQL migrations
echo "Running PostgreSQL migrations..."
pgmigrate -migrations-dir=./migrations -database="postgres://$DB_USER:$DB_PASSWORD@localhost/$DB_NAME?sslmode=disable" up

echo "Database initialization complete."
