# Chaos Theory Internship - Take Home Task

`in-memory` branch - project satisfying minimal requirements with in-memory storage.

`main` branch - later project with additional functionalities I added. Including: SQLite DB as a storage, better validation, more endpoints.

## Available endpoints:

GET `/list`
POST `/add`

GET `/key/{key}`
DELETE `/key/{key}`
GET `/date/before/{year}/{month}/{day}`
GET `/date/after/{year}/{month}/{day}`

## How I developed this webserver?

1. **Laid out basic requirements.** I was required to set up a `/list` and `/add` endpoints, I identified what is needed and what is format for input/output.
2. **Set up test.** Before writing any code I set up a Postman requests with all required payloads, to test my server. I know and used test in Golang before, but decided that it will be too tedious to write tests for such a small project.
3. **Write minimal Chi example.** I didn't use Chi before, so I set up a minimal working example with some "Hello, World!" style endpoints.
4. **Start writing Webserver itself.** I started writing whole program in a *top-down* manner. First, by setting up endpoints, then `ItemServer` and `Item` structs. Then started writing general outline for handler functions.
5. **In-memory store.** When handler functions ready, I defined all functions needed for data storage. I started implementing all the functions one-by-one and testing them. After ensuring that it works, I move to the next one. 
6. **Better storage.** After satisfying basic requirements, I saved my work in  a separate branch and started working on integrating SQLite DB. I rewrote `ItemServer` and handlers to use db connector in `main.go`. Then, I created new package in internal called `sqlitestore`. Implemented functions one by one. At first I deleted and recreated DB each time program launches but later, when needed faster testing, fixed this by checking if .db file already exists.
7. **MORE ENPOINTS!!!** At this point I got minimal working version with SQLite DB. I decided to add more enpoints, that might be useful. I again did it in top-down manner but since main format of functions is outlined, it is much easier to write other handlers. 
8. **Refactoring.** After I decided that there are enough endpoints and adding more will unnecessarily overcomplicate project, I started refactoring code. I decomposed repeated code into helper functions and added error handling in some places I missed before. Added comments where necessary. And created `Dockerfile` for both versions of my project.

