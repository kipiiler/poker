# Husky hold'em

## Set up

### Develop
This application is based on [Gin](https://github.com/gin-gonic/gin) framework

To run the application in development mode:

- Set up environment variable in new `.env` file. See `.env.example`.

- Install dependencies

```
go get .
```

- Synchronize dependencies (optional)
```
go mod tidy
```

- Run application
```
go run main.go
```

- Hot reload option:
```
go install github.com/cosmtrek/air@latest
air
```

Start docker container for db and cache:
```
docker compose up
```

Access postgres cli
```
 docker exec -it <container name> psql -U <user_name> -W <db_name>
```

Access redis cli
```
docker exec -it <container name> redis-cli -a <pass>
```


### Production

- Build project as binaries
```
go build -o bin/main.exe
```

## Rules
### Branches:

**`main`**: This branch is protected and requires merge request. Stand as where production code will live and will be use for deployment only. DO NOT PUSH TO THIS BRANCH

**`staging`**: This branch is protected and requires merge request. Stand as where development code will live. DO NOT PUSH TO THIS BRANCH

**`hotfix`**: You know that this is. Time crunch

**`name-docs`**: Documentation update

**`name-<issue>-<description>`**: Personal branch for any contributors to open up for merge request. The code need to pass all the test on CI/CD before requesting review.
- `<issue>`: issue number of the branch. If this a feature/bug/anything, open an issue and write out a description as well as spec for the branch. This can potentially be empty then the name should be `name-<description>`
- `<description>`: Make it short

Before pushing the code, make sure to check any linting error with

```
go vet
golangci-lint run
```

And clean up dependecies
```
go mod tidy
```

## Documentation:

Documentation is generated with [swaggo](https://github.com/swaggo). To install run following:

```
go install github.com/swaggo/swag/cmd/swag
```

Document should be written for protocol handlers (socket, https, etc...) only, which handle all the endpoints (See `internal/adapters/handlers/http`). To generate the docucmentation run

```
swag init
```

You should be able to see all the documentation in development when running the application using the paths `localhost:<PORT>/docs/index.html`. Make sure that swag init return no error and any type used in document must be defined.

## Contribution
