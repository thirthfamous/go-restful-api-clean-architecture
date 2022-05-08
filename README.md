# Go RESTful API Clean Architecture
Simple Go & MySQL Category CRUD with Proteced with API-Key. Coded with clean architecture and unit test

---

### Features
* Login
* Create Category
* Update Category
* Get Category
* Delete Category

### Prerequisite
* Install Go at [Go official site](https://go.dev/dl/)
* Install MySQL at [here](https://dev.mysql.com/downloads/installer/)

---

### Installation
1. Import/run the populate.sql to your database MySQL.
2. Run the commands 
```sh
// clone the project
git clone https://github.com/thirthfamous/go-restful-api-clean-architecture/

// go to the project directory
cd go-restful-api-clean-architecture

// install the dependencies
go get

// run the tests
go test ./test

// run the server
go run main.go
```

---

### Architecture
The application architecture follow the Clean Architecture

![Clean Architecture drawio](https://user-images.githubusercontent.com/30696403/167250776-f2cda279-12d7-4132-8565-88f45f124d94.png)
