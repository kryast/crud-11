CRUD ke-11

POST
curl -X POST http://localhost:8080/feedbacks \
-H "Content-Type: application/json" \
-d '{
  "customer_name": "Ahmad Syarifuddin",
  "email": "ahmad@example.com",
  "message": "Pelayanan sangat baik, saya puas.",
  "rating": 5
}'
