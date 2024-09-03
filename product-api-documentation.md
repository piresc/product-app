# Product API Documentation

## Base URL
`/api/v1`

## Data Model

### Product

| Field        | Type      | Description                                |
|--------------|-----------|--------------------------------------------|
| id           | uint      | Unique identifier for the product          |
| name         | string    | Name of the product                        |
| description  | string    | Detailed description of the product        |
| price        | float64   | Price of the product                       |
| variety      | string    | Variety or type of the product             |
| rating       | float64   | Average rating of the product              |
| stock        | int       | Available stock of the product             |
| created_at   | timestamp | Timestamp of when the product was created  |
| updated_at   | timestamp | Timestamp of when the product was last updated |

## Endpoints

### 1. Create a Product

- **URL**: `/products`
- **Method**: `POST`
- **Description**: Create a new product
- **Request Body**:
  ```json
  {
    "name": "string",
    "description": "string",
    "price": float,
    "variety": "string",
    "rating": float,
    "stock": integer
  }
  ```
- **Success Response**:
  - **Code**: 201 Created
  - **Content**: 
    ```json
    {
      "id": integer,
      "name": "string",
      "description": "string",
      "price": float,
      "variety": "string",
      "rating": float,
      "stock": integer,
      "created_at": "timestamp",
      "updated_at": "timestamp"
    }
    ```

### 2. Get All Products

- **URL**: `/products`
- **Method**: `GET`
- **Description**: Retrieve a list of all products
- **Success Response**:
  - **Code**: 200 OK
  - **Content**: Array of product objects

### 3. Get Product by ID

- **URL**: `/products/:id`
- **Method**: `GET`
- **Description**: Retrieve a specific product by its ID
- **URL Parameters**: `id=[integer]`
- **Success Response**:
  - **Code**: 200 OK
  - **Content**: Product object

### 4. Get Products by Variant

- **URL**: `/products/variant`
- **Method**: `GET`
- **Description**: Retrieve products by their variety
- **Query Parameters**: `variety=[string]`
- **Success Response**:
  - **Code**: 200 OK
  - **Content**: Array of product objects matching the variety

### 5. Get Products by Title

- **URL**: `/products/title`
- **Method**: `GET`
- **Description**: Retrieve products by their title (name)
- **Query Parameters**: `title=[string]`
- **Success Response**:
  - **Code**: 200 OK
  - **Content**: Array of product objects matching the title

### 6. Update a Product

- **URL**: `/products/:id`
- **Method**: `PUT`
- **Description**: Update an existing product
- **URL Parameters**: `id=[integer]`
- **Request Body**: 
  ```json
  {
    "name": "string",
    "description": "string",
    "price": float,
    "variety": "string",
    "rating": float,
    "stock": integer
  }
  ```
- **Success Response**:
  - **Code**: 200 OK
  - **Content**: Updated product object

### 7. Delete a Product

- **URL**: `/products/:id`
- **Method**: `DELETE`
- **Description**: Delete a specific product
- **URL Parameters**: `id=[integer]`
- **Success Response**:
  - **Code**: 200 OK
  - **Content**: Product object

## Error Responses

- **Code**: 400 Bad Request
  - **Content**: `{ "error": "Invalid input" }`

- **Code**: 404 Not Found
  - **Content**: `{ "error": "Product not found" }`

- **Code**: 500 Internal Server Error
  - **Content**: `{ "error": "Internal server error" }`

## Notes

- All timestamps are in ISO 8601 format.
- The API uses JSON for request and response bodies.
- Authentication and authorization details are not included in this documentation.
