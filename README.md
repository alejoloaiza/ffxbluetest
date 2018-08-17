
# TECHNICAL TEST - Article API

`Test coverage is: `
## Assumptions:
- I'm going to assume that this API requires great performance because will be used on a big company, so I decided to use fasthttp library and couchdb (on a container)


## Pre-Requisites:
- Docker and Docker-compose (On MAC they come together on Linux come separatedly)
`docker --version`
`docker-compose --version`

To install Docker on MAC follow this [link](https://docs.docker.com/docker-for-mac/install/)
## Installation:
Clone this repository
`git clone https://github.com/alejoloaiza/ffxbluetest.git`


Go inside the path ffxbluetest/test and run the start.sh shell: 
`cd ffxbluetest/test && ./start.sh`




## Run the containers using these commands:

`cd test`
`docker-compose -d `
###Important: This will download two docker images, please be patient.

## Run tests and check the coverage:

`go test -cover`

# Test data for manual testing:

**Dataset 1:**      

**Expected output:**

