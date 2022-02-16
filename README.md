# Chaos Theory Internship - Take Home Task

`in-memory` branch - project satisfying minimal requirements with in-memory storage.

`main` branch - later project with additional functionalities I added. Including: SQLite DB as a storage, better validaion, more endpoints.

## Table of Contents

* [To run Webserver](#to-run-webserver)
* [Available endpoints](#available-endpoints)
* [How I developed this webserver?](#how-dev)
* [Curl commands to check server](#curl)

## To run webserver

```bash
docker build -t chaostheory-task
docker run --rm -p 3000:3000 test-server
```

This will run local webserver on port 3000. I used 3000 just because 80 will require running binary as system administrator.

## Available endpoints

GET `/list` - return all items in timestamp descending order. JSON Response:
```json
[{"timestamp": "2019-12-02T06:53:32Z", "key": "a", "value": "some value"},
 {"timestamp": "2019-12-02T06:53:35Z", "key": "asdf", "value": "some other value"}]
```

POST `/add` - adds item into DB, overrides item with same key if exists. Payload:

```json
{"key": "asdf", "value": "some other value"}
```

GET `/key/{key}` - return item with given key. JSON Response:

```json
{"timestamp": "2019-12-02T06:53:32Z", "key": "a", "value": "some value"}
```

DELETE `/key/{key}` - deletes item with given key from DB, does nothing if item does not exists.
GET `/date/before/{year}/{month}/{day}` - returns all items before given date (< not inclusive) in timestamp descending order. Output format same as in `/list`.
GET `/date/after/{year}/{month}/{day}` - returns all items after given date (>= inclusive) in timestamp descending order. Output format same as in `/list`.

## How I developed this webserver? <a name="how-dev" />

1. **Laid out basic requirements.** I was required to set up a `/list` and `/add` endpoints, I identified what is needed and what is format for input/output.
2. **Set up test.** Before writing any code I set up a Postman requests with all required payloads, to test my server. I know and used test in Golang before, but decided that it will be too tedious to write tests for such a small project.
3. **Write minimal Chi example.** I didn't use Chi before, so I set up a minimal working example with some "Hello, World!" style endpoints.
4. **Start writing Webserver itself.** I started writing whole program in a *top-down* manner. First, by setting up endpoints, then `ItemServer` and `Item` structs. Then started writing general outline for handler functions.
5. **In-memory store.** When handler functions ready, I defined all functions needed for data storage. I started implementing all the functions one-by-one and testing them. After ensuring that it works, I move to the next one.
6. **Better storage.** After satisfying basic requirements, I saved my work in  a separate branch and started working on integrating SQLite DB. I rewrote `ItemServer` and handlers to use db connector in `main.go`. Then, I created new package in internal called `sqlitestore`. Implemented functions one by one. At first I deleted and recreated DB each time program launches but later, when needed faster testing, fixed this by checking if .db file already exists.
7. **MORE ENPOINTS!!!** At this point I got minimal working version with SQLite DB. I decided to add more enpoints, that might be useful. I again did it in top-down manner but since main format of functions is outlined, it is much easier to write other handlers.
8. **Refactoring.** After I decided that there are enough endpoints and adding more will unnecessarily overcomplicate project, I started refactoring code. I decomposed repeated code into helper functions and added error handling in some places I missed before. Added comments where necessary. And created `Dockerfile` for both versions of my project.

*During development I assumed any error in `internal/sqlitestore` package will stop program. Because it is relatively simple and that package having error means, something really wrong happened.

## Curl commands to check server <a name="curl" />

`curl localhost:3000/list`

` curl -X POST localhost:3000/add -H 'Content-Type: application/json' -d '{"key":"key5", "value": "value5"}'`

`curl -X GET localhost:3000/key/key5 -H 'Content-Type: application/json' -d '{"key":"key5", "value": "value5"}'`

`curl -X DELETE localhost:3000/key/key5 -H 'Content-Type: application/json' -d '{"key":"key5", "value": "value5"}'`

`curl -X GET localhost:3000/date/before/2022/02/16`

`curl -X GET localhost:3000/date/after/2022/02/16`t
