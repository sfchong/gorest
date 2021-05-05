# gorest
Go rest service example. Use [chi](https://github.com/go-chi/chi) router. Designed to be simple and idiomatic.

## Project Structure
Inspired by [Golang Standard](https://github.com/golang-standards/project-layout) and [Ben Johnson's WTF Dial](https://github.com/benbjohnson/wtf) project structure.

`\ ` Root directory contains all the application domain type e.g. User/Customer/Sale

`cmd` - main application package.

`const` -  global constants 

`http` - http server, router and api logic

`mock` - mock service which returns dummy data for testing.

## Development
```shell
go run ./cmd/rest/main.go
```

## TODO
- [ ] add router middleware
- [ ] add jwt authentication
- [ ] add logging
- [ ] read value from config file
- [ ] add postgresql implementation
- [ ] write test
- [ ] add basic html homepage