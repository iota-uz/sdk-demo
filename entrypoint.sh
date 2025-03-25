#!/bin/sh
set -e

PGPASSWORD="$DB_PASSWORD" psql -h "$DB_HOST" -U "$DB_USER" -c "DROP DATABASE IF EXISTS $DB_NAME;"
PGPASSWORD="$DB_PASSWORD" psql -h "$DB_HOST" -U "$DB_USER" -c "CREATE DATABASE $DB_NAME;"

# Run all commands in sequence with error handling
echo "Running database migration..."
if ! /build/migrate up; then
    echo "Migration failed with exit code $?"
    exit 1
fi
echo "Migration completed successfully."

echo "Seeding database..."
if ! /build/seed_db; then
    echo "Database seeding failed with exit code $?"
    exit 1
fi
echo "Database seeding completed successfully."

echo "Starting server..."
/build/run_server
