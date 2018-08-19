#!/bin/bash
docker kill $(docker ps -q)
docker rm $(docker ps -aq)
docker rmi $(docker images dev-* -q)
docker volume rm $(docker volume ls -qf dangling=true)
set -e
export COMPOSE_PROJECT_NAME=fairfaxtest
docker-compose up -d

