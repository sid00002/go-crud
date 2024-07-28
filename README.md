
# GO CRUD

## Table of Contents

- [Packages Used](#packages-used)
- [Installation](#installation)
- [Usage](#usage)
- [Routes](#routes)

## Packages Used

- **GORM**: ORM to interact with PostgreSQL DB
- **GIN**: For handling HTTP requests
- **CompileDaemon**: Used for hot reload
- **GoDotEnv**: To configure environment files

## Routes

- /posts/ Get aall the books stored in DB
- /posts/ (POST) Create a new Book
- /posts/:id (GET) Get existing book by bookId
- /posts/:id (PUT) Update a Book

## Installation

To install the project dependencies, follow these steps:

```sh
# Clone the repository
git clone https://github.com/sid00002/go-crud.git

# Navigate to the project directory
cd go-crud

# Tidy up Go module dependencies
go mod tidy





