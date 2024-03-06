# Go-Bookstore

Go-Bookstore is a simple web application built with Go that allows users to manage a bookstore. It includes features such as creating, updating, and deleting books, as well as retrieving book details. The application uses the Gorilla Mux router and GORM for database operations.

## Installation

1. Clone the repository: `git clone https://github.com/denilbhatt0814/go-bookstore.git`
2. Navigate to the project directory: `cd go-bookstore`
3. Install the dependencies: `go mod download`

## Configuration

1. Create a new database in your preferred database management system (MySQL, PostgreSQL, SQLite, or SQL Server).
2. Update the database configuration in the `config/config.go` file with your database credentials.

## Usage

1. Build the application: `go build`
2. Run the application: `./go-bookstore`

## API Endpoints

- `GET /books` - Retrieve all books
- `GET /books/{id}` - Retrieve a specific book by ID
- `POST /books` - Create a new book
- `PUT /books/{id}` - Update an existing book
- `DELETE /books/{id}` - Delete a book

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.
