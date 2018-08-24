#!/bin/bash
set -e
export COMPOSE_PROJECT_NAME=fairfaxtest
docker-compose -f docker-compose.yml up -d couchdb
sleep 15

