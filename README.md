# Book-Author CRUD API

This is a RESTful API built with Go and PostgreSQL for managing books and authors. The API provides CRUD operations for both entities, allowing you to create, read, update, and delete books and authors.

## Project Structure

The project follows a typical Go structure:

Sure, here's a README.md file for the book-author CRUD API project:

markdown


Copy code
# Book-Author CRUD API

This is a RESTful API built with Go and PostgreSQL for managing books and authors. The API provides CRUD operations for both entities, allowing you to create, read, update, and delete books and authors.

## Project Structure

The project follows a typical Go structure:

```
bookstore/
├── api/
│   ├── handlers/
│   │   ├── author_handler.go
│   │   └── book_handler.go
│   └── routes.go
├── models/
│   ├── author.go
│   └── book.go
├── repository/
│   ├── author_repository.go
│   └── book_repository.go
├── service/
│   ├── author_service.go
│   └── book_service.go
├── main.go
├── go.mod
└── go.sum
```

## Prerequisites

- Go (version 1.16 or later)
- PostgreSQL

## Setup

1. Clone the repository
2. Install dependencies: `go get ./...`
3. Create a PostgreSQL database (e.g., `mydb`)
4. Update the database connection string in `main.go` with your PostgreSQL credentials

## Running the Application

1. Open a terminal and navigate to the project directory
2. Run the application: `go run main.go`
3. The server will start running on `http://localhost:8080`

## API Endpoints

### Authors

- `POST /authors`: Create a new author
- `GET /authors`: Get a list of all authors
- `GET /authors/:id`: Get an author by ID
- `PUT /authors/:id`: Update an author by ID
- `DELETE /authors/:id`: Delete an author by ID

### Books

- `POST /books`: Create a new book
- `GET /books`: Get a list of all books
- `GET /books/:id`: Get a book by ID
- `PUT /books/:id`: Update a book by ID
- `DELETE /books/:id`: Delete a book by ID

## Testing

You can test the API endpoints using tools like Postman, curl, or a web browser. Here are some example requests:

```bash
# Create an author
curl -X POST -H "Content-Type: application/json" -d '{"name":"John Doe"}' http://localhost:8080/authors

# Get all authors
curl http://localhost:8080/authors

# Create a book
curl -X POST -H "Content-Type: application/json" -d '{"title":"Book Title", "author_id":1}' http://localhost:8080/books

# Get all books
curl http://localhost:8080/books
