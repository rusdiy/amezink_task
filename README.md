# amezink_task

A simple app written in Go with one endpoint.

***

## Prerequisites

- Docker already installed
- No app is listening in the HTTP port (80) already

***

## Set-Up

1) Clone this repository
2) Open the repository directory
3) Run `docker-compose up` to build and run the apps
4) Run `docker exec -i amezink_mysql mysql -utiger -pcat mezink < dump.sql` to import the mock database

***

## Endpoint

```
GET http://localhost/marks?startDate=<startDate>&endDate=<endDate>&minCount=<minCount>&maxCount=<minCount>
```
- `<startDate>` and `<endDate>` should be filled with date in yyyy-mm-dd format. e.g. 2018-01-03
- `<minCount>` and `<maxCount>` should be filled with number. e.g. 500

#### Note
- The original assignment asked me to use JSON Request Payload to input the `startDate`, `endDate`, `minCount`, and `maxCount`. As the endpoint is only used to fetch data without changing anything to the server, I think that GET the most proper HTTP Method that can be used. GET method cannot have any Request body, so I used query parameter in the URL instead.\

### Sample

```
http://localhost/marks?startDate=2000-01-01&endDate=2020-10-10&minCount=500&maxCount=1000
```
or if using cURL
```
curl -X GET localhost/marks?startDate=2000-01-01\&endDate=2020-01-01\&minCount=500\&maxCount=1000
```