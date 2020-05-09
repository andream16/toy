# Toy

Toy is an example of a tiny microservice built with SOLID & DDD principles.

It's separation of concerns, allows to potentially plug-in other transports implementations like GRPC or AMQP re-using the same business logic.

## What does toy do?

- saves toys in memory by their `name` and `description`.
- returns all the available toys. An error is returned if the number of toys is `odd`.
- deletes the oldest toy.

## API

| Endpoint    | Method | Body                                                                     | HTTP Status Codes |
|-------------|:------:|:------------------------------------------------------------------------:|------------------:|
| /           | Get    | `[{"name": "Jotaro Kujo Action Figure","description": "Action Figure"}]` | 200               |
| /           | Put    | `{"name": "Dio Action Figure","description": "Action Figure"}]`          | 201, 400          |
| /           | Delete |                                                                          | 200               |

## What you need to run the project

- go: ~v1.14
- docker: ~v18.09.9
- docker-compose: ~v1.23.2

## Commands

- `make run`: runs the app on port `:8080`
- `make stop`: stops the app
- `make build`: builds the application
- `make test`: runs all go tests
- `make vendor`: downloads golang vendors