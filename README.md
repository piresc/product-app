# Product Service API

## Description
This API manages product details with CRUD (Create, Read, Update, Delete) operations. It allows users to create new products, retrieve existing products, update product information, and delete products from the database.

## Technologies Used
- **Go**: The programming language used for building the API.
- **Echo**: A high-performance, extensible web framework for Go.
- **PostgreSQL**: The database used to store product information.
- **SQLX**: A library that provides extensions on Go's standard database/sql library.
- **Wire**: A dependency injection tool for Go.
- **Viper**: A configuration management library for Go.

## Getting Started

### Prerequisites
- Docker
- Go (1.22.6 or later)

### Running the Application

1. **Docker Compose**:
   configs
   ```json
      POSTGRES_DB: productdb 
      POSTGRES_USER: productuser
      POSTGRES_PASSWORD: productpass
   ````
   command
   ```bash
   docker-compose up -d
   ```


### API Endpoints

- `POST /api/products`: Create a new product
- `GET /api/products`: Retrieve all products
- `GET /api/products/:id`: Retrieve a product by ID
- `GET /api/products/:id`: Retrieve a product by ID
- `GET /api/products/variant`: Retrieve a product by variant
- `PUT /api/products/title`: Update a product by title
- `DELETE /api/products/:id`: Delete a product by ID

### Database Initialization
The application uses PostgreSQL as the database. The initial database schema and sample data can be found in the `init.sql` file. The following table is created:

- **products**
  - `id`: SERIAL PRIMARY KEY
  - `name`: TEXT NOT NULL
  - `description`: TEXT
  - `price`: NUMERIC(10, 2) NOT NULL
  - `variety`: TEXT
  - `rating`: NUMERIC(2, 1)
  - `stock`: INT
  - `created_at`: TIMESTAMP DEFAULT CURRENT_TIMESTAMP
  - `updated_at`: TIMESTAMP DEFAULT CURRENT_TIMESTAMP

### Configuration
The application configuration is managed using Viper. The configuration file is in YAML format and includes settings for the application and database connection.

The configuration: 
```yaml 
app: 
   port: "8080" 
db: 
  host: "localhost"
  port: "5432"
  user: "productuser"
  password: "productpass"
  dbname: "productdb"
  sslmode: "disable"
```

### Running Tests
To run tests, use the following command:
```bash
go test ./...
```
## License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.