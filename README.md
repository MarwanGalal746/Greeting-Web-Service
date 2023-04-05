<div align="center">
  <br>
  <h1>Greeting Web Service</h1>
</div>

## Table of contents

- [**Summary**](#summary)
- [**Getting started**](#getting-started)
- [**Prerequisites**](#prerequisites)
- [**Packages**](#packages)
- [**Features**](#features)
- [**Running**](#running)
- [**Logs**](#logs)
- [**Postman collection**](#postman-collection)

---

## Summary

Simple web service responding to requests with a greeting message. The service is written
in [Golang](https://golang.org/dl/).

## Getting Started

Clone the repository.<br />
Follow the instructions to complete the installation.

## Prerequisites

- [Golang](https://golang.org/dl/)
- [Postman](https://www.postman.com/) (optional)

## Packages

- [Gin](https://github.com/gin-gonic/gin)
- [Validator](https://github.com/go-playground/validator)
- [Testify](https://github.com/stretchr/testify)

## Features

- Upon startup read the HTTP port to listen on from command line arguments and it uses a default value 8080 if command
  line argument is not specified. For example:

```bash
go run main.go --port=9999 #application will run on port 9999
go run main.go  #application will run on port 8080
```

- The webservice have one end point (/greet) that accept POST request with JSON payload looking like this:

```json
{
  “name”: “Sherif”
}
```

and it parses the JSON request and upon success respond with HTTP status 200 OK and a JSON body that looks like this:

```json
{
  “message”: “Hello Sherif!”
}
```

- The webservice responses with HTTP status 405 (Method Not Allowed) if the request method on /greet is anything other
  than POST.
- The webservice responses with HTTP status 400 (Bad Request) for any incorrect requests on the /greet endpoint (for
  example the payload is not valid JSON, the Content-Type header is not application/json or the name property is
  missing).
- The webservice responses with HTTP status 404 (Not Found) on any URL other than /greet.
- The service handles SIGINT or SIGTERM signals (kill command on Linux or Ctrl+C on both Windows and Linux) and exit
  gracefully.

## Running

- In your cloned directory.
- open your terminal and run:

```bash
go run main.go
```

The server will start by default at:

- Local: http://localhost:8080

You can specify port number by passing it as an argument:
```bash
go run main.go --port=9999 #application will run on port 9999
```

## Logs

After running the server, **logs.txt** file will be created in the root directory of the project. It will contain the logs.

## Postman collection

you will find the postman collection [here](postman%20collection/greeting.postman_collection.json).

You can get what is the right structure of JSON file to send requests and receiving responses from the postman
collection after importing it in the [Postman](https://www.postman.com/).