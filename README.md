CRUD - 1

POST
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name":"Ahmad", "email":"ahmad@example.com"}'

GET
curl -X GET http://localhost:8080/users

curl -X GET http://localhost:8080/users/1