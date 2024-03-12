## Receipt Processor

This is the fetch microservice written as per the specs in https://github.com/fetch-rewards/receipt-processor-challenge

### How to run it:

There are two ways you can run this app:

With go directly (must have go installed):

`go mod download && go mod verify && go build && ./fetch-service`

With docker (must have docker installed)

`docker build -t fetch-service . && docker run -p 8083:8083 fetch-service ./fetch-service`

### API Specification

The base path of the api is the following when running locally:

`localhost:8083/api`

You can view the api in swagger by going to the following link after building the service:

http://localhost:8083/swagger/index.html#/receipts

There are two endpoints you can hit:

`POST` `/receipts/process` or localhost:8083/api/receipts/process

`GET` `/receipts/{id}/points` or localhost:8083/api/receipts/{id}/points

They function as required in the source specifications. 

Unit tests are also included, if you want to run them you can type `go test -v ./...` into the command line.






