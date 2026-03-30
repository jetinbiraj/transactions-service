# Transactions Service (Go)

A simple microservice to manage accounts and financial transactions.


### Swagger API Documentation
Open terminal and `cd cmd`
- Required Dependency :- go install github.com/swaggo/swag/cmd/swag@latest
- Format the SWAG comments :- `swag fmt -d ./,../internal/transactions,../internal/accounts`
- Generate swagger doc :- `swag init -o ../swagger -d ./,../internal/transactions,../internal/accounts`
- Swaggo official GitHub :- https://github.com/swaggo/swag