# Weather App
## Overview
This is a minimalistic weather application built using go.
This follows Repository pattern and Domain Driven Design.

## Getting started.
```bash
# Rename the example config file
mv config.example.yml config.yml
# Set up the necessary environment variables for database credentials and JWT secrets.
# Install the dependencies
go mod tidy
# Start the app
go run cmd/main.go
```
### Testing

To run the tests for the authentication service:

```bash
go test ./internal/app/usecases/auth/test
```

## Architecture
This app follows repository pattern which provides a clean separation of concerns between configuration, domain logic, data transfer objects, infrastructure, and use cases.

### Components

- **API Layer**: The service exposes a RESTful API for client interaction. It encapsulates the core business logic and serves various endpoints for user registration, authentication, token generation, and validation.

- **Service Layer**: This layer contains the business logic for user operations and token handling. It communicates with the repository layer to persist data and utilizes JWT for secure token generation.

- **Repository Layer**: The persistence layer abstracts the database interactions. It implements the logic for data retrieval and manipulation, ensuring a separation of concerns and easier maintenance.

- **Database**: PostgreSQL is used as the database to store user credentials and related information securely.

### Security

- **Password Handling**: User passwords are securely hashed using bcrypt before being stored in the database.

- **JWT Tokens**: The service issues JWT tokens for access control, including both access and refresh tokens, ensuring secure and stateless authentication.

- **Environment Variables**: Sensitive information such as database credentials and JWT secrets are managed through environment variables, preventing hard-coded credentials within the codebase.


### Folder Structure

- `cmd/`: Contains the application's entry point.
  - `main.go`: The main executable for the application.

- `internal/`: Encapsulates the internal application logic.
  - `app/`: Core application codebase.
    - `config/`: Configuration related modules.
    - `domain/`: Domain-specific logic.
      - `entities/`: Business entities.
      - `repositories/`: Interfaces for the repository pattern.
    - `dto/`: Data Transfer Objects used to map between API and internal layers.
    - `infrastructure/`: Technical details that support the application.
      - `database/`: Database-related operations.
        - `models/`: Database models.
          - `user.go`: User model for database operations.
          - `weatherSearchHistory.go`: Model for weather search history.
        - `user.repository.impl.go`: Implementation of the user repository.
        - `weather.repository.impl.go`: Implementation of the weather repository.
      - `web/`: Web server and middleware setup.
    - `usecases/`: Application-specific use cases.
      - `auth/`: Authentication logic.
        - `handler.go`: HTTP handlers for authentication.
        - `service.go`: Business logic for authentication services.
        - `test/`: Tests for authentication services.
          - `service_test.go`: Unit tests for the authentication service.
          - `user.repository_mock.go`: Mock repository for user-related testing.
      - `weather/`: Weather information logic.
        - `handler.go`: HTTP handlers for weather information.
        - `service.go`: Business logic for weather services.

## To-Do List for Future Improvements

- [ ] **Folder structure**
  - [ ] Move the implementations to a folder
  - [ ] The middleware is now just a file. In future as more middlewares comes, it should be made into a folder.
  - [ ] Create a service configuration file that returns a struct of services to be used in the main.go
  - [ ] Maybe the `router.go` file also need to be separated into folders in the future.
  - [ ] Use `google wire` to handle dependency injection 

- [ ] **Testing**
  - [ ] Unit testing right now is only implemented in the auth service now. Increase it's coverage.

- [ ] **Error Handling**
  - [ ] Implement centralized error logging for easier debugging and monitoring.
  - [ ] Implement common response handlers. Return a common response structure.

- [ ] **Documentation**
  - [ ] Generate API documentation using tools like Swagger for easier API consumption.