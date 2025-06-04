# go-api-restFul 

# how-to-implement-a-restful-web-service-in-go

A simple Go based REST API server built using gorilla/mux HTTP router and MongoDB database.

Prerequisites
Basic familiarity with Go and MongoDB.

Run 
$ go run main.go
or
$ go build main.go
$ ./main

Endpoints
You can use Postman to perform CRUD operations using following endpoints:

-> GET /api/books
-> POST /api/book
-> GET /api/book/{id}
-> PUT /api/book/{id}
-> DELETE /api/book/{id}