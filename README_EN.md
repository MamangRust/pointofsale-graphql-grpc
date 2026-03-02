# Point Of Sale (GraphQL & gRPC)

This project is an Ecommerce system that provides two API interfaces: GraphQL and gRPC. The system is built using Go and follows Clean Architecture principles to separate application layers.

## Key Features

- **Dual API:** Provides GraphQL endpoints for query flexibility and gRPC for high-performance inter-service communication.
- **Authentication & Authorization:** Uses JWT to secure endpoints and role-based access control for managing user permissions.
- **User Management:** Registration, authentication, and role-based access management
- **Merchant Management:** Multi-merchant support with comprehensive configuration
- **Product Management:** Inventory tracking, barcode generation, categorization
- **Category Management:** Flexible product organization with categories
- **Order and Transaction System:** Real-time and accurate sales processing

---

## Architecture Overview

This project adopts a Layered Architecture approach inspired by Clean Architecture. It separates concerns into clear layers: `Handler`, `Service`, `Repository`, and `Domain`.

- **Handlers (gRPC & GraphQL):** The outermost layer that receives requests from clients. This layer is responsible for input validation, calling the appropriate services, and formatting responses.
- **Services:** Contains the core business logic of the application. Services orchestrate data from various repositories and perform complex operations.
- **Repositories:** Data access layer responsible for communicating with the database. It abstracts database queries from the business logic layer.
- **Domain/Model:** Represents the core entities and value objects of the system.

```mermaid
graph TD
    subgraph Client
        A[User / Client App]
    end

    subgraph API Gateway
        B[GraphQL API Gateway]
    end

    subgraph Application Core
        C[gRPC API Server]
        D[Middlewares Auth, Rate Limit, etc]
        E[Handlers gRPC & GraphQL Resolvers]
        F[Services Business Logic]
        G[Repositories Data Access Layer]
    end

    subgraph Infrastructure
        H[PostgreSQL Database]
        I[Proto & GraphQL Schema]
    end

    A --> B
    B -->|Resolver calls via gRPC| C
    C --> D
    D --> E
    E --> F
    F --> G
    G --> H
    E -- Uses contracts from --> I

```

---

## Entity Relationship Diagram (ERD)

The following diagram represents the relationships between the main entities in the database.

```mermaid
erDiagram
    USERS {
        INT user_id PK
        VARCHAR firstname
        VARCHAR lastname
        VARCHAR email
        VARCHAR password
    }

    ROLES {
        INT role_id PK
        VARCHAR role_name
    }

    USER_ROLES {
        INT user_role_id PK
        INT user_id FK
        INT role_id FK
    }

    REFRESH_TOKENS {
        INT refresh_token_id PK
        INT user_id FK
        VARCHAR token
    }

    MERCHANTS {
        INT merchant_id PK
        INT user_id FK
        VARCHAR name
        TEXT address
    }

    CASHIERS {
        INT cashier_id PK
        INT merchant_id FK
        INT user_id FK
        VARCHAR name
    }

    CATEGORIES {
        INT category_id PK
        VARCHAR name
        VARCHAR slug_category
    }

    PRODUCTS {
        INT product_id PK
        INT merchant_id FK
        INT category_id FK
        VARCHAR name
        INT price
        INT count_in_stock
    }

    ORDERS {
        INT order_id PK
        INT merchant_id FK
        INT cashier_id FK
        BIGINT total_price
    }

    ORDER_ITEMS {
        INT order_item_id PK
        INT order_id FK
        INT product_id FK
        INT quantity
        INT price
    }

    TRANSACTIONS {
        INT transaction_id PK
        INT order_id FK
        INT merchant_id FK
        VARCHAR payment_method
        INT amount
    }

    USERS ||--o{ USER_ROLES : "has"
    ROLES ||--o{ USER_ROLES : "has"
    USERS ||--o{ REFRESH_TOKENS : "has"
    USERS ||--|| MERCHANTS : "owns"
    MERCHANTS ||--o{ CASHIERS : "employs"
    USERS ||--o{ CASHIERS : "is"
    MERCHANTS ||--o{ PRODUCTS : "sells"
    CATEGORIES ||--o{ PRODUCTS : "contains"
    MERCHANTS ||--o{ ORDERS : "generates"
    CASHIERS ||--o{ ORDERS : "creates"
    ORDERS ||--o{ ORDER_ITEMS : "includes"
    PRODUCTS ||--o{ ORDER_ITEMS : "appears in"
    ORDERS ||--|| TRANSACTIONS : "results in"
    MERCHANTS ||--o{ TRANSACTIONS : "processes"
```

---

## Usage Guide

### Prerequisites

- [Go](https://golang.org/doc/install) (version 1.18 or newer)
- [Docker](https://www.docker.com/get-started) and Docker Compose
- [Make](https://www.gnu.org/software/make/)
- [protoc](https://grpc.io/docs/protoc-installation/)

### Installation & Setup

1.  **Clone Repository**

    ```bash
    git clone https://github.com/MamangRust/pointofsale-graphql-grpc.git
    cd pointofsale-graphql-grpc
    ```

2.  **Environment Configuration**
    Create a `.env` file in the project root directory by copying from the example.

        ```bash
        cp .env.example .env
        ```

        Adjust the variables in the `.env` file according to your local configuration.

        **Example `.env`:**

        ```env
        # Postgres
        DB_DRIVER=postgres

        DB_HOST=postgres
        DB_HOST=localhost
        DB_PORT=5432
        DB_USERNAME=postgres
        DB_PASSWORD=postgres
        DB_NAME=ecommerce_grpc

        DB_MAX_OPEN_CONNS=50
        DB_MAX_IDLE_CONNS=10
        DB_CONN_MAX_LIFETIME=30m

        DB_SEEDER=false

        SECRET_KEY=yantopedia

        DB_URL=postgres://postgres:postgres@localhost:5432/ecommerce_grpc
        ```

    If not, make sure you have a running PostgreSQL instance and its configuration matches the `.env` file.

3.  **Install Dependencies**

    ```bash
    go mod tidy
    ```

4.  **Generate Code from Proto**
    Every time you modify `.proto` files, run this command to update the generated Go code.

    ```bash
    make generate-proto
    ```

5.  **Run Database Migrations**
    This command will create all required tables in your database.

    ```bash
    make migrate
    ```

6.  **Run the Server**
    ```bash
    go run cmd/server/main.go
    ```
    The server will now be running.
    - **gRPC Server:** `localhost:50051`
    - **GraphQL Server:** `http://localhost:8080/query`

### Makefile Commands

- `make migrate`: Run database migrations to the latest version.
- `make migrate-down`: Revert the last database migration.
- `make generate-proto`: Generate Go code from protobuf files.
