#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/miration -database "$DB_SOURCE" -verbose up

echo "start the app"
exec "$@"