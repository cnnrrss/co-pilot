# Co-Pilot

Your co-pilot to navigating the Autopilot Contacts API.

## Install

`go get -u github.com/cnnrrss/co-pilot`

OR 

`git clone github.com/cnnrrss/co-pilot`

### Prerequisites

- Go >=1.11
- Docker

## Build the app

Switch to co-pilot working directory
`cd path/to/dir`

Build the services environment (redis container)
`docker-compose up --build`

Build the copilot app
`go build .`

## Start the Server

Run the server (make sure to set your redis host environment variable)
`CACHE_HOST=<redis-host> ./bin/copilot start`

Display app version
`./bin/copilot version`

Get help and display app usage
`./bin/copilot help`

### Implementation Strategy
Docker / Docker-Compose - Run redis from a docker container
Cache - Redis Cache
Gin - Middleware and routing framework

### Usage

Use curl or a tool like postman to test the api.

##### Single Contact

**GET** 
