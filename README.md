# Banking Ledger Microservice

## Overview
The **Banking Ledger Microservice** is a distributed system built using **Golang**, **Kafka**, **PostgreSQL**, **MongoDB**, and **gRPC**. It facilitates account management, transaction processing, and notification services with a clean and scalable microservices architecture. The project includes an **API Gateway** for request routing, a **Banking Service** for managing accounts and transactions, and a **Notification Service** for event-driven messaging.

## Features
- **Account Management**: Create, retrieve, and manage user accounts.
- **Transactions**: Process deposits, withdrawals, and track balances.
- **Notifications**: Event-driven notifications for transaction updates.
- **API Gateway**: Routes requests to appropriate microservices.
- **Event-Driven Architecture**: Uses Kafka for real-time message streaming.
- **Mock Testing**: Implements `sqlmock` and `gomock` for unit testing.
- **gRPC Communication**: Efficient inter-service communication with protocol buffers.
- **Dockerized Deployment**: Containerized with Docker and orchestrated via Docker Compose.

## Tech Stack
- **Programming Language**: Golang
- **Frameworks & Libraries**: Gin, gRPC
- **Databases**: PostgreSQL (Relational), MongoDB (NoSQL)
- **Messaging**: Apache Kafka
- **DevOps Tools**: Docker, Docker Compose
- **Testing**: sqlmock, gomock

## Project Structure
```
Banking Ledger
│   docker-compose.yaml
│
├───Api-Gateway
│   ├── cmd/
│   │   └── main.go
│   ├── pkg/
│   │   ├── client/
│   │   ├── config/
│   │   ├── di/
│   │   ├── handler/
│   │   ├── model/
│   │   ├── pb/
│   │   └── routes/
│
├───Banking Service
│   ├── cmd/
│   │   └── main.go
│   ├── pkg/
│   │   ├── config/
│   │   ├── db/
│   │   ├── di/
│   │   ├── domain/
│   │   ├── mock/
│   │   ├── repository/
│   │   ├── server/
│   │   ├── usecase/
│   │   └── utils/
│
└───Notification Service
    ├── cmd/
    │   └── main.go
    ├── pkg/
    │   ├── config/
    │   ├── db/
    │   ├── di/
    │   ├── model/
    │   ├── pb/
    │   ├── repository/
    │   ├── server/
    │   └── usecase/
```

## Installation & Setup
### Prerequisites
Ensure you have the following installed:
- Golang
- Docker & Docker Compose
- Kafka
- PostgreSQL & MongoDB

### Running the Application
#### 1. Clone the repository
```sh
 git clone https://github.com/ShahabazSulthan/Banking-Ledger.git
 cd Banking-Ledger
```

#### 2. Start services with Docker Compose
```sh
 docker-compose up --build
```

#### 3. Run individual services manually (if needed)
```sh
# Run API Gateway
cd Api-Gateway
 go run cmd/main.go

# Run Banking Service
cd ../Banking Service
 go run cmd/main.go

# Run Notification Service
cd ../Notification Service
 go run cmd/main.go
```

## API Endpoints
### Account Management (API Gateway)
| Method | Endpoint | Description |
|--------|-----------------|--------------------------|
| POST | `/api/account` | Create a new account |
| GET | `/api/account/:id` | Get account details |
| PATCH | `/api/account/:id/balance` | Update account balance |

### Transactions
| Method | Endpoint | Description |
|--------|-----------------|--------------------------|
| POST | `/api/transaction` | Process a transaction |
| GET | `/api/transaction/:id` | Get transaction details |
| GET | `/api/transactions/:account_id` | Get transactions by account |

### Notifications
| Method | Endpoint | Description |
|--------|-----------------|--------------------------|
| GET | `/api/notifications/:account_id` | Fetch notifications |

## Kafka Event Flow
1. **Banking Service** publishes transaction events to Kafka.
2. **Notification Service** consumes the events and generates notifications.
3. **API Gateway** retrieves notifications when requested.

## Testing
### Mock Testing
- **`sqlmock`**: Used for mocking database interactions in repository tests.
- **`gomock`**: Used for mocking interfaces in use case testing.

Run tests with:
```sh
# Run tests for Banking Service
cd Banking Service
 go test ./...

```



## Contact
- **Author**: Shahabaz Sulthan
- **Email**: [shahabazsulthan4@gmail.com](mailto:shahabazsulthan4@gmail.com)
- **GitHub**: [ShahabazSulthan](https://github.com/ShahabazSulthan)
- **LinkedIn**: [Shahabaz Sulthan](https://www.linkedin.com/in/shahabaz-sulthan-a256252b3/)

