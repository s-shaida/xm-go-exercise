# XM gateway service
## Development
### How-to run locally
1. Make sure your local Kubernetes cluster is up and running
2. Execute
```shell
make kubernetes-start
```
3. For work with API you can use postman collection or
```shell
# register user
curl --request POST \
        --url http://localhost:32717/auth/register \
        --header 'Content-Type: application/json' \
        --data '{
        "email": "test@user.com",
        "password": "12345678"
        }'
```
```shell
# login user
curl --request POST \
        --url http://localhost:32717/auth/login \
        --header 'Content-Type: application/json' \
        --data '{
        "email": "test@user.com",
        "password": "12345678"
        }'
```
```shell
# create company
curl --request POST \
  --url http://localhost:32717/company/ \
  --header 'Authorization: Bearer <token>' \
  --header 'Content-Type: application/json' \
  --data '{
    "name": "Company 1",
    "description": "Small description",
    "amount": 15,
    "registered": false,
    "type": "NonProfit"
    }'
```
```shell
# get company by id
curl --request GET \
  --url http://localhost:32717/company/<company_id> \
  --header 'Authorization: Bearer <token>'
```
```shell
# patch company
curl --request PATCH \
  --url http://localhost:32717/company/ \
  --header 'Authorization: Bearer <token>' \
  --header 'Content-Type: application/json' \
  --data '{
    "id": <company_id>,
    "name": "Company 1",
    "description": "Small description",
    "amount": 10,
    "registered": false,
    "type": "NonProfit"
    }'
```
```shell
# delete company
curl --request DELETE \
  --url http://localhost:32717/company/<company_id> \
  --header 'Authorization: Bearer <token>'
```