# Bookstore api CRUD app
 ### API Methods
 
`id` is integer
1. Get all books: `GET` `books/`
2. Get book: `GET` `books/{id}` with JSON response `{
  "title": "Start with Why",
  "author": "Simon Sinek"
}`
3. Create book: `POST` `books/` with data structure like  `{
  "title": "Start with Why",
  "author": "Simon Sinek"
}`
4. Update book: `PATCH` `books/{id}` with data structure like  `{
  "title": "The Infinite Game"
}`
4. Delete book: `DELETE` `books/{id}`
