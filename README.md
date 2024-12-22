Calculator

to run 
in cmd run: go run cmd/main/main.go



Description
калькулятор
чтобы проверить введите в cmd

curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
"expression": "2+2*2"
}'



curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
"expression": "2+2*2)"
}'


curl --location 'localhost:8080/api/v1/calculate' \
--header 'Content-Type: application/json' \
--data '{
"expression": "asdfafd"
}'