curl -X POST http://192.168.1.194:1122/api/v1/houses \
  -H 'Content-Type: application/json' \
  -d '{
        "request_id": "create_house_req_id",
        "house": {
            "id": "house_id"
        }
    }' | jq

curl -X POST http://localhost:1122/api/v1/users \
  -H 'Content-Type: application/json' \
  -d '{
        "request_id": "create_user_req_id",
        "user": {
            "id": "user_id"
        }
    }' | jq

