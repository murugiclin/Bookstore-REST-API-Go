# 📚 Go REST API - Book Management

A simple **CRUD REST API** built in Go using the **Gorilla Mux router**. This project manages a collection of books with in-memory storage (no database needed). Perfect as a starter template for learning REST API development in Go.

---

## 🚀 Features

* Get all books (`GET /books`)
* Get a single book by ID (`GET /books/{id}`)
* Create a new book (`POST /books`)
* Update an existing book (`PUT /books/{id}`)
* Delete a book (`DELETE /books/{id}`)
* Seeded with sample data on startup

---

## 🛠️ Technologies Used

* **Go (Golang)** – main programming language
* **net/http** – built-in Go HTTP server
* **encoding/json** – JSON serialization/deserialization
* **strconv** – string ↔ int conversions
* **Gorilla Mux** – router for handling URL parameters

---

## 📂 Project Structure

```
.
├── main.go        # Entry point, seeds data and starts the server
├── go.mod         # Module definition & dependencies
└── go.sum         # Dependency checksums
```

---

## ⚙️ Installation & Setup

### 1. Clone the Repository

```bash
git clone https://github.com/murugiclin/Bookstore-REST-API-Go.git
cd Bookstore-REST-API-Go
```

### 2. Initialize Go Module

```bash
go mod init go-books-api
go mod tidy
```

### 3. Install Dependencies

```bash
go get -u github.com/gorilla/mux
```

### 4. Run the Server

```bash
go run main.go
```

The server will start on:

```
http://localhost:8080
```

---

## 🌍 API Endpoints

### 1. **Get All Books**

```http
GET /books
```

**Response:**

```json
[
  {"id":1,"title":"1984","author":"George Orwell","year":1949},
  {"id":2,"title":"To Kill a Mockingbird","author":"Harper Lee","year":1960},
  {"id":3,"title":"The Great Gatsby","author":"F. Scott Fitzgerald","year":1925}
]
```

---

### 2. **Get Book by ID**

```http
GET /books/{id}
```

**Example:**

```bash
curl http://localhost:8080/books/1
```

**Response:**

```json
{"id":1,"title":"1984","author":"George Orwell","year":1949}
```

---

### 3. **Create Book**

```http
POST /books
```

**Example:**

```bash
curl -X POST http://localhost:8080/books \
  -H "Content-Type: application/json" \
  -d '{"title":"Brave New World","author":"Aldous Huxley","year":1932}'
```

**Response:**

```json
{"id":4,"title":"Brave New World","author":"Aldous Huxley","year":1932}
```

---

### 4. **Update Book**

```http
PUT /books/{id}
```

**Example:**

```bash
curl -X PUT http://localhost:8080/books/1 \
  -H "Content-Type: application/json" \
  -d '{"title":"Nineteen Eighty-Four","author":"George Orwell","year":1949}'
```

**Response:**

```json
{"id":1,"title":"Nineteen Eighty-Four","author":"George Orwell","year":1949}
```

---

### 5. **Delete Book**

```http
DELETE /books/{id}
```

**Example:**

```bash
curl -X DELETE http://localhost:8080/books/2
```

**Response:**

```
HTTP/1.1 204 No Content
```

---

## 🧪 Testing

You can test endpoints using:

* `curl` (examples above)
* Postman
* Any REST client of your choice

---

## 📌 Notes

* Data is stored **in-memory** → changes are lost when the server restarts.
* To persist data, integrate with a database (e.g., PostgreSQL, MySQL, MongoDB).

---

## 🤝 Contributing

1. Fork the repo
2. Create a new branch (`feature/your-feature`)
3. Commit your changes
4. Push to your branch
5. Open a Pull Request

---

## 📜 License

This project is licensed under the **MIT License**.

---

## ✨ Author

👤 **Your Name**
🔗 [GitHub](https://github.com/murugiclin)


---

Happy coding! 🚀
