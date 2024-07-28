
# Book Store Management




## Table of Contents

- Packages used
- Installation
- Usage
- Routes

## Packages Used

- GORM : ORM to interact with postgres DB
- GIN : For handelling http Requests
- CompileDaemon : Used for hot reload
- GoDotEnv : To configure Env file.## Installation

To install the project dependencies, follow these steps:

```sh
# Clone the repository
git clone https://github.com/sid00002/go-crud.git

# Navigate to the project directory
cd your-repository

# Tidy up Go module dependencies
go mod tidy

## Usage

- go run main.go
  Server will start on http://localhost:3000
  
## Routes

- /posts/ Get aall the books stored in DB
- /posts/ (POST) Create a new Book
- /posts/:id (GET) Get existing book by bookId
- /posts/:id (PUT) Update a Book
