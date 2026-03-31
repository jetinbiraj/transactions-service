# Transactions Service (Go)

A simple microservice to manage accounts and financial transactions.

## App requirements

- HomeBrew :- `/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"`
- Git :- `brew install git`
- Docker :- `brew install docker`
- Colima :- `brew install colima`
- Docker Compose :- `brew install docker-compose`

## Clone Application

`git clone https://github.com/jetinbiraj/transactions-service.git && cd transactions-service`

## Database Setup

- Database is optional and postgres db is enabled by config, if you do not wish to setup postgres simply comment out the
  line
  `DB_NAME` in config.yaml, application will use then in-memory database
- Follow the steps for postgres setup
  ```shell
  # Pull the docker image if does not exists and run
  docker run --name postgres \                                        
  -e POSTGRES_PASSWORD=postgres \
  -e POSTGRES_DB=postgres \
  -p 5432:5432 \
  -d postgres
  
  # Once the docker image is pulled, start the container using the command:
  docker start postgres
  ```

## Build and Run Application

1. Start the docker instance using colima :- `colima start`
2. Build the application :- `docker build -t transactions-service .`
3. Run application docker image only if you're not using postgres:- `docker run -p 8080:8080 transactions-service`
4. Run the application using docker compose with postgres enabled by config, It spins off both application and postgres
   db together :- `docker-compose up`
5. To stop the docker compose, since no storage attached to container, the data will be lost :- `docker compose down`

## Service Interaction

- Via Swagger UI
    - Open the url in the browser `http://localhost:8080/swagger/index.html`, expand the API, click on `Try it out` and
      click on `Execute`
- Via http files in the codebase
    - If you cloned the repo in IntelliJ IDE then open `http` directory
    - Open any http file and hit the green play button for associated http request to hit endpoint
- Bruno API Collection
    - Bruno API collection added in the codebase, install the Bruno App in your machine and import the collection in
      Bruno app
    - Once the collections are imported, you can navigate to http request and hit the endpoint
- Curl requests
  ```shell
    # Creata Account
    curl --request POST \
        --url http://localhost:8080/accounts \
        --header 'content-type: application/json' \
        --data '{"document_number": "12345678900"}'
  
    # Get Account
    curl --request GET \
       --url http://localhost:8080/accounts/1

     # Create Transaction
      curl --request POST \
         --url http://localhost:8080/transactions \
         --header 'content-type: application/json' \
         --data '{
              "account_id": 1,
              "operation_type_id": 1,
              "amount": -55
         }'

## Swagger API Documentation

Open terminal and `cd cmd`

- Required Dependency :- go install github.com/swaggo/swag/cmd/swag@latest
- Format the SWAG comments :- `swag fmt -d ./,../internal/transactions,../internal/accounts`
- Generate swagger doc :- `swag init -o ../swagger -d ./,../internal/transactions,../internal/accounts --pd`
- Swaggo official GitHub :- https://github.com/swaggo/swag

## Architecture Overview

The service follows a clean layered architecture:

### Layers

- **Handler**: Handles HTTP requests, validation, and response formatting
- **Service**: Contains business logic and domain rules
- **Repository**: Abstracts data persistence (in-memory implementation)

### Design Principles

- Separation of concerns
- Dependency injection via interfaces
- Testability (mocked repositories)
- Simplicity over over-engineering

## Data Model

### Account

- `account_id` (int64) – unique identifier
- `document_number` (string) – same document number can be used to create multiple accounts

### Transaction

- `transaction_id` (int64)
- `account_id` (int64)
- `operation_type_id` (int)
- `amount` (float64)
- `event_date` (timestamp, UTC)

## Business Rules

### Account

- `document_number` must:
    - be non-empty
    - contain only digits

### Transaction

- `account_id` must exist
- `operation_type_id` must be one of:
    - 1: Normal Purchase
    - 2: Installment Purchase
    - 3: Withdrawal
    - 4: Credit Voucher

- Amount rules:
    - For operation types 1, 2, 3 → amount must be **negative**
    - For operation type 4 → amount must be **positive**

- `event_date` is generated automatically in UTC

## API Endpoints

### Accounts

- `POST /accounts` → Create account
- `GET /accounts/{accountId}` → Get account details

### Transactions

- `POST /transactions` → Create transaction

---

## Testing Section

```md
## Running Tests

Run unit & integration test:

```bash
   go test ./...
```

## Configuration Section

```md
## Configuration

Configuration is loaded using Viper from: config/config.yaml
```

### Notes

- Application will fail to start if config file is missing
- In Docker, config directory is bundled inside the container

## Docker Notes

- Multi-stage build is used for a minimal runtime image
- Config file bundled into the container

### Verify container

```bash
curl http://localhost:8080/health
```

# Trade-offs Section

## Design Decisions & Trade-offs

### In-Memory Storage

- Used for simplicity and faster execution
- Simulates indexed access using maps (O(1) lookups)

### Why not a database?

- Keeps focus on API design and business logic
- Repository abstraction allows easy replacement with DB

### Time Handling

- All timestamps stored in UTC
- Generated at service layer

### Validation

- Explicit validation functions used instead of external libraries
- Provides better control over business rules

## Future Improvements

- Restrict user to have limited number of accounts per documentId
- Add authentication/authorization
- Add rate limiting
- Add structured logging
- Add metrics (Prometheus)