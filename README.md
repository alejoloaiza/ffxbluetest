
# TECHNICAL TEST - Article API

`Test coverage is: `
## Introduction:
I decided to use Golang (my favorite programming language) to solve this problem. For me the main problem is PERFORMANCE, that is what I always think about when designing any solution, even if it's a simple API, so I dedicated my energy mainly on PERFORMANCE. I decided to use the Go library fasthttp (which 10 times faster than Go native http) and decided to use CouchDB a fast Non-SQL database written in Erlang (a high performance programming language).
But dont worry, you dont have to install anything (not even Go), everything is dockerized.


## Assumptions:
- If the "POST /articles" API receives an article request with the same ID (Already exists), will not update it, articles can be created only once. 

## Benchmarks:
Luckily, Golang its so good that has a benchmark package included, I dont need to install anything new to run the benchmarks, just create a couple of unit testing files, run a command and we are done.

These are the results for the "POST /articles API":

root@alloaiza:~/go/src/ffxbluetest# go test -bench=. -test.benchtime=3s -run=XXX
goos: linux
goarch: amd64
pkg: ffxbluetest
BenchmarkArticles-4   	     500	   8066636 ns/op

root@alloaiza:~/go/src/ffxbluetest# go test -bench=. -test.benchtime=5s -run=XXX
goos: linux
goarch: amd64
pkg: ffxbluetest
BenchmarkArticles-4   	    1000	   8354991 ns/op

root@alloaiza:~/go/src/ffxbluetest# go test -bench=. -test.benchtime=10s -run=XXX
goos: linux
goarch: amd64
pkg: ffxbluetest
BenchmarkArticles-4   	    2000	   9774476 ns/op

root@alloaiza:~/go/src/ffxbluetest# go test -bench=. -test.benchtime=20s -run=XXX
goos: linux
goarch: amd64
pkg: ffxbluetest
BenchmarkArticles-4   	    5000	   9797122 ns/op

## Results Summary: 
Machine details:                     
    Description: Notebook
    Product: Latitude 7280 (079F)
    Vendor: Dell Inc.
    Width: 64 bits
    Memory Size: 16GiB
    CPU: Intel(R) Core(TM) i7-7600U CPU @ 2.80GHz

- For 500 the Api took 8066636 nano seconds per each article creation, including the DB operation, that is 0.008066636 seconds (API + DB). 
- For 1000 the Api took 8354991 nano seconds per each article creation, including the DB operation, that is 0.008354991 seconds (API + DB). 
- For 2000 the Api took 9774476 nano seconds per each article creation, including the DB operation, that is 0.009774476 seconds (API + DB). 
- For 5000 the Api took 9797122 nano seconds per each article creation, including the DB operation, that is 0.009797122 seconds (API + DB).

## Pre-Requisites:
- Git (just to clone the repo)
- Docker and Docker-compose (On MAC they come together on Linux come separatedly)
Check using these commands
`docker --version`
`docker-compose --version`

To install Docker on MAC follow this [link](https://docs.docker.com/docker-for-mac/install/)

## Installation:
Clone this repository
`git clone https://github.com/alejoloaiza/ffxbluetest.git`


Go inside the path ffxbluetest/env and run the start.sh shell: 
`cd ffxbluetest/env && ./start.sh`

###Important: This will download two docker images, please be patient.

## Run tests and check the coverage:

`go test -cover`

# Test data for manual testing:

**Dataset 1:**      

**Expected output:**

