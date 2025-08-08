CRUD ke-11

POST
curl -X POST http://localhost:8080/feedback \
-H "Content-Type: application/json" \
-d '{
  "customer_name": "Ahmad Syarifuddin",
  "email": "ahmad@example.com",
  "message": "Pelayanan sangat baik, saya puas.",
  "rating": 5
}'


GET
curl http://localhost:8080/feedback
curl http://localhost:8080/feedback/1


PUT
curl -X PUT http://localhost:8080/feedback/1 \
-H "Content-Type: application/json" \
-d '{
  "customer_name": "Ahmad Updated",
  "email": "ahmad.updated@example.com",
  "message": "Pelayanan semakin baik setelah update.",
  "rating": 4
}'

DELETE
curl -X DELETE http://localhost:8080/feedback/1