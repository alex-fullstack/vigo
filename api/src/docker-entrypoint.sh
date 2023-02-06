#!/bin/sh

while ! nc -z $ES_HOST $ES_PORT; do
  echo "Wait for DB connection..."
  sleep 1
done

exec "$@"
