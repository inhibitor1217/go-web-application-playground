# go-web-application-playground

- This is a simple web application written in Go.
- Implements basic password authentication and authorization.

## How to run

```bash
# Setup local environment
make init

# Run the application
make dev
```

- The server will be running on port 8000.
- On port 8001, swagger documentation will be available.

## Environment variables

Fill out the following environment variables before running the application.

`.env.development`

```
APP_DOMAIN=localhost
APP_STAGE=development
AUTH_JWT_SECRET=<secret>
PSQL_HOST=localhost
PSQL_USER=<user>
PSQL_PASSWORD=<password>
PSQL_DATABASE=<database>
PUBLIC_HTTP_BASE_URL=http://localhost:8000
PUBLIC_HTTP_PORT=8000
SWAGGER_BASE_URL=http://localhost:8001
SWAGGER_PORT=8001
```
