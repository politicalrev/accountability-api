# Accountability API

We can mourn the state of affairs, or we can use the elections as a wake-up call. This API seeks to provide a generalized and unbiased political promise tracker. The idea is that we can take the things that politicians claim they'll do and keep a record of it. Any time they accomplish a goal or break a promise, we can track it -- with sources. When it comes time to vote again, this tool could be used to gauge how effective a politician was at accomplishing their goals. The judgement as to why isn't in the domain of this tool. That, as always, is up to the voter.

## Contributing

### WIP

This API isn't complete. It can only spit out the information already in the database. Check repo issues for the TODO.

### Prerequisites

- Docker

### Building

First time:

1. `docker-compose up --build` to build the application and data containers
2. `docker -exec -it {accountabilityapi_web_1 container ID} bash` to connect to the application container
3. `goose up` to run any outstanding migrations

The container uses [gin](https://github.com/codegangsta/gin) to auto-reload the server when a change is detected. So if you make a change, you don't need to rebuild.

To bring up the container after it has already been built, use `docker-compose up` (without `--build`).
