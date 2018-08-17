#!/bin/bash
set -e
export COMPOSE_PROJECT_NAME=fairfaxtest
docker-compose up -d

sleep 5
docker exec couchdb curl -X PUT http://127.0.0.1:5984/articles
