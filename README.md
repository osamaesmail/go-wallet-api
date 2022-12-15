# Wallet API

# dependencies
* go v18
* docker
* docker-compose
* postgres
* mockery
* golangci-lint
* protobuf

# Setup
* run `make` for help
* run `make dev` to run all the services in docker

# features
* [x] mockery mocks
* [x] SQL Migrations
* [x] lint
* [x] protobuf generator
* [x] protobuf lint
* [x] protobuf check breaking changes
* [x] Business logic unit tests
* [x] Dockerized
* [x] Docker Compose
* [x] Logging service
* [x] REST Transport (all requests are post)
* [x] gRPC Transport
* [ ] gRPC client
* [ ] Read from os env
* [ ] Seed command to fill db with test data
* [ ] Integration tests


# endpoints
## Create user Account
### Request
```json
{
  "user_id": "8b92b25e-c2d3-4d2d-a1ef-3694df09f48e",
  "balance": 200,
  "currency": "USD"
}
```
### Response
```json
{
    "id": "ed3129f7-e3c6-413b-b291-e4dd24f3c039",
    "user_id": "8b92b25e-c2d3-4d2d-a1ef-3694df09f48e",
    "balance": "200",
    "currency": "USD"
}
```
# List user accounts
##Request
```json
{
  "user_id": "1a6fb0d2-aa91-4162-9930-161e0550adc9"
}
```
## Response
```json
[
  {
    "id": "9d15e3f5-2804-4e26-9380-062d7ffd7d7c",
    "user_id": "1a6fb0d2-aa91-4162-9930-161e0550adc9",
    "balance": "190",
    "currency": "BTC"
  },
  {
    "id": "8b92b25e-c2d3-4d2d-a1ef-3694df09f48e",
    "user_id": "1a6fb0d2-aa91-4162-9930-161e0550adc9",
    "balance": "240",
    "currency": "USD"
  }
]
```
# Create transaction
## Request
```json
{
  "from_account": "ef907719-4820-4917-b161-4cd3f418c6aa",
  "to_account": "8b92b25e-c2d3-4d2d-a1ef-3694df09f48e",
  "amount": 10
}
```
## Response
```json
{
  "ID": "eb7bb9b0-6bbe-4497-b3b9-6b7d370f4d1d",
  "FromAccount": "ef907719-4820-4917-b161-4cd3f418c6aa",
  "ToAccount": "8b92b25e-c2d3-4d2d-a1ef-3694df09f48e",
  "Amount": "10",
  "Currency": "",
  "CreatedAt": "2022-12-15T21:51:12.141933881+02:00"
}
```

# List transactions
# Request
```json
{
    "user_id": "ef907719-4820-4917-b161-4cd3f418c6aa",
    "page": 1,
    "per_page": 10
}
```
# Response
```json
{
    "data": [
        {
            "ID": "88f3cdef-f9e1-4afb-864f-4b6d37140eb8",
            "FromAccount": "ef907719-4820-4917-b161-4cd3f418c6aa",
            "ToAccount": "ef825f46-61e5-4a03-8a01-e5bd7fb4aee6",
            "Amount": "10",
            "Currency": "USD",
            "CreatedAt": "2022-12-15T14:24:31.679026Z"
        },
        {
            "ID": "9ffb0699-efbe-404e-af78-a0509298e4ce",
            "FromAccount": "9d15e3f5-2804-4e26-9380-062d7ffd7d7c",
            "ToAccount": "ef825f46-61e5-4a03-8a01-e5bd7fb4aee6",
            "Amount": "10",
            "Currency": "BTC",
            "CreatedAt": "2022-12-15T15:00:41.080961Z"
        }
    ],
    "page": 1,
    "per_page": 10,
    "total_pages": 1,
    "total_records": 2
}
```
