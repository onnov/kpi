#!/bin/sh

GOOSE_DRIVER=mysql \
GOOSE_DBSTRING="kpi:kpi@tcp(127.0.0.1:3366)/kpi" \
goose -dir migrations up