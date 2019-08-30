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

**Environment Variables**

`export CACHE_HOST=<redis-host>` (if running the docker redis container, set to your DOCKER_HOST_IP)

`export API_KEY=<autopilot-api-key>`

Optional: `export CACHE_FLUSH=true` clears the redis cache

Run the server (make sure to set all env variables)

`./bin/copilot start`

Display app version

`./bin/copilot version`

Get help and display app usage

`./bin/copilot help`

### Implementation Strategy
Docker / Docker-Compose - Run redis from a docker container

Cache - Redis Cache

Gin - Middleware and routing framework, this might be overkill for this challenge but allowed me to focus on the core task instead of middleware routing details.

### Usage

Use curl or a tool like postman to test the api.

**GET** - Get contact
```bash
curl -X GET \
  http://localhost:8080/api/contact/person_9EAF39E4-9AEC-4134-964A-D9D8D54162E7 \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache'
```

Forwards request to ...
```bash
curl -X GET \
https://private-82d61-autopilot.apiary-mock.com/v1/contact/person_9EAF39E4 \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache'
```

**POST** - Upsert Contact
```bash
curl -X POST \
  http://localhost:8080/api/contact/ \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{"contact": {"FirstName": "Slarty","LastName": "Bartfast","Email":"test@slarty.com","custom": { "string--Test--Field":}}'
```

Forwards request to ...

```bash
curl -X POST \
  private-82d61-autopilot.apiary-mock.com/v1/contact \
  -H 'Content-Type: application/json' \
  -H 'cache-control: no-cache' \
  -d '{"contact": {"FirstName": "Slarty","LastName": "Bartfast","Email":"test@slarty.com","custom": { "string--Test--Field":}}'
```

### Known bugs:
- Support all APIs (right now only supports mockserver) 
- Handle case where data in cache cannot be unmarshalled (should we hit API?)
- Handle name and contact_id as key for caching (right now only supports contact_id)

### Improvements, given more time..
- Remove some hardcoded things
- Better logging / api response formatting
- Testcases (mock http requests)
- Better error handling and logging
- Improved cache invalidation
- Improve availability (handle timeouts, rate-limiting, graceful shutdown, retries, etc..)
- Refactor some of the API methods that got out of hand and cleanup code
- Refactor the middleware API to accept interfaces and return requested structs of data
- Better cache invalidation (right now I just delete from cache when a PUT/POST request is made)
- Improved handling of environment variables/config
- Remove globals
