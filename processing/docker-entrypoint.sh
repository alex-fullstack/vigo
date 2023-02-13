#!/bin/sh

while ! nc -z bus:4222; do
  echo "Wait for bus connection..."
  sleep 1
done

exec "$@"