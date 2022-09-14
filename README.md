# DDD

> Documentation Driven Development

A small demo demonstrating how using documentation to generate server
implementation can eliminate the incompatibilities of the API contracts with
clients.

## Getting Started

To run the server
```sh
$ make
```

To generate code from the API documentation in [`openapi.yaml`](openapi.yaml)
```sh
$ make gen
```

## Tools
* [`oapi-codegen`](https://github.com/deepmap/oapi-codegen)
