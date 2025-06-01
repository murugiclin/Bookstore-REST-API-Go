📚 Documentation Summary
🔹 Book Struct

    Holds metadata of each book

    Fields: ID, Title, Author, Year

🔹 Endpoints
Method	Endpoint	Description
GET	/books	List all books
GET	/books/{id}	Get a single book
POST	/books	Add a new book
PUT	/books/{id}	Update book info
DELETE	/books/{id}	Delete a book
🔹 Data Handling

    Books are stored in-memory in a global slice books

    IDs are assigned incrementally

🔹 Libraries

    Uses github.com/gorilla/mux for robust routing

        Install it:

        go get -u github.com/gorilla/mux

✅ How to Run

    Save the code as main.go

    Install dependencies:

go get -u github.com/gorilla/mux

Run:

go run main.go

Test with Postman, cURL, or browser on http://localhost:8080/books
