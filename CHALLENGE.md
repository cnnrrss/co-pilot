### Autopilot API caching middleware server

Implement a Go caching middleware server for Autopilot API /contact method
https://autopilot.docs.apiary.io/

#introduction/getting-help
- Create GET / POST / PUT /contact endpoint
- Retrieve a requested contact from redis, if it is not present retrieve it from Autopilot API and store in redis
- Create / Update a contact and invalidate redis cache after POST / PUT requests
- Cover the necessary methods with unit tests
- Write README file with instructions how to run and test it


Should only take "a couple of hours"
- Upload to github