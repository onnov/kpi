#!/bin/sh

go build ./cmd/kpi

LISTENED_ADDRESS=":3000" \
DB_HOST=127.0.0.1:3366 \
DB_NAME=kpi \
DB_USER=kpi \
DB_PASSWORD=kpi \
./kpi