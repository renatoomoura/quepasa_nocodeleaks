#!/bin/sh
set -e

echo "Container's IP address: `awk 'END{print $1}' /etc/hosts`"
echo "Working dir: `pwd`"

cd /opt/quepasa/
go build -o service main.go
./service
exec "$@"