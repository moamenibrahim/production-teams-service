# production-teams-service

A simple web service for managing production teams at a company

## Requirements

- GoLang v1.17.1: https://go.dev/learn/
- PostgreSQL for windows: https://www.postgresql.org/download/
- Docker Desktop for Windows v4.5.1: https://www.docker.com/products/docker-desktop
- Docker compose version v2.2.3 (part of Docker Desktop for Windows)

## About the service

A simple web service for managing production teams at a company. The data structure consists of three entities called `Hub`, `Team`, and `User` in the order of managerial hierarchy: [lucidchart database diagram](https://lucid.app/lucidchart/c67b5a0e-473a-482a-a6aa-5bc14e2774d1/edit?invitationId=inv_775213ab-46cf-4590-b080-497ba4f3eff4)

- `Hub` is an entity that associates with `Team` based on their geological location. 

- `Team` is an entity that associates with `User` based on their production type. 

- `User` is an entity that stores users information.


## Run the Application

Depending on your preference, there are two options for running the application:

### Option 1: Using `golang` from terminal:

Make sure `postgreSQL` database is up and running and especially `pgadmin` to make sure service is available

Create database manually named `productionService`, we use separate database for this service to allow isolation from the main database of postgres

Finally, you should be able to run the application straight from terminal:

```bash
go run .
```

### Option 2: Using `docker-compose`:


If you have Docker desktop installed you should be able to build and run with docker-compose:

```bash
docker-compose up -d --build
```

## Tests

### tests types

- Unit tests: using mocks
- Integeration tests: trigger endpoints using `http-clients` and parsing responses

### Run tests

```bash
go test ./... --cover -p 1
```

## Future improvements

- Relations between entities is still not fully implemented, planned implementation to be as in the diagram here: [lucidchart database diagram](https://lucid.app/lucidchart/c67b5a0e-473a-482a-a6aa-5bc14e2774d1/edit?invitationId=inv_775213ab-46cf-4590-b080-497ba4f3eff4) but first iteration of the backend does not include `Foreign keys`
- `PostgreSQL` provides `GIS` data support that can be used as datatype for `geo_location` column in `hubs` table
- Kubernetes deployment should be straight forward since we have `Dockerfile` in place, using `Gitlab` free docker registry, for example, we can push the generated docker image to the remote registry and use it in the deployment scripts
- Move to `gin`: soon after starting the work, I noticed that Gin provides better functionalities and testing setups, I would prefer to see the project use such library more

## Author

- Moamen Ibrahim <mmibrahem76@gmail.com>
