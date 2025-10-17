# Point Of Sale (GraphQL & gRPC)

Proyek ini adalah sebuah sistem Ecommerce yang menyediakan dua antarmuka API: GraphQL dan gRPC. Sistem ini dibangun menggunakan Go dan mengikuti prinsip-prinsip arsitektur bersih (Clean Architecture) untuk memisahkan lapisan-lapisan aplikasi.

## Fitur Utama

- **Dual API:** Menyediakan endpoint GraphQL untuk fleksibilitas query dan gRPC untuk komunikasi _high-performance_ antar servis.
- **Autentikasi & Otorisasi:** Menggunakan JWT untuk mengamankan endpoint dan sistem berbasis peran (role-based) untuk mengatur hak akses.
- **Manajemen Pengguna**: Registrasi, autentikasi, dan manajemen role-based access
- **Manajemen Merchant**: Multi-merchant support dengan konfigurasi lengkap
- **Manajemen Produk**: Inventory tracking, barcode generation, categorization
- **Manajemen Kategori**: Organisasi produk dengan kategori yang fleksibel
- **Sistem Order dan Transaksi**: Proses penjualan yang real-time dan akurat

---

## Arsitektur Overview

Arsitektur proyek ini mengadopsi pendekatan _Layered Architecture_ yang terinspirasi dari _Clean Architecture_. Ini memisahkan _concerns_ menjadi beberapa lapisan yang jelas: `Handler`, `Service`, `Repository`, dan `Domain`.

- **Handlers (gRPC & GraphQL):** Lapisan terluar yang menerima permintaan dari klien. Lapisan ini bertanggung jawab untuk validasi input, memanggil service yang sesuai, dan memformat respons.
- **Services:** Lapisan ini berisi logika bisnis inti dari aplikasi. Services mengorkestrasi data dari berbagai _repository_ dan melakukan operasi yang kompleks.
- **Repositories:** Lapisan akses data yang bertanggung jawab untuk berkomunikasi dengan database. Ini mengabstraksi query database dari lapisan bisnis.
- **Domain/Model:** Merepresentasikan entitas inti dan objek nilai dari sistem.

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

Diagram berikut merepresentasikan hubungan antar entitas utama dalam database.

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

## Cara Penggunaan

### Prasyarat

- [Go](https://golang.org/doc/install) (versi 1.18 atau lebih baru)
- [Docker](https://www.docker.com/get-started) dan Docker Compose
- [Make](https://www.gnu.org/software/make/)
- [protoc](https://grpc.io/docs/protoc-installation/)

### Instalasi & Setup

1.  **Clone Repositori**

    ```bash
    git clone https://github.com/MamangRust/pointofsale-graphql-grpc.git
    cd pointofsale-graphql-grpc
    ```

2.  **Konfigurasi Environment**
    Buat file `.env` di root direktori proyek dengan menyalin dari contoh.

        ```bash
        cp .env.example .env
        ```

        Sesuaikan variabel di dalam file `.env` sesuai dengan konfigurasi lokal Anda.

        **Contoh `.env`:**

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

    Jika tidak, pastikan Anda memiliki instance PostgreSQL yang berjalan dan konfigurasinya sesuai dengan file `.env`.

3.  **Instalasi Dependensi**

    ```bash
    go mod tidy
    ```

4.  **Generate Kode dari Proto**
    Setiap kali Anda mengubah file `.proto`, jalankan perintah ini untuk memperbarui kode Go yang di-generate.

    ```bash
    make generate-proto
    ```

5.  **Jalankan Migrasi Database**
    Perintah ini akan membuat semua tabel yang diperlukan di database Anda.

    ```bash
    make migrate
    ```

6.  **Jalankan Server**
    ```bash
    go run cmd/server/main.go
    ```
    Server sekarang akan berjalan.
    - **gRPC Server:** `localhost:50051`
    - **GraphQL Server:** `http://localhost:8080/query`

### Perintah Makefile

- `make migrate`: Menjalankan migrasi database ke versi terbaru.
- `make migrate-down`: Me-revert migrasi database terakhir.
- `make generate-proto`: Men-generate kode Go dari file-file protobuf.
