
# TECHNICAL TEST - Article API

`Test coverage is: `
## Introduction:
I decided to take this problem personal an decided to use Golang to solve it. For me the main problem is PERFORMANCE, so I dedicated my energy mainly on that point. I decided to use the Go library fasthttp (which 10 times faster than native) and decided to use a fast Non-SQL database (CouchDB) written in Erlang (a high performance programming language).
But dont worry, you dont have to install anything. Everything is dockerized.


## Assumptions:
- If the "POST /articles" API recieves a article request with the same ID (Already exists), will not update it, articles can be created only once. 


## Pre-Requisites:
- No Go or No CouchDB required.
- Git (to clone the repo)
- Docker and Docker-compose (On MAC they come together on Linux come separatedly)
Check using this commands
`docker --version`
`docker-compose --version`

To install Docker on MAC follow this [link](https://docs.docker.com/docker-for-mac/install/)

## Installation:
Clone this repository
`git clone https://github.com/alejoloaiza/ffxbluetest.git`


Go inside the path ffxbluetest/test and run the start.sh shell: 
`cd ffxbluetest/test && ./start.sh`

###Important: This will download two docker images, please be patient.

## Run tests and check the coverage:

`go test -cover`

# Test data for manual testing:

**Dataset 1:**      

**Expected output:**

