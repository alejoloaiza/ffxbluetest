#!/bin/bash
./startdb.sh
export COMPOSE_PROJECT_NAME=fairfaxtest
docker-compose -f docker-compose.yml up -d goserver
sleep 15