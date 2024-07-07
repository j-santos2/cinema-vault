# Cinema Vault API

The Cinema Vault API is a web service written in Go that provides endpoints for managing movies and user authentification.
This README contains an overview of the API endpoints, how to use them and additional information.

## Routes

### Health Check

- **URL**: `/v1/healthcheck`
- **Method**: `GET`
- **Description**: Endpoint for checking the health status of the API.

### Movies

#### List Movies

- **URL**: `/v1/movies`
- **Method**: `GET`
- **Description**: Retrieves a list of movies.
- **Permissions Required**: `movies:read`

#### Create Movie

- **URL**: `/v1/movies`
- **Method**: `POST`
- **Description**: Creates a new movie.
- **Permissions Required**: `movies:write`

#### Show Movie

- **URL**: `/v1/movies/:id`
- **Method**: `GET`
- **Description**: Retrieves details of a specific movie identified by `:id`.
- **Permissions Required**: `movies:read`

#### Update Movie

- **URL**: `/v1/movies/:id`
- **Method**: `PATCH`
- **Description**: Updates details of a specific movie identified by `:id`.
- **Permissions Required**: `movies:write`

#### Delete Movie

- **URL**: `/v1/movies/:id`
- **Method**: `DELETE`
- **Description**: Deletes a specific movie identified by `:id`.
- **Permissions Required**: `movies:write`

### Users

#### Register User

- **URL**: `/v1/users`
- **Method**: `POST`
- **Description**: Registers a new user.

#### Activate User

- **URL**: `/v1/users/activated`
- **Method**: `PUT`
- **Description**: Activates a user.

### Authentication

#### Create Authentication Token

- **URL**: `/v1/tokens/authentication`
- **Method**: `POST`
- **Description**: Generates an authentication token.

### Metrics

- **URL**: `/v1/metrics`
- **Method**: `GET`
- **Description**: Provides metrics about the API using expvar.

## Middleware

The application uses various middleware functions to handle authentication, rate limiting, CORS, error recovery, and metrics:

- `authenticate`: Handles user authentication.
- `rateLimit`: Implements rate limiting for API endpoints.
- `enableCORS`: Enables Cross-Origin Resource Sharing (CORS) for API requests.
- `recoverPanic`: Recovers from panics in the application and returns a proper error response.
- `metrics`: Tracks metrics related to each request.

## License

This project is licensed under the MIT License - see the LICENSE file for details.
